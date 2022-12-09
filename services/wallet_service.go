package services

import (
	es "github.com/novabankapp/common.data/eventstore"
	"github.com/novabankapp/common.data/repositories/base"
	"github.com/novabankapp/common.infrastructure/logger"
	"github.com/novabankapp/wallet.application/commands"
	"github.com/novabankapp/wallet.application/queries"
	"github.com/novabankapp/wallet.data/es/models"
)

type WalletService struct {
	Commands *commands.WalletCommands
	Queries  *queries.WalletQueries
}

func NewWalletService(
	log logger.Logger,
	es es.AggregateStore,
	repo base.NoSqlRepository[models.WalletProjection],
	//elasticRepository repository.ElasticOrderRepository,
) *WalletService {

	createWalletHandler := commands.NewCreateWalletHandler(log, es)
	lockWalletHandler := commands.NewLockWalletHandler(log, es)
	unlockWalletHandler := commands.NewUnlockWalletHandler(log, es)
	blockWalletHandler := commands.NewBlockWalletHandler(log, es)
	unblockWalletHandler := commands.NewUnblockWalletHandler(log, es)
	deleteWalletHandler := commands.NewDeleteWalletHandler(log, es)
	debitWalletHandler := commands.NewDebitWalletHandler(log, es)
	creditWalletHandler := commands.NewCreditWalletHandler(log, es)

	getWalletByIDHandler := queries.NewGetWalletByIDHandler(log, es, repo)
	getUserWalletsHandler := queries.NewGetUserWalletsByIDQueryHandler(log, es, repo)

	walletCommands := commands.NewWalletCommands(
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
