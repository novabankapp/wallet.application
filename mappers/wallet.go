package mappers

import (
	"github.com/novabankapp/wallet.data/es/aggregate"
	"github.com/novabankapp/wallet.data/es/models"
)

func WalletProjectionFromAggregate(walletAggregate *aggregate.WalletAggregate) *models.WalletProjection {
	return &models.WalletProjection{
		WalletID:           aggregate.GetWalletAggregateID(walletAggregate.GetID()),
		ID:                 walletAggregate.ID,
		Wallet:             *walletAggregate.Wallet,
		WalletState:        *walletAggregate.WalletState,
		WalletTransactions: *walletAggregate.WalletTransactions,
	}
}
