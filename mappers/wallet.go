package mappers

import (
	"github.com/novabankapp/wallet.application/internal/dtos"
	"github.com/novabankapp/wallet.data/domain"
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
func WalletProjectionDtoFromProjection(walletProjection models.WalletProjection) dtos.WalletProjectionDto {
	walletStateP, _ := aggregate.GetEntityFromJsonString[domain.WalletState](walletProjection.WalletState)
	var walletState = *walletStateP
	walletP, _ := aggregate.GetEntityFromJsonString[domain.Wallet](walletProjection.Wallet)
	var wallet = *walletP
	walletTransactionsP, _ := aggregate.GetEntityArrayFromJsonString[domain.WalletTransaction](walletProjection.WalletTransactions)
	var walletTransactions []domain.WalletTransaction = *walletTransactionsP
	return dtos.WalletProjectionDto{
		ID:                 walletProjection.ID,
		WalletID:           walletProjection.WalletID,
		WalletState:        walletState,
		Wallet:             wallet,
		WalletTransactions: walletTransactions,
	}
}
