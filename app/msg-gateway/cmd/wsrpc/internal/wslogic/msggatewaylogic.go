package wslogic

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/showurl/Zero-IM-Server/app/auth/cmd/rpc/pb"
	chatpb "github.com/showurl/Zero-IM-Server/app/msg/cmd/rpc/pb"
	"github.com/showurl/Zero-IM-Server/common/utils"
	"github.com/zeromicro/go-zero/core/trace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"net/http"
	"sync"
	"time"

	"github.com/showurl/Zero-IM-Server/app/msg-gateway/cmd/wsrpc/internal/types"
	"github.com/showurl/Zero-IM-Server/app/msg-gateway/cmd/wsrpc/internal/wssvc"

	"github.com/zeromicro/go-zero/core/logx"
	oteltrace "go.opentelemetry.io/otel/trace"
)

type UserConn struct {
	*websocket.Conn
	w *sync.Mutex
}

type MsggatewayLogic struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *wssvc.ServiceContext
	wsMaxConnNum int
	wsUpGrader   *websocket.Upgrader
	wsConnToUser map[*UserConn]map[string]string
	wsUserToConn map[string]map[string]*UserConn
}

func (l *MsggatewayLogic) runWithCtx(f func(ctx context.Context), kv ...attribute.KeyValue) {
	propagator := otel.GetTextMapPropagator()
	tracer := otel.GetTracerProvider().Tracer(trace.TraceName)
	ctx := propagator.Extract(context.Background(), propagation.HeaderCarrier(http.Header{}))
	spanName := utils.CallerFuncName()
	spanCtx, span := tracer.Start(
		ctx,
		spanName,
		oteltrace.WithSpanKind(oteltrace.SpanKindServer),
		oteltrace.WithAttributes(kv...),
	)
	defer span.End()
	propagator.Inject(spanCtx, propagation.HeaderCarrier(http.Header{}))
	f(spanCtx)
}

var msgGatewayLogic *MsggatewayLogic

func NewMsggatewayLogic(ctx context.Context, svcCtx *wssvc.ServiceContext) *MsggatewayLogic {
	if msgGatewayLogic != nil {
		return msgGatewayLogic
	}
	ws := &MsggatewayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
	ws.wsMaxConnNum = ws.svcCtx.Config.Websocket.MaxConnNum
	ws.wsConnToUser = make(map[*UserConn]map[string]string)
	ws.wsUserToConn = make(map[string]map[string]*UserConn)
	ws.wsUpGrader = &websocket.Upgrader{
		HandshakeTimeout: time.Duration(ws.svcCtx.Config.Websocket.TimeOut) * time.Second,
		ReadBufferSize:   ws.svcCtx.Config.Websocket.MaxMsgLen,
		CheckOrigin:      func(r *http.Request) bool { return true },
	}
	msgGatewayLogic = ws
	return msgGatewayLogic
}

func (l *MsggatewayLogic) Msggateway(req *types.Request) (*types.Response, bool) {
	if len(req.Token) != 0 && len(req.SendID) != 0 && len(req.PlatformID) != 0 {
		// 调用rpc验证token
		resp, err := l.svcCtx.AuthService.VerifyToken(l.ctx, &pb.VerifyTokenReq{
			Token:    req.Token,
			Platform: req.PlatformID,
			SendID:   req.SendID,
		})
		if err != nil {
			logx.WithContext(l.ctx).Errorf("调用 VerifyToken 失败, err: %s", err.Error())
			return &types.Response{
				Uid:     "",
				ErrMsg:  "调用 VerifyToken 失败",
				Success: false,
			}, false
		}
		if !resp.Success {
			logx.WithContext(l.ctx).Infof("VerifyToken 失败, err: %s", resp.ErrMsg)
			return &types.Response{
				Uid:     resp.Uid,
				ErrMsg:  resp.ErrMsg,
				Success: false,
			}, false
		}
		return &types.Response{
			Uid:     resp.Uid,
			ErrMsg:  "",
			Success: true,
		}, true
	}
	return &types.Response{
		Uid:     "",
		ErrMsg:  "参数错误",
		Success: false,
	}, false
}

func (l *MsggatewayLogic) WsUpgrade(uid string, req *types.Request, w http.ResponseWriter, r *http.Request, header http.Header) error {
	conn, err := l.wsUpGrader.Upgrade(w, r, header)
	if err != nil {
		return err
	}
	newConn := &UserConn{conn, new(sync.Mutex)}
	userCount++
	l.addUserConn(uid, req.PlatformID, newConn, req.Token)
	go l.readMsg(newConn, uid, req.PlatformID)
	return nil
}

func (l *MsggatewayLogic) readMsg(conn *UserConn, uid string, platformID string) {
	for {
		messageType, msg, err := conn.ReadMessage()
		if messageType == websocket.PingMessage {
			l.sendMsg(context.Background(), conn, Resp{
				ReqIdentifier: 0,
				MsgIncr:       "",
				ErrCode:       0,
				ErrMsg:        "",
				Data:          []byte("pong"),
			})
		}
		if err != nil {
			uid, platform := l.getUserUid(conn)
			logx.Error("WS ReadMsg error ", " userIP ", conn.RemoteAddr().String(), " userUid ", uid, " platform ", platform, " error ", err.Error())
			userCount--
			l.delUserConn(conn)
			return
		}
		l.runWithCtx(func(ctx context.Context) {
			l.msgParse(ctx, conn, msg)
		}, attribute.KeyValue{
			Key:   "uid",
			Value: attribute.StringValue(uid),
		}, attribute.KeyValue{
			Key:   "platformID",
			Value: attribute.StringValue(platformID),
		})
	}
}

func (l *MsggatewayLogic) getSeqReq(ctx context.Context, conn *UserConn, m *Req) {
	rpcReq := chatpb.GetMaxAndMinSeqReq{}
	nReply := new(chatpb.GetMaxAndMinSeqResp)
	rpcReq.UserID = m.SendID
	rpcReply, err := l.svcCtx.MsgRpc.GetMaxAndMinSeq(ctx, &rpcReq)
	if err != nil {
		logx.WithContext(ctx).Error("rpc call failed to getSeqReq", err, rpcReq.String())
		nReply.ErrCode = 500
		nReply.ErrMsg = err.Error()
		l.getSeqResp(ctx, conn, m, nReply)
	} else {
		logx.WithContext(ctx).Info("rpc call success to getSeqReq", rpcReply.String())
		l.getSeqResp(ctx, conn, m, rpcReply)
	}
}

func (l *MsggatewayLogic) getSeqResp(ctx context.Context, conn *UserConn, m *Req, pb *chatpb.GetMaxAndMinSeqResp) {
	var mReplyData chatpb.GetMaxAndMinSeqResp
	mReplyData.MaxSeq = pb.GetMaxSeq()
	mReplyData.MinSeq = pb.GetMinSeq()
	b, _ := proto.Marshal(&mReplyData)
	mReply := Resp{
		ReqIdentifier: m.ReqIdentifier,
		MsgIncr:       m.MsgIncr,
		ErrCode:       pb.GetErrCode(),
		ErrMsg:        pb.GetErrMsg(),
		Data:          b,
	}
	l.sendMsg(ctx, conn, mReply)
}
