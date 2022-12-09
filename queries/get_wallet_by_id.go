package queries

import (
	"context"
	cons "github.com/novabankapp/common.data/constants"
	es "github.com/novabankapp/common.data/eventstore"
	"github.com/novabankapp/common.data/repositories/base"
	"github.com/novabankapp/common.infrastructure/logger"
	"github.com/novabankapp/wallet.application/internal/dtos"
	"github.com/novabankapp/wallet.application/mappers"
	"github.com/novabankapp/wallet.data/constants"
	"github.com/novabankapp/wallet.data/es/aggregate"
	"github.com/novabankapp/wallet.data/es/models"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type GetWalletByIDQueryHandler interface {
	Handle(ctx context.Context, command *GetWalletByIDQuery) (*dtos.WalletProjectionDto, error)
}

type getWalletByIDHandler struct {
	log  logger.Logger
	es   es.AggregateStore
	repo base.NoSqlRepository[models.WalletProjection]
}

func NewGetWalletByIDHandler(log logger.Logger,
	es es.AggregateStore,
	repo base.NoSqlRepository[models.WalletProjection]) *getWalletByIDHandler {
	return &getWalletByIDHandler{log: log, es: es, repo: repo}
}

func (q *getWalletByIDHandler) Handle(ctx context.Context, query *GetWalletByIDQuery) (*dtos.WalletProjectionDto, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "getWalletByIDHandler.Handle")
	defer span.Finish()
	span.LogFields(log.String(constants.AggregateID, query.ID))
	var walletProjection *models.WalletProjection
	queries := make([]map[string]string, 1)
	m := make(map[string]string)
	m[cons.Column] = constants.WalletID
	m[cons.Compare] = cons.Equal
	m[cons.Value] = query.ID
	queries = append(queries, m)
	walletProjection, err := q.repo.GetByCondition(ctx, queries)
	if err != nil {
		return nil, err
	}
	if walletProjection != nil {
		dto := mappers.WalletProjectionDtoFromProjection(*walletProjection)
		return &dto, nil
	}

	walletAgg := aggregate.NewWalletAggregateWithID(query.ID)
	if err := q.es.Load(ctx, walletAgg); err != nil {
		return nil, err
	}

	if aggregate.IsAggregateNotFound(walletAgg) {
		return nil, aggregate.ErrOrderNotFound
	}

	walletProjection = mappers.WalletProjectionFromAggregate(walletAgg)

	_, err = q.repo.Create(ctx, *walletProjection)
	if err != nil {
		return nil, err
	}

	dto := mappers.WalletProjectionDtoFromProjection(*walletProjection)
	return &dto, nil
}
