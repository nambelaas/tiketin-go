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

//	@summary		Create Ticket
//	@description	Menambahkan data ticket untuk event
//	@description	* Hanya bisa diakses oleh admin atau organizer
//	@tags			Ticket
//	@accept			json
//	@produce		json
//	@Param			Authorization	header		string					true	"Bearer token"
//	@Param			ticket			body		structs.Ticket			true	"Ticket Data"
//	@Success		200				{object}	structs.SuccessStruct	"Berhasil menambahkan ticket"
//	@Failure		400				{object}	structs.ErrorStruct		"Gagal menambahkan ticket"
//	@Router			/api/events/:event_id/tickets/create [post]
func CreateTicketHandle(ctx *gin.Context) {
	err := ticketService.CreateTicket(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponse(ctx, "berhasil menambahkan ticket")
}

//	@summary		List Ticket
//	@description	Mendapatkan daftar ticket untuk event
//	@description	* Hanya bisa diakses oleh admin atau organizer
//	@tags			Ticket
//	@accept			json
//	@produce		json
//	@Param			Authorization	header		string							true	"Bearer token"
//	@Success		200				{object}	structs.SuccessStructWithData	"Berhasil mendapatkan ticket"
//	@Failure		400				{object}	structs.ErrorStruct				"Gagal mendapatkan ticket"
//	@Router			/api/events/:event_id/tickets/list [get]
func GetAllTicketEventHandle(ctx *gin.Context) {
	data, err := ticketService.GetAllTicketEvent(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponseWithData(ctx, "berhasil mendapatkan data ticket", data)
}

//	@summary		Update Ticket
//	@description	Memperbarui data ticket untuk event
//	@description	* Hanya bisa diakses oleh admin atau organizer
//	@tags			Ticket
//	@accept			json
//	@produce		json
//	@Param			Authorization	header		string					true	"Bearer token"
//	@Param			ticket			body		structs.Ticket			true	"Ticket Data"
//	@Success		200				{object}	structs.SuccessStruct	"Berhasil memperbarui ticket"
//	@Failure		400				{object}	structs.ErrorStruct		"Gagal memperbarui ticket"
//	@Router			/api/events/:event_id/tickets/:ticket_id/update [put]
func UpdateTicketHandle(ctx *gin.Context) {
	err := ticketService.UpdateTicket(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponse(ctx, "berhasil update ticket")
}

//	@summary		Delete Ticket
//	@description	Menghapus data ticket untuk event
//	@description	* Hanya bisa diakses oleh admin atau organizer
//	@tags			Ticket
//	@accept			json
//	@produce		json
//	@Param			Authorization	header		string					true	"Bearer token"
//	@Success		200				{object}	structs.SuccessStruct	"Berhasil menghapus ticket"
//	@Failure		400				{object}	structs.ErrorStruct		"Gagal menghapus ticket"
//	@Router			/api/events/:event_id/tickets/:ticket_id/delete [put]
func DeleteTicketHandle(ctx *gin.Context) {
	err := ticketService.DeleteTicket(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Param("ticket_id")
	helper.PrintSuccessResponse(ctx, fmt.Sprintf("berhasil menghapus ticket %s", id))
}
