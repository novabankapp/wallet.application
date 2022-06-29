package mappers

import (
	"github.com/novabankapp/wallet.data/es/aggregate"
	"github.com/novabankapp/wallet.data/es/models"
)

func WalletProjectionFromAggregate(walletAggregate *aggregate.WalletAggregate) *models.WalletProjection {
	wallet := aggregate.GetJsonString(*walletAggregate.Wallet)
	walletState := aggregate.GetJsonString(*walletAggregate.WalletState)
	walletTransactions := aggregate.GetJsonString(*walletAggregate.WalletTransactions)
	return &models.WalletProjection{
		WalletID:           aggregate.GetWalletAggregateID(walletAggregate.GetID()),
		ID:                 walletAggregate.ID,
		Wallet:             wallet,
		WalletState:        walletState,
		WalletTransactions: walletTransactions,
	}
}
