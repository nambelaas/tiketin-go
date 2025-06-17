package eventtype

import (
	"github.com/gin-gonic/gin"
	"github.com/tiketin-management-api-with-go/model/repository/event"
	"github.com/tiketin-management-api-with-go/structs"
)

type EventServiceInterface interface {
	CreateEvent(ctx *gin.Context) error
	GetAllEvent(ctx *gin.Context) ([]structs.Event, error)
	GetEventById(ctx *gin.Context) (structs.Event, error)
	GetEventByUser(ctx *gin.Context) ([]structs.Event, error)
	UpdateEvent(ctx *gin.Context) error
	DeleteEvent(ctx *gin.Context) error
}

type EventService struct {
	repo event.EventRepositoryInterface
}

