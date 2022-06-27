package queries

type WalletQueries struct {
	GetWalletByID GetWalletByIDQueryHandler
	//SearchOrders SearchOrdersQueryHandler
}

func NewWalletQueries(getWalletByID GetWalletByIDQueryHandler) *WalletQueries {
	return &WalletQueries{
		GetWalletByID: getWalletByID,
	}
}

type GetWalletByIDQuery struct {
	ID string
}
