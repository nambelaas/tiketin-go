package ticket

import "github.com/tiketin-management-api-with-go/structs"

type TicketRepositoryInterface interface {
	CreateTicket(ticket structs.Ticket) error
	GetAllTicketEvent(eventId int) ([]structs.Ticket,error)
	UpdateTicket(ticketId string, ticket structs.Ticket) error
	DeleteTicket(ticketId string) error
}

type TicketRepository struct {
	
}