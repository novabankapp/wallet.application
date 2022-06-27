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

type CreateWalletCommandHandler interface {
	Handle(ctx context.Context, command *internal.CreateWalletCommand) error
}

type createWalletHandler struct {
	log logger.Logger
	cfg *config.Config
	es  es.AggregateStore
}

func NewCreateWalletHandler(log logger.Logger, cfg *config.Config, es es.AggregateStore) *createWalletHandler {
	return &createWalletHandler{log: log, cfg: cfg, es: es}
}

func (c *createWalletHandler) Handle(ctx context.Context, command *internal.CreateWalletCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "createOrderHandler.Handle")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", command.GetAggregateID()))

	wallet := aggregate.NewWalletAggregateWithID(command.AggregateID)
	err := c.es.Exists(ctx, wallet.GetID())
	if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
		return err
	}

	if err := wallet.CreateWallet(
		ctx,
		command.WalletDetails.Amount,
		command.WalletDetails.Description,
		command.WalletDetails.UserId,
		command.WalletDetails.AccountId,
		command.WalletDetails.Id,

	); err != nil {
		return err
	}

	span.LogFields(log.String("order", wallet.String()))
	return c.es.Save(ctx, wallet)
}
