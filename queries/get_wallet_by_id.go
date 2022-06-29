package queries

import (
	"context"
	es "github.com/novabankapp/common.data/eventstore"
	"github.com/novabankapp/common.data/repositories/base"
	"github.com/novabankapp/common.infrastructure/logger"
	"github.com/novabankapp/wallet.application/mappers"
	"github.com/novabankapp/wallet.data/es/aggregate"
	"github.com/novabankapp/wallet.data/es/models"
	"github.com/olivere/elastic/v7/config"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type GetWalletByIDQueryHandler interface {
	Handle(ctx context.Context, command *GetWalletByIDQuery) (*models.WalletProjection, error)
}

type getWalletByIDHandler struct {
	log  logger.Logger
	cfg  *config.Config
	es   es.AggregateStore
	repo base.NoSqlRepository[models.WalletProjection]
}

func NewGetWalletByIDHandler(log logger.Logger, cfg *config.Config, es es.AggregateStore, repo base.NoSqlRepository[models.WalletProjection]) *getWalletByIDHandler {
	return &getWalletByIDHandler{log: log, cfg: cfg, es: es, repo: repo}
}

func (q *getWalletByIDHandler) Handle(ctx context.Context, query *GetWalletByIDQuery) (*models.WalletProjection, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "getWalletByIDHandler.Handle")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", query.ID))

	walletProjection, err := q.repo.GetById(ctx, query.ID)
	if err != nil {
		return nil, err
	}
	if walletProjection != nil {
		return walletProjection, nil
	}

	wallet := aggregate.NewWalletAggregateWithID(query.ID)
	if err := q.es.Load(ctx, wallet); err != nil {
		return nil, err
	}

	if aggregate.IsAggregateNotFound(wallet) {
		return nil, aggregate.ErrOrderNotFound
	}

	walletProjection = mappers.WalletProjectionFromAggregate(wallet)

	_, err = q.repo.Create(ctx, *walletProjection)
	if err != nil {
		return nil, err
	}

	return walletProjection, nil
}
