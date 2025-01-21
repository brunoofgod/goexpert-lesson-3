package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	List(page, limit int) ([]Order, int, error)
}
