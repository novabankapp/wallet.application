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

type DebitWalletCommandHandler interface {
	Handle(ctx context.Context, command *DebitWalletCommand) error
}

type debitWalletHandler struct {
	log logger.Logger
	es  es.AggregateStore
}

func NewDebitWalletHandler(log logger.Logger, es es.AggregateStore) *debitWalletHandler {
	return &debitWalletHandler{log: log, es: es}
}

func (c *debitWalletHandler) Handle(ctx context.Context, command *DebitWalletCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "debitWalletHandler.Handle")
	defer span.Finish()
	span.LogFields(log.String(constants.AggregateID, command.GetAggregateID()))

	wallet := aggregate.NewWalletAggregateWithID(command.AggregateID)
	err := c.es.Exists(ctx, wallet.GetID())
	if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
		return err
	}
	creditWallet := aggregate.NewWalletAggregateWithID(command.CreditWalletAggregateID)
	err = c.es.Exists(ctx, creditWallet.GetID())
	if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
		return err
	}

	if err := wallet.DebitWallet(
		ctx,
		creditWallet.Wallet.ID.String(),
		command.Amount,
		command.Description,
	); err != nil {
		return err
	}

	span.LogFields(log.String("wallet", wallet.String()))
	return c.es.Save(ctx, wallet)
}
