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

func CreateEventTypeHandler(ctx *gin.Context) {
	err := eventTypeService.CreateEventType(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponse(ctx, "berhasil menambahkan event type")
}

func GetAllEventType(ctx *gin.Context) {
	data, err := eventTypeService.GetAllEventType(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponseWithData(ctx, "berhasil mendapatkan event type", data)
}

func GetEventByEventTypeId(ctx *gin.Context) {
	data, err := eventTypeService.GetEventByEventTypeId(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Param("id")
	helper.PrintSuccessResponseWithData(ctx, fmt.Sprintf("berhasil mendapatkan list event dengan event type %s", id), data)
}

func UpdateEventTypeHandler(ctx *gin.Context) {
	err := eventTypeService.UpdateEventType(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponse(ctx, "berhasil update event type")
}

func DeleteEventTypeHandler(ctx *gin.Context) {
	err := eventTypeService.DeleteEventType(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Param("id")
	helper.PrintSuccessResponse(ctx, fmt.Sprintf("berhasil menghapus event type %s", id))
}
