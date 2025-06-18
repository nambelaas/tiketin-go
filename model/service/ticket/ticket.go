package ticket

import (
	"github.com/gin-gonic/gin"
	"github.com/tiketin-management-api-with-go/model/repository/ticket"
	"github.com/tiketin-management-api-with-go/structs"
)

type TicketServiceInterface interface {
	CreateTicket(ctx *gin.Context) error
	GetAllTicketEvent(ctx *gin.Context) ([]structs.Ticket, error)
	UpdateTicket(ctx *gin.Context) error
	DeleteTicket(ctx *gin.Context) error
}

type TicketService struct {
	repo ticket.TicketRepositoryInterface
}