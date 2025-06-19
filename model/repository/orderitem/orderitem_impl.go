package orderitem

import (
	"github.com/tiketin-management-api-with-go/database"
	"github.com/tiketin-management-api-with-go/structs"
)

func NewOrderItemRepository() OrderItemRepositoryInterface {
	return &OrderItemRepository{}
}

func (r *OrderItemRepository) GetAllOrderItemByOrderId(orderId int) ([]structs.OrderItem, error) {
	var result []structs.OrderItem
	query := `select * from order_items where order_id=$1`

	rows, err := database.DBConn.Query(query, orderId)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		var item = structs.OrderItem{}
		var err = rows.Scan(&item.Id, &item.OrderId, &item.TicketTypeId,&item.Quantity, &item.QrCodeUrl, &item.IsCheckIn, &item.CreatedAt, &item.ModifiedAt)
		if err != nil {
			return result, err
		}

		result = append(result, item)
	}

	return result, nil
}
