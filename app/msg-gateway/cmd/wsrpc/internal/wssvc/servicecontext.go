package wssvc

import (
	"github.com/showurl/Zero-IM-Server/app/auth/cmd/rpc/authservice"
	"github.com/showurl/Zero-IM-Server/app/msg-gateway/cmd/wsrpc/internal/wsconfig"
	"github.com/showurl/Zero-IM-Server/app/msg/cmd/rpc/chat"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      wsconfig.Config
	AuthService authservice.AuthService
	MsgRpc      chat.Chat
}

func NewServiceContext(c wsconfig.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		AuthService: authservice.NewAuthService(zrpc.MustNewClient(c.AuthRpc)),
		MsgRpc:      chat.NewChat(zrpc.MustNewClient(c.MsgRpc)),
	}
}
