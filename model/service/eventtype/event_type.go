package eventtype

import (
	"github.com/gin-gonic/gin"
	"github.com/tiketin-management-api-with-go/model/repository/eventtype"
	"github.com/tiketin-management-api-with-go/structs"
)

type EventTypeServiceInterface interface {
	CreateEventType(ctx *gin.Context) error
	GetAllEventType(ctx *gin.Context) ([]structs.EventType, error)
	GetEventByEventTypeId(ctx *gin.Context) ([]structs.Event, error)
	UpdateEventType(ctx *gin.Context) error
	DeleteEventType(ctx *gin.Context) error
}

type EventTypeService struct {
	repo eventtype.EventTypeRepositoryInterface
}

