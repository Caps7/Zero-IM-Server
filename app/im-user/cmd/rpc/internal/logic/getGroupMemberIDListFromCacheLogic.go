package logic

import (
	"context"

	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/internal/svc"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupMemberIDListFromCacheLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupMemberIDListFromCacheLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupMemberIDListFromCacheLogic {
	return &GetGroupMemberIDListFromCacheLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupMemberIDListFromCacheLogic) GetGroupMemberIDListFromCache(in *pb.GetGroupMemberIDListFromCacheReq) (*pb.GetGroupMemberIDListFromCacheResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetGroupMemberIDListFromCacheResp{
		CommonResp: &pb.CommonResp{
			ErrCode: 0,
			ErrMsg:  "",
		},
		MemberIDList: []string{
			"1",
			"2",
			"3",
			"4",
			"5",
			"6",
			"7",
			"8",
			"9",
			"10",
		},
	}, nil
}
