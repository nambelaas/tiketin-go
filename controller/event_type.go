package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiketin-management-api-with-go/helper"
	repository "github.com/tiketin-management-api-with-go/model/repository/eventtype"
	service "github.com/tiketin-management-api-with-go/model/service/eventtype"
)

var (
	eventTypeRepository = repository.NewEventTypeRepository()
	eventTypeService    = service.NewEventTypeService(eventTypeRepository)
)

//	@summary		Create Event Type
//	@description	Mendaftarkan jenis event
//	@description	* Hanya bisa diakses oleh admin
//	@tags			Event Type
//	@accept			json
//	@produce		json
//	@Param			Authorization	header		string					true	"Bearer token"
//	@Param			event-type		body		structs.EventType		true	"Event Type Data"
//	@Success		200				{object}	structs.SuccessStruct	"Berhasil menambahkan event type"
//	@Failure		400				{object}	structs.ErrorStruct		"Gagal menambahkan event type"
//	@Router			/api/event-types/create [post]
func CreateEventTypeHandle(ctx *gin.Context) {
	err := eventTypeService.CreateEventType(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponse(ctx, "berhasil menambahkan event type")
}

//	@summary		Get All Event Type
//	@description	Mendapatkan semua jenis event
//	@tags			Event Type
//	@accept			json
//	@produce		json
//	@Success		200	{object}	structs.SuccessStructWithData	"Berhasil mendapatkan event type"
//	@Failure		400	{object}	structs.ErrorStruct				"Gagal mendapatkan event type"
//	@Router			/api/event-types/list [get]
func GetAllEventTypeHandle(ctx *gin.Context) {
	data, err := eventTypeService.GetAllEventType(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponseWithData(ctx, "berhasil mendapatkan event type", data)
}

//	@summary		Get Event By Event Type Id
//	@description	Mendapatkan daftar event berdasarkan jenis event
//	@tags			Event Type
//	@accept			json
//	@produce		json
//	@Success		200	{object}	structs.SuccessStructWithData	"Berhasil mendapatkan list event"
//	@Failure		400	{object}	structs.ErrorStruct				"Gagal mendapatkan list event"
//	@Router			/api/event-types/:id/events [get]
func GetEventByEventTypeIdHandle(ctx *gin.Context) {
	data, err := eventTypeService.GetEventByEventTypeId(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Param("id")
	helper.PrintSuccessResponseWithData(ctx, fmt.Sprintf("berhasil mendapatkan list event dengan event type %s", id), data)
}

//	@summary		Update Event Type
//	@description	Memperbarui jenis event
//	@description	* Hanya bisa diakses oleh admin
//	@tags			Event Type
//	@accept			json
//	@produce		json
//	@Param			Authorization	header		string					true	"Bearer token"
//	@Param			event-type		body		structs.EventType		true	"Event Type Data"
//	@Success		200				{object}	structs.SuccessStruct	"Berhasil memperbarui event type"
//	@Failure		400				{object}	structs.ErrorStruct		"Gagal memperbarui event type"
//	@Router			/api/event-types/:id/update [put]
func UpdateEventTypeHandle(ctx *gin.Context) {
	err := eventTypeService.UpdateEventType(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponse(ctx, "berhasil update event type")
}

//	@summary		Delete Event Type
//	@description	Menghapus jenis event
//	@description	* Hanya bisa diakses oleh admin
//	@tags			Event Type
//	@accept			json
//	@produce		json
//	@Param			Authorization	header		string					true	"Bearer token"
//	@Success		200				{object}	structs.SuccessStruct	"Berhasil menghapus event type"
//	@Failure		400				{object}	structs.ErrorStruct		"Gagal menghapus event type"
//	@Router			/api/event-types/:id/delete [delete]
func DeleteEventTypeHandle(ctx *gin.Context) {
	err := eventTypeService.DeleteEventType(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Param("id")
	helper.PrintSuccessResponse(ctx, fmt.Sprintf("berhasil menghapus event type %s", id))
}
