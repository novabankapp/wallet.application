package queries

type WalletQueries struct {
	GetWalletByID      GetWalletByIDQueryHandler
	GetUserWalletsByID GetUserWalletsByIDQueryHandler
}

func NewWalletQueries(getWalletByID GetWalletByIDQueryHandler) *WalletQueries {
	return &WalletQueries{
		GetWalletByID: getWalletByID,
	}
}

type GetWalletByIDQuery struct {
	ID string
}

type GetUserWalletsByIDQuery struct {
	UserID string
}
