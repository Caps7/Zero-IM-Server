package logic

import (
	"context"

	"github.com/showurl/Zero-IM-Server/app/conversation/cmd/rpc/internal/svc"
	"github.com/showurl/Zero-IM-Server/app/conversation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyConversationFieldLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyConversationFieldLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyConversationFieldLogic {
	return &ModifyConversationFieldLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifyConversationFieldLogic) ModifyConversationField(in *pb.ModifyConversationFieldReq) (*pb.ModifyConversationFieldResp, error) {
	// todo: add your logic here and delete this line

	return &pb.ModifyConversationFieldResp{}, nil
}
