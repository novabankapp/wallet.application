package internal

import (
	es "github.com/novabankapp/common.data/eventstore"
	"github.com/shopspring/decimal"
)

type WalletDetails struct {
	Amount      decimal.Decimal
	Description string
	UserId      string
	AccountId   string
	Id          string
}
type CreateWalletCommand struct {
	es.BaseCommand
	WalletDetails WalletDetails
}

func NewCreateWalletCommand(aggregateID string, amount decimal.Decimal, description, userId, accountId, id string) *CreateWalletCommand {
	return &CreateWalletCommand{BaseCommand: es.NewBaseCommand(aggregateID), WalletDetails: WalletDetails{
		amount,
		description,
		userId,
		accountId,
		id,
	}}
}
