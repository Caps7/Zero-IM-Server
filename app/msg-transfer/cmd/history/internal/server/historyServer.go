package server

import (
	"context"
	"errors"
	"github.com/Shopify/sarama"
	"github.com/showurl/Zero-IM-Server/app/msg-transfer/cmd/history/internal/logic"
	"github.com/showurl/Zero-IM-Server/app/msg-transfer/cmd/history/internal/svc"
	"github.com/showurl/Zero-IM-Server/common/utils"
	"github.com/showurl/Zero-IM-Server/common/xkafka"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/trace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	oteltrace "go.opentelemetry.io/otel/trace"
	"net/http"
	"sync"
	"time"
)

func NewMsgTransferHistoryServer(svcCtx *svc.ServiceContext) *MsgTransferHistoryServer {
	m := &MsgTransferHistoryServer{svcCtx: svcCtx}
	m.cmdCh = make(chan Cmd2Value, 10000)
	m.w = new(sync.Mutex)
	m.msgHandle = make(map[string]fcb)
	m.msgHandle[svcCtx.Config.Kafka.Online.Topic] = m.ChatMs2Mongo
	m.historyConsumerGroup = xkafka.NewMConsumerGroup(&xkafka.MConsumerGroupConfig{
		KafkaVersion:   sarama.V0_10_2_0,
		OffsetsInitial: sarama.OffsetNewest, IsReturnErr: false,
	}, []string{svcCtx.Config.Kafka.Online.Topic},
		svcCtx.Config.Kafka.Online.Brokers, svcCtx.Config.Kafka.Online.MsgToMongoGroupID)
	return m
}

func (s *MsgTransferHistoryServer) Start() {
	s.historyConsumerGroup.RegisterHandleAndConsumer(s)
}

func (s *MsgTransferHistoryServer) runWithCtx(f func(ctx context.Context), kv ...attribute.KeyValue) {
	propagator := otel.GetTextMapPropagator()
	tracer := otel.GetTracerProvider().Tracer(trace.TraceName)
	ctx := propagator.Extract(context.Background(), propagation.HeaderCarrier(http.Header{}))
	spanName := utils.CallerFuncName()
	spanCtx, span := tracer.Start(
		ctx,
		spanName,
		oteltrace.WithSpanKind(oteltrace.SpanKindServer),
		oteltrace.WithAttributes(kv...),
	)
	defer span.End()
	propagator.Inject(spanCtx, propagation.HeaderCarrier(http.Header{}))
	f(spanCtx)
}

func (s *MsgTransferHistoryServer) ChatMs2Mongo(msg []byte, msgKey string) {
	s.runWithCtx(func(ctx context.Context) {
		logic.NewMsgTransferHistoryOnlineLogic(ctx, s.svcCtx).ChatMs2Mongo(msg, msgKey)
	}, attribute.String("msgKey", msgKey))
}

func (s *MsgTransferHistoryServer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		s.SetOnlineTopicStatus(OnlineTopicBusy)
		s.msgHandle[msg.Topic](msg.Value, string(msg.Key))
		sess.MarkMessage(msg, "")
		if claim.HighWaterMarkOffset()-msg.Offset <= 1 {
			s.SetOnlineTopicStatus(OnlineTopicVacancy)
			s.TriggerCmd(context.Background(), OnlineTopicVacancy)
		}
	}
	return nil
}

func (s *MsgTransferHistoryServer) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (s *MsgTransferHistoryServer) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (s *MsgTransferHistoryServer) SetOnlineTopicStatus(busy int) {
	s.w.Lock()
	defer s.w.Unlock()
	s.OnlineTopicStatus = busy

}
func (s *MsgTransferHistoryServer) TriggerCmd(ctx context.Context, status int) {
	for {
		err := s.sendCmd(ctx, s.cmdCh, Cmd2Value{Cmd: status, Value: ""}, 1)
		if err != nil {
			logx.WithContext(ctx).Errorf("send cmd error: %v", err)
			continue
		}
		return
	}
}
func (s *MsgTransferHistoryServer) sendCmd(ctx context.Context, ch chan Cmd2Value, value Cmd2Value, timeout int64) error {
	var flag = 0
	select {
	case ch <- value:
		flag = 1
	case <-time.After(time.Second * time.Duration(timeout)):
		flag = 2
	}
	if flag == 1 {
		return nil
	} else {
		return errors.New("send cmd timeout")
	}
}
