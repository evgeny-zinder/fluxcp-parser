package domain

const (
	_                     = iota
	CurrencyZeny Currency = iota
	CurrencyCash
)

type Currency int

type Position struct {
	Map string
	X   int
	Y   int
}

type Seller struct {
	Name     string
	ShopName string
	Position Position
}

type Item struct {
	ID    int
	Name  string
	Card0 string
	Card1 string
	Card2 string
	Card3 string
}

type Order struct {
	Item     Item
	Price    int64
	Currency Currency
	Amount   int
}

type Shop struct {
	Orders   []Order
	Currency Currency
	Seller   Seller
}

type MarketStats struct {
	TotalOrders   int
	TotalPages    int
	OrdersPerPage int
}
