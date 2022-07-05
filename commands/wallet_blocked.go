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

type BlockWalletCommandHandler interface {
	Handle(ctx context.Context, command *BlockWalletCommand) error
}

type blockWalletHandler struct {
	log logger.Logger
	cfg *config.Config
	es  es.AggregateStore
}

func NewBlockWalletHandler(log logger.Logger, cfg *config.Config, es es.AggregateStore) *blockWalletHandler {
	return &blockWalletHandler{log: log, cfg: cfg, es: es}
}

func (c *blockWalletHandler) Handle(ctx context.Context, command *BlockWalletCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "blockWalletHandler.Handle")
	defer span.Finish()
	span.LogFields(log.String(constants.WalletID, command.GetAggregateID()))

	wallet := aggregate.NewWalletAggregateWithID(command.AggregateID)
	err := c.es.Exists(ctx, wallet.GetID())
	if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
		return err
	}

	if err := wallet.BlacklistWallet(
		ctx,
		command.Description,
	); err != nil {
		return err
	}

	span.LogFields(log.String("wallet", wallet.String()))
	return c.es.Save(ctx, wallet)
}
