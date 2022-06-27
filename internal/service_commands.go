package internal

import "github.com/novabankapp/wallet.application/commands"

type WalletCommands struct {
	CreateWallet commands.CreateWalletCommandHandler
}

func NewWalletCommands(
	createWallet commands.CreateWalletCommandHandler,

) *WalletCommands {
	return &WalletCommands{
		CreateWallet: createWallet,
	}
}
