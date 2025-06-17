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

func CreateEventHandler(ctx *gin.Context) {
	err := eventService.CreateEvent(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponse(ctx, "berhasil menambahkan event")
}

func GetAllEvent(ctx *gin.Context) {
	data, err := eventService.GetAllEvent(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponseWithData(ctx, "berhasil mendapatkan event", data)
}

func GetEventById(ctx *gin.Context) {
	data, err := eventService.GetEventById(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Param("event_id")
	helper.PrintSuccessResponseWithData(ctx, fmt.Sprintf("berhasil mendapatkan event  %s", id), data)
}

func GetEventByUser(ctx *gin.Context) {
	data, err := eventService.GetEventByUser(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponseWithData(ctx, "berhasil mendapatkan list item user", data)
}

func UpdateEventHandler(ctx *gin.Context) {
	err := eventService.UpdateEvent(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponse(ctx, "berhasil update event")
}

func DeleteEventHandler(ctx *gin.Context) {
	err := eventService.DeleteEvent(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Param("event_id")
	helper.PrintSuccessResponse(ctx, fmt.Sprintf("berhasil menghapus event %s", id))
}
