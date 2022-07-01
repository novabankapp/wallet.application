package internal

import "github.com/novabankapp/wallet.application/commands"

type WalletCommands struct {
	CreateWallet         commands.CreateWalletCommandHandler
	LockWalletCommand    commands.LockWalletCommandHandler
	UnlockWalletCommand  commands.UnlockWalletCommandHandler
	BlockWalletCommand   commands.BlockWalletCommandHandler
	UnblockWalletCommand commands.UnBlockWalletCommandHandler
	DeleteWalletCommand  commands.DeleteWalletCommandHandler
	DebitWalletCommand   commands.DebitWalletCommandHandler
	CreditWalletCommand  commands.CreditWalletCommandHandler
}

func NewWalletCommands(
	createWallet commands.CreateWalletCommandHandler,
	lockWallet commands.LockWalletCommandHandler,
	unlockWallet commands.UnlockWalletCommandHandler,
	blockWallet commands.BlockWalletCommandHandler,
	unblockWallet commands.UnBlockWalletCommandHandler,
	deleteWalletCommand commands.DeleteWalletCommandHandler,
	debitWalletCommand commands.DebitWalletCommandHandler,
	creditWalletCommand commands.CreditWalletCommandHandler,

) *WalletCommands {
	return &WalletCommands{
		CreateWallet:         createWallet,
		LockWalletCommand:    lockWallet,
		UnlockWalletCommand:  unlockWallet,
		BlockWalletCommand:   blockWallet,
		UnblockWalletCommand: unblockWallet,
		DeleteWalletCommand:  deleteWalletCommand,
		DebitWalletCommand:   debitWalletCommand,
		CreditWalletCommand:  creditWalletCommand,
	}
}
