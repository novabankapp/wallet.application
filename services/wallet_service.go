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
	lockWalletHandler := commands.NewLockWalletHandler(log, cfg, es)
	unlockWalletHandler := commands.NewUnlockWalletHandler(log, cfg, es)
	blockWalletHandler := commands.NewBlockWalletHandler(log, cfg, es)
	unblockWalletHandler := commands.NewUnblockWalletHandler(log, cfg, es)
	deleteWalletHandler := commands.NewDeleteWalletHandler(log, cfg, es)
	debitWalletHandler := commands.NewDebitWalletHandler(log, cfg, es)
	creditWalletHandler := commands.NewCreditWalletHandler(log, cfg, es)

	getWalletByIDHandler := queries.NewGetWalletByIDHandler(log, cfg, es, repo)
	getUserWalletsHandler := queries.NewGetUserWalletsByIDQueryHandler(log, cfg, es, repo)

	walletCommands := internal.NewWalletCommands(
		createWalletHandler,
		lockWalletHandler,
		unlockWalletHandler,
		blockWalletHandler,
		unblockWalletHandler,
		deleteWalletHandler,
		debitWalletHandler,
		creditWalletHandler,
	)

	walletQueries := queries.NewWalletQueries(
		getWalletByIDHandler,
		getUserWalletsHandler,
	)

	return &WalletService{Commands: walletCommands, Queries: walletQueries}
}
