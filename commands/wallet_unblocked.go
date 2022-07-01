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

type UnBlockWalletCommandHandler interface {
	Handle(ctx context.Context, command *internal.UnblockWalletCommand) error
}

type unBlockWalletHandler struct {
	log logger.Logger
	cfg *config.Config
	es  es.AggregateStore
}

func NewUnblockWalletHandler(log logger.Logger, cfg *config.Config, es es.AggregateStore) *unBlockWalletHandler {
	return &unBlockWalletHandler{log: log, cfg: cfg, es: es}
}

func (c *unBlockWalletHandler) Handle(ctx context.Context, command *internal.UnblockWalletCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "unblockWalletHandler.Handle")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", command.GetAggregateID()))

	wallet := aggregate.NewWalletAggregateWithID(command.AggregateID)
	err := c.es.Exists(ctx, wallet.GetID())
	if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
		return err
	}

	if err := wallet.UnBlacklistWallet(
		ctx,
		command.Description,
	); err != nil {
		return err
	}

	span.LogFields(log.String("wallet", wallet.String()))
	return c.es.Save(ctx, wallet)
}
