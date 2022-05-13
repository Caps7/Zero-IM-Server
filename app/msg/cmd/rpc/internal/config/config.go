package config

import (
	"github.com/showurl/Zero-IM-Server/common/xkafka"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Kafka           KafkaConfig
	Callback        CallbackConfig
	MessageVerify   MessageVerifyConfig
	ImUserRpc       zrpc.RpcClientConf
	ConversationRpc zrpc.RpcClientConf
	MsgCallbackRpc  zrpc.RpcClientConf
}
type CallbackConfig struct {
	CallbackWordFilter          CallbackConfigItem
	CallbackBeforeSendGroupMsg  CallbackConfigItem
	CallbackAfterSendGroupMsg   CallbackConfigItem
	CallbackBeforeSendSingleMsg CallbackConfigItem
	CallbackAfterSendSingleMsg  CallbackConfigItem
}
type CallbackConfigItem struct {
	Enable          bool
	ContinueOnError bool
}
type MessageVerifyConfig struct {
	FriendVerify bool // 只有好友才能发送消息
}
type KafkaConfig struct {
	Online  xkafka.ProducerConfig
	Offline xkafka.ProducerConfig
}
