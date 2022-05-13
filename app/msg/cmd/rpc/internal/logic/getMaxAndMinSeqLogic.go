package logic

import (
	"context"

	"github.com/showurl/Zero-IM-Server/app/msg/cmd/rpc/internal/svc"
	"github.com/showurl/Zero-IM-Server/app/msg/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMaxAndMinSeqLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMaxAndMinSeqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMaxAndMinSeqLogic {
	return &GetMaxAndMinSeqLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMaxAndMinSeqLogic) GetMaxAndMinSeq(in *pb.GetMaxAndMinSeqReq) (*pb.GetMaxAndMinSeqResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetMaxAndMinSeqResp{}, nil
}
