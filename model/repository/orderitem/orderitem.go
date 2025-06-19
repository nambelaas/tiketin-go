package orderitem

import "github.com/tiketin-management-api-with-go/structs"

type OrderItemRepositoryInterface interface {
	GetAllOrderItemByOrderId(orderId int) ([]structs.OrderItem, error)
}

type OrderItemRepository struct {
}
