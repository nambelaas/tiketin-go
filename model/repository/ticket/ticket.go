package ticket

import "github.com/tiketin-management-api-with-go/structs"

type TicketRepositoryInterface interface {
	CreateTicket(ticket structs.Ticket) error
	GetAllTicketEvent(eventId int) ([]structs.Ticket, error)
	UpdateTicket(id int, ticket structs.Ticket) error
	DeleteTicket(id int) error
	ReduceQuota(ticketId int, quantity int) error
	RestoreQuota(ticketId int, quantity int) error
	AvailableQuota(ticketId int) (int, error)
}

type TicketRepository struct {
}
