package ticket

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tiketin-management-api-with-go/helper"
	"github.com/tiketin-management-api-with-go/model/repository/event"
	"github.com/tiketin-management-api-with-go/model/repository/ticket"
	"github.com/tiketin-management-api-with-go/structs"
)

func NewTicketService(repo ticket.TicketRepositoryInterface) TicketServiceInterface {
	return &TicketService{
		repo: repo,
	}
}

func (s *TicketService) CreateTicket(ctx *gin.Context) error {
	var newTicket structs.Ticket
	err := ctx.ShouldBindJSON(&newTicket)
	if err != nil {
		result := helper.ValidationCheck(err)
		if result != nil {
			return result
		}
		return err
	}

	exists := helper.IsTicketExists(newTicket)
	if exists {
		return errors.New("gagal menambahkan ticket, data sudah ada")
	}

	eventId := ctx.Param("event_id")
	dataEvent, err := event.NewEventRepository().GetEventById(eventId)
	if err != nil {
		return errors.New("id event tidak ditemukan")
	}

	newTicket.EventId = dataEvent.Id

	err = s.repo.CreateTicket(newTicket)
	if err != nil {
		return err
	}

	return nil
}

func (s *TicketService) UpdateTicket(ctx *gin.Context) error {
	var updateTicket structs.Ticket
	err := ctx.ShouldBindJSON(&updateTicket)
	if err != nil {
		result := helper.ValidationCheck(err)
		if result != nil {
			return result
		}
		return err
	}

	eventId := ctx.Param("event_id")
	dataEvent, err := event.NewEventRepository().GetEventById(eventId)
	if err != nil {
		return errors.New("id event tidak ditemukan")
	}

	updateTicket.EventId = dataEvent.Id

	ticketId, _ := strconv.Atoi(ctx.Param("ticket_id"))

	err = s.repo.UpdateTicket(ticketId, updateTicket)
	if err != nil {
		return err
	}

	return nil
}

func (s *TicketService) DeleteTicket(ctx *gin.Context) error {
	ticketId, _ := strconv.Atoi(ctx.Param("ticket_id"))

	err := s.repo.DeleteTicket(ticketId)
	if err != nil {
		return err
	}

	return nil
}

func (s *TicketService) GetAllTicketEvent(ctx *gin.Context) ([]structs.Ticket, error) {
	paramEventId := ctx.Param("event_id")

	dataEvent, err := event.NewEventRepository().GetEventById(paramEventId)
	if err != nil {
		return []structs.Ticket{}, errors.New("id event tidak ditemukan")
	}

	eventId := dataEvent.Id

	data, err := s.repo.GetAllTicketEvent(eventId)
	if err != nil {
		return []structs.Ticket{}, err
	}

	return data, nil
}
