package order

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tiketin-management-api-with-go/helper"
	"github.com/tiketin-management-api-with-go/model/repository/order"
	"github.com/tiketin-management-api-with-go/model/repository/orderitem"
	"github.com/tiketin-management-api-with-go/model/repository/ticket"
	"github.com/tiketin-management-api-with-go/structs"
)

func NewOrderService(repo order.OrderRepositoryInterface) OrderServiceInterface {
	return &OrderService{
		repo: repo,
	}
}

func (s *OrderService) CreateOrder(ctx *gin.Context) error {
	var newOrder structs.OrderCreate
	err := ctx.ShouldBindJSON(&newOrder)
	if err != nil {
		result := helper.ValidationCheck(err)
		if result != nil {
			return result
		}
		return err
	}

	for _, item := range newOrder.Ticket {
		availableQuota, err := ticket.NewTicketRepository().AvailableQuota(item.TicketTypeId)
		if err != nil {
			return err
		}

		calculateLeftQuota := availableQuota - item.Quantity

		if calculateLeftQuota < 0 {
			return errors.New("kuota tiket tidak cukup untuk jenis tiket " + strconv.Itoa(item.TicketTypeId))
		}
	}

	dataJwt, err := helper.GetJwtData(ctx)
	if err != nil {
		return err
	}

	userId := dataJwt.UserId

	err = s.repo.CreateOrder(userId, newOrder)
	if err != nil {
		return err
	}

	return nil
}

func (s *OrderService) PayOrder(ctx *gin.Context) error {
	var dataOrder structs.Order
	err := ctx.ShouldBindJSON(&dataOrder)
	if err != nil {
		result := helper.ValidationCheck(err)
		if result != nil {
			return result
		}
		return err
	}

	orderId, _ := strconv.Atoi(ctx.Param("order_id"))

	err = s.repo.PayOrder(orderId, dataOrder)
	if err != nil {
		return err
	}

	orderItems, _ := orderitem.NewOrderItemRepository().GetAllOrderItemByOrderId(orderId)

	for _, item := range orderItems {
		err := ticket.NewTicketRepository().ReduceQuota(item.TicketTypeId, item.TicketTypeId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *OrderService) CancelOrder(ctx *gin.Context) error {
	orderId, _ := strconv.Atoi(ctx.Param("order_id"))

	err := s.repo.CancelOrder(orderId)
	if err != nil {
		return err
	}

	orderItems, _ := orderitem.NewOrderItemRepository().GetAllOrderItemByOrderId(orderId)

	for _, item := range orderItems {
		err := ticket.NewTicketRepository().RestoreQuota(item.TicketTypeId, item.TicketTypeId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *OrderService) GetAllOrder(ctx *gin.Context) ([]structs.Order, error) {
	data, err := s.repo.GetAllOrder()
	if err != nil {
		return []structs.Order{}, err
	}

	return data, nil
}

func (s *OrderService) GetOrderByUser(ctx *gin.Context) ([]structs.Order, error) {
	userId, _ := strconv.Atoi(ctx.Param("user_id"))

	data, err := s.repo.GetOrderByUser(userId)
	if err != nil {
		return []structs.Order{}, err
	}

	return data, nil
}

func (s *OrderService) GetOrderById(ctx *gin.Context) (structs.Order, error) {
	orderId, _ := strconv.Atoi(ctx.Param("order_id"))

	data, err := s.repo.GetOrderById(orderId)
	if err != nil {
		return structs.Order{}, err
	}

	return data, nil
}

func (s *OrderService) CheckIn(ctx *gin.Context) error {
	orderId, _ := strconv.Atoi(ctx.Query("orderId"))
	orderItemId, _ := strconv.Atoi(ctx.Query("orderItemId"))

	// cek apakah orderId valid
	orderExist := helper.IsOrderExists(orderId)
	if !orderExist {
		return errors.New("tidak bisa check in karena data order tidak ditemukan")
	}

	// cek apakah order item sudah dicheck in
	alreadyCheckIn, err := helper.IsOrderAlreadyCheckIn(orderItemId)
	if err != nil {
		return errors.New("terdapat error saat check in")
	}

	if alreadyCheckIn {
		return errors.New("order sudah dicheck in")
	}

	err = s.repo.CheckIn(orderItemId)
	if err != nil {
		return err
	}

	// cek apakah order sudah dibayar
	isOrderPaid, err := s.repo.IsOrderPaid(orderId)
	if err != nil {
		return err
	}

	if !isOrderPaid {
		return errors.New("tidak bisa check in karena order belum dibayar")
	}

	// update status order ke complete jika semua data order item sudah di check in
	allCheckedIn, err := s.repo.IsAllOrderItemCheckedIn(orderId)
	if err != nil {
		return errors.New("gagal cek status order item")
	}

	if allCheckedIn {
		err = s.repo.UpdateOrderStatus(orderId, "complete")
		if err != nil {
			return errors.New("gagal update status order menjadi complete")
		}
	}

	return nil
}
