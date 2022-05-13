package logic

import (
	"context"

	"github.com/showurl/Zero-IM-Server/app/msg-callback/cmd/rpc/internal/svc"
	"github.com/showurl/Zero-IM-Server/app/msg-callback/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CallbackWordFilterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCallbackWordFilterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackWordFilterLogic {
	return &CallbackWordFilterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CallbackWordFilterLogic) CallbackWordFilter(in *pb.CallbackWordFilterReq) (*pb.CallbackWordFilterResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CallbackWordFilterResp{}, nil
}
