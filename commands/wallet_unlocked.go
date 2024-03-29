package commands

import (
	"context"
	"errors"
	"github.com/EventStore/EventStore-Client-Go/esdb"
	es "github.com/novabankapp/common.data/eventstore"
	"github.com/novabankapp/common.infrastructure/logger"
	"github.com/novabankapp/wallet.data/constants"
	"github.com/novabankapp/wallet.data/es/aggregate"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type UnlockWalletCommandHandler interface {
	Handle(ctx context.Context, command *UnlockWalletCommand) error
}

type unlockWalletHandler struct {
	log logger.Logger
	es  es.AggregateStore
}

func NewUnlockWalletHandler(log logger.Logger, es es.AggregateStore) *unlockWalletHandler {
	return &unlockWalletHandler{log: log, es: es}
}

func (c *unlockWalletHandler) Handle(ctx context.Context, command *UnlockWalletCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "unlockWalletHandler.Handle")
	defer span.Finish()
	span.LogFields(log.String(constants.AggregateID, command.GetAggregateID()))

	wallet := aggregate.NewWalletAggregateWithID(command.AggregateID)
	err := c.es.Exists(ctx, wallet.GetID())
	if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
		return err
	}

	if err := wallet.UnlockWallet(
		ctx,
		command.Description,
	); err != nil {
		return err
	}

	span.LogFields(log.String("wallet", wallet.String()))
	return c.es.Save(ctx, wallet)
}
