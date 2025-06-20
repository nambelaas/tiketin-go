package order

import (
	"errors"
	"fmt"
	"time"

	"github.com/tiketin-management-api-with-go/database"
	"github.com/tiketin-management-api-with-go/helper"
	"github.com/tiketin-management-api-with-go/model/repository/orderitem"
	"github.com/tiketin-management-api-with-go/structs"
)

func NewOrderRepository() OrderRepositoryInterface {
	return &OrderRepository{}
}

func (r *OrderRepository) CreateOrder(userId int, order structs.OrderCreate) error {
	var totalPrice float32
	for _, ticket := range order.Ticket {
		var price float32
		err := database.DBConn.QueryRow(`SELECT price FROM tickets WHERE id = $1`, ticket.TicketTypeId).Scan(&price)
		if err != nil {
			return errors.New("gagal mengambil harga tiket")
		}
		totalPrice += price * float32(ticket.Quantity)
	}

	query := `insert into orders (user_id, event_id, total_price) values ($1,$2,$3) returning id`

	var orderId int
	err := database.DBConn.QueryRow(query, userId, order.EventId, totalPrice).Scan(&orderId)
	if err != nil {
		return err
	}

	// insert to order item table
	for _, ticket := range order.Ticket {
		for i := 0; i < ticket.Quantity; i++ {
			queryOrderItem := `insert into order_items (order_id,ticket_type_id,quantity) values ($1,$2,$3) returning id`

			var orderItemId int
			err = database.DBConn.QueryRow(queryOrderItem, orderId, ticket.TicketTypeId, 1).Scan(&orderItemId)
			if err != nil {
				fmt.Println(err)
				_, rollbackErr := database.DBConn.Exec(`DELETE FROM orders WHERE id = $1`, orderId)
				if rollbackErr != nil {
					return fmt.Errorf("gagal rollback order %d setelah gagal menambahkan order item: %v", orderId, rollbackErr)
				}

				return fmt.Errorf("gagal menambahkan order item untuk order %d", orderId)
			}

			qrUrl, err := helper.GenerateQRCode(orderId, orderItemId)
			if err != nil {
				fmt.Println(err)
				_, rollbackErr := database.DBConn.Exec(`DELETE FROM orders WHERE id = $1`, orderId)
				if rollbackErr != nil {
					return fmt.Errorf("gagal rollback order %d setelah gagal menambahkan order item: %v", orderId, rollbackErr)
				}

				return err
			}

			_, err = database.DBConn.Exec(`update order_items set qr_code_url=$1, modified_at=$3 where id=$2`, qrUrl, orderItemId, time.Now())
			if err != nil {
				fmt.Println(err)
				_, rollbackErr := database.DBConn.Exec(`DELETE FROM orders WHERE id = $1`, orderId)
				if rollbackErr != nil {
					return fmt.Errorf("gagal rollback order %d setelah gagal menambahkan order item: %v", orderId, rollbackErr)
				}
				return fmt.Errorf("gagal update qr_code_url untuk order item %d", orderItemId)
			}
		}
	}

	return nil
}

func (r *OrderRepository) PayOrder(id int, order structs.Order) error {
	query := `update orders set status=$1, paid_at=$2, payment_method=$3, modified_at=$4 where id=$5`

	res, err := database.DBConn.Exec(query, order.Status, time.Now(), order.PaymentMethod, time.Now(), id)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("gagal update status pembayaran order")
	}

	return nil
}

func (r *OrderRepository) CancelOrder(id int) error {
	query := `update orders set status=$1, modified_at=$2 where id=$3`

	res, err := database.DBConn.Exec(query, "cancelled", time.Now(), id)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("gagal membatalkan order")
	}

	return nil
}

func (r *OrderRepository) GetAllOrder() ([]structs.Order, error) {
	var result []structs.Order
	query := `select * from orders`

	rows, err := database.DBConn.Query(query)
	if err != nil {
		return result, errors.New("gagal mengambil data order")
	}

	for rows.Next() {
		var data = structs.Order{}
		var err = rows.Scan(&data.Id, &data.UserId, &data.EventId, &data.Status, &data.TotalPrice, &data.PaidAt, &data.PaymentMethod, &data.CreatedAt, &data.ModifiedAt)
		if err != nil {
			return result, err
		}

		orderItems, err := orderitem.NewOrderItemRepository().GetAllOrderItemByOrderId(data.Id)
		if err != nil {
			return result, err
		}

		data.OrderItem = orderItems

		result = append(result, data)
	}

	return result, nil
}

func (r *OrderRepository) GetOrderByUser(userId int) ([]structs.Order, error) {
	var result []structs.Order
	query := `select * from orders where user_id = $1`

	rows, err := database.DBConn.Query(query, userId)
	if err != nil {
		return result, errors.New("gagal mengambil data order")
	}

	for rows.Next() {
		var data = structs.Order{}
		var err = rows.Scan(&data.Id, &data.UserId, &data.EventId, &data.Status, &data.TotalPrice, &data.PaidAt, &data.PaymentMethod, &data.CreatedAt, &data.ModifiedAt)
		if err != nil {
			return result, err
		}

		orderItems, err := orderitem.NewOrderItemRepository().GetAllOrderItemByOrderId(data.Id)
		if err != nil {
			return result, err
		}

		data.OrderItem = orderItems

		result = append(result, data)
	}

	return result, nil
}

func (r *OrderRepository) GetOrderById(id int) (structs.Order, error) {
	var result structs.Order
	query := `select * from orders where id = $1`

	err := database.DBConn.QueryRow(query, id).Scan(&result.Id, &result.UserId, &result.EventId, &result.Status, &result.TotalPrice, &result.PaidAt, &result.PaymentMethod, &result.CreatedAt, &result.ModifiedAt)
	if err != nil {
		return result, errors.New("gagal mengambil data order")
	}

	orderItems, err := orderitem.NewOrderItemRepository().GetAllOrderItemByOrderId(id)
	if err != nil {
		return result, err
	}

	result.OrderItem = orderItems

	return result, nil
}

func (r *OrderRepository) CheckIn(id int) error {
	query := `update order_items set is_check_in=$1, modified_at=$2 where id=$3`

	res, err := database.DBConn.Exec(query, true, time.Now(), id)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("gagal update status check in order item %d", id)
	}

	return nil
}

func (r *OrderRepository) IsAllOrderItemCheckedIn(orderId int) (bool, error) {
	var total, checkedIn int
	err := database.DBConn.QueryRow(`SELECT COUNT(*) FROM order_items WHERE order_id = $1`, orderId).Scan(&total)
	if err != nil {
		return false, err
	}
	err = database.DBConn.QueryRow(`SELECT COUNT(*) FROM order_items WHERE order_id = $1 AND is_check_in = true`, orderId).Scan(&checkedIn)
	if err != nil {
		return false, err
	}
	return total > 0 && total == checkedIn, nil
}

func (r *OrderRepository) UpdateOrderStatus(orderId int, status string) error {
	_, err := database.DBConn.Exec(
		`UPDATE orders SET status = $1 WHERE id = $2`, status, orderId,
	)
	return err
}

func (r *OrderRepository) IsOrderPaid(orderId int) (bool, error) {
	query := `select status from orders where id=$1`
	var status string
	err := database.DBConn.QueryRow(query, orderId).Scan(&status)
	if err != nil {
		return false, err
	}

	if status != "paid" {
		return false, nil
	}

	return true, nil
}
