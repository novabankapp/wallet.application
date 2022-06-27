package services

import (
	es "github.com/novabankapp/common.data/eventstore"
	"github.com/novabankapp/common.data/repositories/base"
	"github.com/novabankapp/common.infrastructure/logger"
	"github.com/novabankapp/wallet.application/commands"
	"github.com/novabankapp/wallet.application/internal"
	"github.com/novabankapp/wallet.application/queries"
	"github.com/novabankapp/wallet.data/es/models"
	"github.com/olivere/elastic/v7/config"
)

type WalletService struct {
	Commands *internal.WalletCommands
	Queries  *queries.WalletQueries
}

func NewWalletService(
	log logger.Logger,
	cfg *config.Config,
	es es.AggregateStore,
	repo base.NoSqlRepository[models.WalletProjection],
//elasticRepository repository.ElasticOrderRepository,
) *WalletService {

	createWalletHandler := commands.NewCreateWalletHandler(log, cfg, es)

	//getOrderByIDHandler := queries.NewGetOrderByIDHandler(log, cfg, es, mongoRepo)
	//searchOrdersHandler := queries.NewSearchOrdersHandler(log, cfg, es, elasticRepository)

	walletCommands := internal.NewWalletCommands(
		createWalletHandler,
	)
	orderQueries := queries.NewWalletQueries( /*getOrderByIDHandler, searchOrdersHandler*/)

	return &WalletService{Commands: walletCommands, Queries: orderQueries}
}
