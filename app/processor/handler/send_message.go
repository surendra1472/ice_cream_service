package handler

import (
	"context"
	"encoding/json"
	"ic-service/app/builder"
	"ic-service/app/config"
	"ic-service/app/model/bo"
	"ic-service/kafka_producers"
)

//go:generate sh -c "$GOPATH/bin/mockery -case=underscore -dir=. -name=SendMessageToKafkaHandler"
type SendMessageToKafkaHandler interface {
	SendMessage(ctx context.Context, icecream *bo.Icecream, isDeleted bool) error
}

type sendMessageToKafkaHandler struct {
	syncProducer  kafka_producers.SyncProducer
	icecreamTopic string
}

func NewSendMessageToKafkaHandler() *sendMessageToKafkaHandler {
	return &sendMessageToKafkaHandler{
		syncProducer:  config.GetSyncProducer(),
		icecreamTopic: config.GetConfig().Kafka.IceCreamCreateTopic,
	}
}

func (smtkh sendMessageToKafkaHandler) SendMessage(ctx context.Context, icecream *bo.Icecream, isDeleted bool) error {

	indexerBuilder := builder.NewIcecreamIndexerBuilder().IcecreamIndexerBuilder(icecream, isDeleted)
	val, _ := json.Marshal(indexerBuilder)
	message := kafka_producers.Message{Topic: smtkh.icecreamTopic, Value: val}
	return smtkh.syncProducer.SendMessage(ctx, message)
}
