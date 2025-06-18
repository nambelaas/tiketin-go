package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiketin-management-api-with-go/helper"
	repository "github.com/tiketin-management-api-with-go/model/repository/ticket"
	service "github.com/tiketin-management-api-with-go/model/service/ticket"
)

var (
	ticketRepository = repository.NewTicketRepository()
	ticketService    = service.NewTicketService(ticketRepository)
)

func CreateTicketHandle(ctx *gin.Context) {
	err := ticketService.CreateTicket(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponse(ctx, "berhasil menambahkan ticket")
}

func GetAllTicketEventHandle(ctx *gin.Context) {
	data, err := ticketService.GetAllTicketEvent(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponseWithData(ctx, "berhasil mendapatkan data ticket", data)
}

func UpdateTicketHandle(ctx *gin.Context) {
	err := ticketService.UpdateTicket(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponse(ctx, "berhasil update ticket")
}

func DeleteTicketHandle(ctx *gin.Context) {
	err := ticketService.DeleteTicket(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Param("ticket_id")
	helper.PrintSuccessResponse(ctx, fmt.Sprintf("berhasil menghapus ticket %s", id))
}
