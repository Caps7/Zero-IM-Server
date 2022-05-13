package rpclogic

import (
	"context"
	"github.com/showurl/Zero-IM-Server/app/msg-gateway/cmd/wsrpc/internal/rpcsvc"

	"github.com/showurl/Zero-IM-Server/app/msg-gateway/cmd/wsrpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OnlinePushMsgLogic struct {
	ctx    context.Context
	svcCtx *rpcsvc.ServiceContext
	logx.Logger
}

func NewOnlinePushMsgLogic(ctx context.Context, svcCtx *rpcsvc.ServiceContext) *OnlinePushMsgLogic {
	return &OnlinePushMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OnlinePushMsgLogic) OnlinePushMsg(in *pb.OnlinePushMsgReq) (*pb.OnlinePushMsgResp, error) {
	// todo: add your logic here and delete this line

	return &pb.OnlinePushMsgResp{}, nil
}
