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

func NewCreateWalletCommand(aggregateID string, amount decimal.Decimal,
	description, userId, accountId, id string) *CreateWalletCommand {
	return &CreateWalletCommand{BaseCommand: es.NewBaseCommand(aggregateID), WalletDetails: WalletDetails{
		amount,
		description,
		userId,
		accountId,
		id,
	}}
}

type LockWalletCommand struct {
	es.BaseCommand
	WalletID    string
	Description string
}

func NewLockWalletCommand(aggregateID, walletID string, description string) *LockWalletCommand {
	return &LockWalletCommand{
		BaseCommand: es.NewBaseCommand(aggregateID),
		WalletID:    walletID,
		Description: description,
	}
}

type UnlockWalletCommand struct {
	es.BaseCommand
	WalletID    string
	Description string
}

func NewUnlockWalletCommand(aggregateID, walletID string, description string) *UnlockWalletCommand {
	return &UnlockWalletCommand{
		BaseCommand: es.NewBaseCommand(aggregateID),
		WalletID:    walletID,
		Description: description,
	}
}

type BlockWalletCommand struct {
	es.BaseCommand
	WalletID    string
	Description string
}

func NewBlockWalletCommand(aggregateID, walletID string, description string) *BlockWalletCommand {
	return &BlockWalletCommand{
		BaseCommand: es.NewBaseCommand(aggregateID),
		WalletID:    walletID,
		Description: description,
	}
}

type UnblockWalletCommand struct {
	es.BaseCommand
	WalletID    string
	Description string
}

func NewUnblockWalletCommand(aggregateID, walletID string, description string) *UnblockWalletCommand {
	return &UnblockWalletCommand{
		BaseCommand: es.NewBaseCommand(aggregateID),
		WalletID:    walletID,
		Description: description,
	}
}

type DeleteWalletCommand struct {
	es.BaseCommand
	WalletID    string
	Description string
}

func NewDeleteWalletCommand(aggregateID, walletID string, description string) *DeleteWalletCommand {
	return &DeleteWalletCommand{
		BaseCommand: es.NewBaseCommand(aggregateID),
		WalletID:    walletID,
		Description: description,
	}
}

type DebitWalletCommand struct {
	es.BaseCommand
	CreditWalletID string
	Amount         decimal.Decimal
	Description    string
}

func NewDebitWalletCommand(aggregateID, creditWalletID string, amount decimal.Decimal, description string) *DebitWalletCommand {
	return &DebitWalletCommand{
		BaseCommand:    es.NewBaseCommand(aggregateID),
		CreditWalletID: creditWalletID,
		Amount:         amount,
		Description:    description,
	}
}

type CreditWalletCommand struct {
	es.BaseCommand
	DebitWalletID string
	Amount        decimal.Decimal
	Description   string
}

func NewCreditWalletCommand(aggregateID, debitWalletID string, amount decimal.Decimal, description string) *CreditWalletCommand {
	return &CreditWalletCommand{
		BaseCommand:   es.NewBaseCommand(aggregateID),
		DebitWalletID: debitWalletID,
		Amount:        amount,
		Description:   description,
	}
}
