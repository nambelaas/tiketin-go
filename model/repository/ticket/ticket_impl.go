package ticket

import (
	"errors"
	"time"

	"github.com/tiketin-management-api-with-go/database"
	"github.com/tiketin-management-api-with-go/helper"
	"github.com/tiketin-management-api-with-go/structs"
)

func NewTicketRepository() TicketRepositoryInterface {
	return &TicketRepository{}
}

func (t *TicketRepository) CreateTicket(ticket structs.Ticket) error {
	ticketId := helper.GenerateRandomString(8)

	query := `insert into tickets (ticket_id, event_id, name, price, quota) values ($1,$2,$3,$4,$5)`

	res, err := database.DBConn.Exec(query, ticketId, ticket.EventId, ticket.Name, ticket.Price, ticket.Quota)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("gagal menambahkan ticket")
	}
	return nil
}

func (t *TicketRepository) UpdateTicket(ticketId string, ticket structs.Ticket) error {
	query := `update tickets set name=$1, quota=$2, price=$3, modified_at=$4 where ticket_id=$5`

	res, err := database.DBConn.Exec(query, ticket.Name, ticket.Quota, ticket.Price, time.Now(), ticketId)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("gagal update ticket")
	}
	return nil
}

func (t *TicketRepository) DeleteTicket(ticketId string) error {

	query := `delete from tickets where ticket_id=$1`

	res, err := database.DBConn.Exec(query, ticketId)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("gagal menghapus ticket")
	}
	return nil
}

func (t *TicketRepository) GetAllTicketEvent(eventId int) ([]structs.Ticket, error) {
	var result []structs.Ticket
	query := `select * from tickets where event_id=$1`

	rows, err := database.DBConn.Query(query, eventId)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		var data = structs.Ticket{}
		var err = rows.Scan(&data.Id, &data.TicketId, &data.EventId, &data.Name, &data.Price, &data.Quota, &data.CreatedAt, &data.ModifiedAt)
		if err != nil {
			return result, err
		}

		result = append(result, data)
	}

	return result, nil
}
