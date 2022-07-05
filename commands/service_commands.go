package commands

type WalletCommands struct {
	CreateWallet         CreateWalletCommandHandler
	LockWalletCommand    LockWalletCommandHandler
	UnlockWalletCommand  UnlockWalletCommandHandler
	BlockWalletCommand   BlockWalletCommandHandler
	UnblockWalletCommand UnBlockWalletCommandHandler
	DeleteWalletCommand  DeleteWalletCommandHandler
	DebitWalletCommand   DebitWalletCommandHandler
	CreditWalletCommand  CreditWalletCommandHandler
}

func NewWalletCommands(
	createWallet CreateWalletCommandHandler,
	lockWallet LockWalletCommandHandler,
	unlockWallet UnlockWalletCommandHandler,
	blockWallet BlockWalletCommandHandler,
	unblockWallet UnBlockWalletCommandHandler,
	deleteWalletCommand DeleteWalletCommandHandler,
	debitWalletCommand DebitWalletCommandHandler,
	creditWalletCommand CreditWalletCommandHandler,

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
