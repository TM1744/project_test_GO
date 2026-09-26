package enum

type Role string

const (
	Admin      Role = "ADMIN"
	Manager    Role = "MANAGER"
	Seller     Role = "SELLER"
	StockClerk Role = "STOCK_CLERK"
)

func (r Role) IsValid () bool {
	switch r {
	case Admin, Manager, Seller, StockClerk:
		return true
	}
	return false
}
