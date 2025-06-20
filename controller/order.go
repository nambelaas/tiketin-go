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

//	@summary		Create Order
//	@description	Menambahkan data order
//	@description	* Hanya bisa diakses oleh user
//	@tags			Order
//	@accept			json
//	@produce		json
//	@Param			Authorization	header		string					true	"Bearer token"
//	@Param			order			body		structs.OrderCreate		true	"Order Data"
//	@Success		200				{object}	structs.SuccessStruct	"Berhasil menambahkan order"
//	@Failure		400				{object}	structs.ErrorStruct		"Gagal menambahkan order"
//	@Router			/api/orders/create [post]
func CreateOrderHandle(ctx *gin.Context) {
	err := orderService.CreateOrder(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponse(ctx, "berhasil menambahkan order")
}

//	@summary		Pay Order
//	@description	Memperbarui status pembayaran order
//	@description	* Hanya bisa diakses oleh user
//	@tags			Order
//	@accept			json
//	@produce		json
//	@Param			Authorization	header		string					true	"Bearer token"
//	@Param			order			body		structs.Order			true	"Order Data"
//	@Success		200				{object}	structs.SuccessStruct	"Berhasil membayar order"
//	@Failure		400				{object}	structs.ErrorStruct		"Gagal membayar order"
//	@Router			/api/orders/:order_id/pay [put]
func PayOrderHandle(ctx *gin.Context) {
	err := orderService.PayOrder(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	orderId := ctx.Param("order_id")
	helper.PrintSuccessResponse(ctx, fmt.Sprintf("berhasil membayar order %s", orderId))
}

//	@summary		Cancel Order
//	@description	Membatalkan order
//	@description	* Hanya bisa diakses oleh user
//	@tags			Order
//	@accept			json
//	@produce		json
//	@Param			Authorization	header		string					true	"Bearer token"
//	@Success		200				{object}	structs.SuccessStruct	"Berhasil membatalkan order"
//	@Failure		400				{object}	structs.ErrorStruct		"Gagal membatalkan order"
//	@Router			/api/orders/:order_id/cancel [put]
func CancelOrderHandle(ctx *gin.Context) {
	err := orderService.CancelOrder(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	orderId := ctx.Param("order_id")
	helper.PrintSuccessResponse(ctx, fmt.Sprintf("berhasil membatalkan order %s", orderId))
}

//	@summary		List Order
//	@description	Mendapatkan daftar order
//	@tags			Order
//	@accept			json
//	@produce		json
//	@Param			Authorization	header		string							true	"Bearer token"
//	@Success		200				{object}	structs.SuccessStructWithData	"Berhasil mendapatkan daftar order"
//	@Failure		400				{object}	structs.ErrorStruct				"Gagal mendapatkan daftar order"
//	@Router			/api/orders/list [put]
func GetAllOrderHandle(ctx *gin.Context) {
	order, err := orderService.GetAllOrder(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponseWithData(ctx, "berhasil mendapatkan semua data order", order)
}

//	@summary		Detail Order
//	@description	Mendapatkan order berdasarkan id
//	@tags			Order
//	@accept			json
//	@produce		json
//	@Param			Authorization	header		string							true	"Bearer token"
//	@Success		200				{object}	structs.SuccessStructWithData	"Berhasil mendapatkan detail order"
//	@Failure		400				{object}	structs.ErrorStruct				"Gagal mendapatkan detail order"
//	@Router			/api/orders/:order_id [get]
func GetOrderByIdHandle(ctx *gin.Context) {
	order, err := orderService.GetOrderById(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	orderId := ctx.Param("order_id")
	helper.PrintSuccessResponseWithData(ctx, fmt.Sprintf("berhasil mendapatkan data order %s",orderId), order)
}

//	@summary		List Order User
//	@description	Mendapatkan daftar order user
//	@tags			Order
//	@accept			json
//	@produce		json
//	@Param			Authorization	header		string							true	"Bearer token"
//	@Success		200				{object}	structs.SuccessStructWithData	"Berhasil mendapatkan daftar order user"
//	@Failure		400				{object}	structs.ErrorStruct				"Gagal mendapatkan daftar order user"
//	@Router			/api/orders/user/:user_id/list [get]
func GetOrderByUserHandle(ctx *gin.Context) {
	order, err := orderService.GetOrderByUser(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userId := ctx.Param("user_id")
	helper.PrintSuccessResponseWithData(ctx, fmt.Sprintf("berhasil mendapatkan data order dari user %s",userId), order)
}

//	@summary		Check In Order
//	@description	Melakukan check-in pada order
//	@tags			Order
//	@accept			json
//	@produce		json
//	@Param			orderId		query		string					true	"Order Id"
//	@Param			orderItemId	query		string					true	"Order Item Id"
//	@Param			ticketId	query		string					true	"Ticket Id"
//	@Success		200			{object}	structs.SuccessStruct	"Berhasil melakukan check-in order"
//	@Failure		400			{object}	structs.ErrorStruct		"Gagal melakukan check-in order"
//	@Router			/api/orders/checkin/ticket [get]
func CheckInHandle(ctx *gin.Context) {
	err := orderService.CheckIn(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	orderId := ctx.Param("order_id")
	helper.PrintSuccessResponse(ctx, fmt.Sprintf("berhasil checkin order %s",orderId))
}
