package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiketin-management-api-with-go/helper"
	repository "github.com/tiketin-management-api-with-go/model/repository/order"
	service "github.com/tiketin-management-api-with-go/model/service/order"
)

var (
	orderRepository = repository.NewOrderRepository()
	orderService    = service.NewOrderService(orderRepository)
)

func CreateOrderHandle(ctx *gin.Context) {
	err := orderService.CreateOrder(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponse(ctx, "berhasil menambahkan order")
}

func PayOrderHandle(ctx *gin.Context) {
	err := orderService.PayOrder(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	orderId := ctx.Param("order_id")
	helper.PrintSuccessResponse(ctx, fmt.Sprintf("berhasil membayar order %s", orderId))
}

func CancelOrderHandle(ctx *gin.Context) {
	err := orderService.CancelOrder(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	orderId := ctx.Param("order_id")
	helper.PrintSuccessResponse(ctx, fmt.Sprintf("berhasil membatalkan order %s", orderId))
}

func GetAllOrderHandle(ctx *gin.Context) {
	order, err := orderService.GetAllOrder(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponseWithData(ctx, "berhasil mendapatkan semua data order", order)
}

func GetOrderByIdHandle(ctx *gin.Context) {
	order, err := orderService.GetOrderById(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	orderId := ctx.Param("order_id")
	helper.PrintSuccessResponseWithData(ctx, fmt.Sprintf("berhasil mendapatkan data order %s",orderId), order)
}

func GetOrderByUserHandle(ctx *gin.Context) {
	order, err := orderService.GetOrderByUser(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userId := ctx.Param("user_id")
	helper.PrintSuccessResponseWithData(ctx, fmt.Sprintf("berhasil mendapatkan data order dari user %s",userId), order)
}

func CheckInHandle(ctx *gin.Context) {
	err := orderService.CheckIn(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	orderId := ctx.Param("order_id")
	helper.PrintSuccessResponse(ctx, fmt.Sprintf("berhasil checkin order %s",orderId))
}
