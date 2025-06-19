package order

import "github.com/tiketin-management-api-with-go/structs"

type OrderRepositoryInterface interface {
	CreateOrder(userId int, order structs.OrderCreate) error
	PayOrder(id int, order structs.Order) error
	CancelOrder(id int) error
	GetAllOrder() ([]structs.Order, error)
	GetOrderByUser(userId int) ([]structs.Order, error)
	GetOrderById(id int) (structs.Order, error)
	CheckIn(id int) error
	IsAllOrderItemCheckedIn(orderId int) (bool, error)
	UpdateOrderStatus(orderId int, status string) error
	IsOrderPaid(orderId int) (bool, error)
}

type OrderRepository struct {
}
