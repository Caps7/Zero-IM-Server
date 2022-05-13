package wsconfig

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	AuthRpc   zrpc.RpcClientConf
	MsgRpc    zrpc.RpcClientConf
	Websocket WebsocketConfig
}
type WebsocketConfig struct {
	MaxConnNum int
	TimeOut    int
	MaxMsgLen  int
}
