package commands

import (
	"context"
	"errors"
	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/novabankapp/common.application/services/message_queue"
	es "github.com/novabankapp/common.data/eventstore"
	kafkaClient "github.com/novabankapp/common.infrastructure/kafka"
	"github.com/novabankapp/common.infrastructure/logger"
	"github.com/novabankapp/wallet.data/constants"
	"github.com/novabankapp/wallet.data/es/aggregate"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type DeleteWalletCommandHandler interface {
	Handle(ctx context.Context, command *DeleteWalletCommand) error
}

type deleteWalletHandler struct {
	log          logger.Logger
	es           es.AggregateStore
	topics       *kafkaClient.KafkaTopics
	messageQueue message_queue.MessageQueue
}

func NewDeleteWalletHandler(log logger.Logger,
	es es.AggregateStore,
	topics *kafkaClient.KafkaTopics,
	messageQueue message_queue.MessageQueue) *deleteWalletHandler {
	return &deleteWalletHandler{log: log, es: es, topics: topics, messageQueue: messageQueue}
}

func (c *deleteWalletHandler) Handle(ctx context.Context, command *DeleteWalletCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "deleteWalletHandler.Handle")
	defer span.Finish()
	span.LogFields(log.String(constants.AggregateID, command.GetAggregateID()))

	wallet := aggregate.NewWalletAggregateWithID(command.AggregateID)
	err := c.es.Exists(ctx, wallet.GetID())
	if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
		return err
	}

	if err := wallet.DeleteWallet(
		ctx,
		command.Description,
	); err != nil {
		return err
	}

	/*res := new(bytes.Buffer)
	e := json.NewEncoder(res).Encode(userDto)
	if e == nil {
		msgBytes := res.Bytes()

		_, err2 = c.messageQueue.PublishMessage(ctx, msgBytes, *userId, r.topics.UserCreated.TopicName)
	}*/

	span.LogFields(log.String("wallet", wallet.String()))
	return c.es.Save(ctx, wallet)
}
