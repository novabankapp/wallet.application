package queries

type WalletQueries struct {
	GetWalletByID      GetWalletByIDQueryHandler
	GetUserWalletsByID GetUserWalletsByIDQueryHandler
}

func NewWalletQueries(getWalletByID GetWalletByIDQueryHandler, getUserWallets GetUserWalletsByIDQueryHandler) *WalletQueries {
	return &WalletQueries{
		GetWalletByID:      getWalletByID,
		GetUserWalletsByID: getUserWallets,
	}
}

type GetWalletByIDQuery struct {
	ID string
}

type GetUserWalletsByIDQuery struct {
	UserID string
}
