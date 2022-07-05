package commands

import (
	"context"
	"errors"
	"github.com/EventStore/EventStore-Client-Go/esdb"
	es "github.com/novabankapp/common.data/eventstore"
	"github.com/novabankapp/common.infrastructure/logger"
	"github.com/novabankapp/wallet.data/constants"
	"github.com/novabankapp/wallet.data/es/aggregate"
	"github.com/olivere/elastic/v7/config"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type DeleteWalletCommandHandler interface {
	Handle(ctx context.Context, command *DeleteWalletCommand) error
}

type deleteWalletHandler struct {
	log logger.Logger
	cfg *config.Config
	es  es.AggregateStore
}

func NewDeleteWalletHandler(log logger.Logger, cfg *config.Config, es es.AggregateStore) *deleteWalletHandler {
	return &deleteWalletHandler{log: log, cfg: cfg, es: es}
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

	span.LogFields(log.String("wallet", wallet.String()))
	return c.es.Save(ctx, wallet)
}
