package logic

import (
	"context"

	"github.com/showurl/Zero-IM-Server/app/msg/cmd/rpc/internal/svc"
	"github.com/showurl/Zero-IM-Server/app/msg/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelMsgListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelMsgListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelMsgListLogic {
	return &DelMsgListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelMsgListLogic) DelMsgList(in *pb.WrapDelMsgListReq) (*pb.WrapDelMsgListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.WrapDelMsgListResp{}, nil
}
