package queries

import (
	"context"
	es "github.com/novabankapp/common.data/eventstore"
	"github.com/novabankapp/common.data/repositories/base"
	"github.com/novabankapp/common.infrastructure/logger"
	"github.com/novabankapp/wallet.data/es/models"
	"github.com/olivere/elastic/v7/config"
	"github.com/opentracing/opentracing-go"
)

type GetUserWalletsByIDQueryHandler interface {
	Handle(ctx context.Context, command *GetWalletByIDQuery) (*models.WalletProjection, error)
}

type getUserWalletsByIDQueryHandler struct {
	log  logger.Logger
	cfg  *config.Config
	es   es.AggregateStore
	repo base.NoSqlRepository[models.WalletProjection]
}

func NewGetUserWalletsByIDQueryHandler(log logger.Logger,
	cfg *config.Config, es es.AggregateStore,
	repo base.NoSqlRepository[models.WalletProjection]) *getUserWalletsByIDQueryHandler {
	return &getUserWalletsByIDQueryHandler{log: log, cfg: cfg, es: es, repo: repo}
}

func (q *getUserWalletsByIDQueryHandler) Handle(ctx context.Context, query *GetUserWalletsByIDQuery, pageSize int, page []byte) (results *[]models.WalletProjection, pageState []byte, error error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "getUserWalletsByIDHandler.Handle")
	defer span.Finish()

	queries := make([]map[string]string, 1)
	m := make(map[string]string)
	m["column"] = "UserId"
	m["compare"] = "="
	m["value"] = query.UserID
	queries = append(queries, m)

	return q.repo.Get(ctx, page, pageSize, queries)

}