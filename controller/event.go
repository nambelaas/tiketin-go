package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiketin-management-api-with-go/helper"
	repository "github.com/tiketin-management-api-with-go/model/repository/event"
	service "github.com/tiketin-management-api-with-go/model/service/event"
)

var (
	eventRepository = repository.NewEventRepository()
	eventService    = service.NewEventService(eventRepository)
)

//	@summary		Create Event
//	@description	Menambahkan data event
//	@description	* Hanya bisa diakses oleh admin atau organizer
//	@tags			Event
//	@accept			json
//	@produce		json
//	@Param			Authorization	header		string					true	"Bearer token"
//	@Param			event			body		structs.Event			true	"Event Data"
//	@Success		200				{object}	structs.SuccessStruct	"Berhasil menambahkan event"
//	@Failure		400				{object}	structs.ErrorStruct		"Gagal menambahkan event"
//	@Router			/api/events/create [post]
func CreateEventHandle(ctx *gin.Context) {
	err := eventService.CreateEvent(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponse(ctx, "berhasil menambahkan event")
}

//	@summary		List Event
//	@description	Mendapatkan semua daftar event
//	@tags			Event
//	@accept			json
//	@produce		json
//	@Success		200	{object}	[]structs.SuccessStructWithData	"Berhasil mendapatkan daftar event"
//	@Failure		400	{object}	structs.ErrorStruct				"Gagal mendapatkan daftar event"
//	@Router			/api/events/list [get]
func GetAllEventHandle(ctx *gin.Context) {
	data, err := eventService.GetAllEvent(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponseWithData(ctx, "berhasil mendapatkan event", data)
}

//	@summary		Detail Event
//	@description	Mendapatkan detail event berdasarkan id
//	@tags			Event
//	@accept			json
//	@produce		json
//	@Success		200	{object}	structs.SuccessStructWithData	"Berhasil mendapatkan daftar event"
//	@Failure		400	{object}	structs.ErrorStruct				"Gagal mendapatkan daftar event"
//	@Router			/api/events/:event_id [get]
func GetEventByIdHandle(ctx *gin.Context) {
	data, err := eventService.GetEventById(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Param("event_id")
	helper.PrintSuccessResponseWithData(ctx, fmt.Sprintf("berhasil mendapatkan event  %s", id), data)
}

//	@summary		Detail Event By User
//	@description	Mendapatkan daftar event berdasarkan id user
//	@description	* Hanya bisa diakses oleh admin atau organizer
//	@tags			Event
//	@accept			json
//	@produce		json
//	@Param			Authorization	header		string							true	"Bearer token"
//	@Success		200				{object}	structs.SuccessStructWithData	"Berhasil mendapatkan daftar event"
//	@Failure		400				{object}	structs.ErrorStruct				"Gagal mendapatkan daftar event"
//	@Router			/api/events/me [get]
func GetEventByUserHandle(ctx *gin.Context) {
	data, err := eventService.GetEventByUser(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponseWithData(ctx, "berhasil mendapatkan list item user", data)
}

//	@summary		Update event
//	@description	Memperbarui data event
//	@description	* Hanya bisa diakses oleh admin atau organizer
//	@tags			Event
//	@accept			json
//	@produce		json
//	@Param			Authorization	header		string					true	"Bearer token"
//	@Param			event			body		structs.Event			true	"Event Data"
//	@Success		200				{object}	structs.SuccessStruct	"Berhasil memperbarui event"
//	@Failure		400				{object}	structs.ErrorStruct		"Gagal memperbarui event"
//	@Router			/api/events/:event_id/update [put]
func UpdateEventHandle(ctx *gin.Context) {
	err := eventService.UpdateEvent(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponse(ctx, "berhasil update event")
}

//	@summary		Delete Event
//	@description	Menghapus event berdasarkan id
//	@description	* Hanya bisa diakses oleh admin atau organizer
//	@tags			Event
//	@accept			json
//	@produce		json
//	@Param			Authorization	header		string					true	"Bearer token"
//	@Success		200				{object}	structs.SuccessStruct	"Berhasil menghapus event"
//	@Failure		400				{object}	structs.ErrorStruct		"Gagal menghapus event"
//	@Router			/api/events/:event_id/delete [delete]
func DeleteEventHandle(ctx *gin.Context) {
	err := eventService.DeleteEvent(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Param("event_id")
	helper.PrintSuccessResponse(ctx, fmt.Sprintf("berhasil menghapus event %s", id))
}
