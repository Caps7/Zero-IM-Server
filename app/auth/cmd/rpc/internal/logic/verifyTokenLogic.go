package logic

import (
	"context"

	"github.com/showurl/Zero-IM-Server/app/auth/cmd/rpc/internal/svc"
	"github.com/showurl/Zero-IM-Server/app/auth/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyTokenLogic {
	return &VerifyTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyTokenLogic) VerifyToken(in *pb.VerifyTokenReq) (*pb.VerifyTokenResp, error) {
	// todo: add your logic here and delete this line

	return &pb.VerifyTokenResp{
		Uid:     in.SendID,
		Success: true,
		ErrMsg:  "",
	}, nil
}
