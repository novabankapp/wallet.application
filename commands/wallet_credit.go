package commands

import (
	"context"
	"errors"
	"github.com/EventStore/EventStore-Client-Go/esdb"
	es "github.com/novabankapp/common.data/eventstore"
	"github.com/novabankapp/common.infrastructure/logger"
	"github.com/novabankapp/wallet.application/internal"
	"github.com/novabankapp/wallet.data/es/aggregate"
	"github.com/olivere/elastic/v7/config"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type CreditWalletCommandHandler interface {
	Handle(ctx context.Context, command *internal.CreditWalletCommand) error
}

type creditWalletHandler struct {
	log logger.Logger
	cfg *config.Config
	es  es.AggregateStore
}

func NewCreditWalletHandler(log logger.Logger, cfg *config.Config, es es.AggregateStore) *creditWalletHandler {
	return &creditWalletHandler{log: log, cfg: cfg, es: es}
}

func (c *creditWalletHandler) Handle(ctx context.Context, command *internal.CreditWalletCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "creditWalletHandler.Handle")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", command.GetAggregateID()))

	wallet := aggregate.NewWalletAggregateWithID(command.AggregateID)
	err := c.es.Exists(ctx, wallet.GetID())
	if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
		return err
	}

	if err := wallet.CreditWallet(
		ctx,
		command.DebitWalletID,
		command.Amount,
		command.Description,
	); err != nil {
		return err
	}

	span.LogFields(log.String("wallet", wallet.String()))
	return c.es.Save(ctx, wallet)
}
