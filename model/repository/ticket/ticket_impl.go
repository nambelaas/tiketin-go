package ticket

import (
	"errors"
	"time"

	"github.com/tiketin-management-api-with-go/database"
	"github.com/tiketin-management-api-with-go/structs"
)

func NewTicketRepository() TicketRepositoryInterface {
	return &TicketRepository{}
}

func (t *TicketRepository) CreateTicket(ticket structs.Ticket) error {
	query := `insert into tickets (event_id, name, price, quota) values ($1,$2,$3,$4)`

	res, err := database.DBConn.Exec(query, ticket.EventId, ticket.Name, ticket.Price, ticket.Quota)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("gagal menambahkan ticket")
	}
	return nil
}

func (t *TicketRepository) UpdateTicket(id int, ticket structs.Ticket) error {
	query := `update tickets set name=$1, quota=$2, price=$3, modified_at=$4 where id=$5`

	res, err := database.DBConn.Exec(query, ticket.Name, ticket.Quota, ticket.Price, time.Now(), id)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("gagal update ticket")
	}
	return nil
}

func (t *TicketRepository) DeleteTicket(id int) error {

	query := `delete from tickets where id=$1`

	res, err := database.DBConn.Exec(query, id)
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
		var err = rows.Scan(&data.Id, &data.EventId, &data.Name, &data.Price, &data.Quota, &data.CreatedAt, &data.ModifiedAt)
		if err != nil {
			return result, err
		}

		result = append(result, data)
	}

	return result, nil
}

func (t *TicketRepository) ReduceQuota(ticketId int, quantity int) error {
	query := `update tickets set quota=quota-$1, modified_at=$2 where id=$3`

	res, err := database.DBConn.Exec(query, quantity, time.Now(), ticketId)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("gagal mengurangi kuota ticket")
	}

	return nil
}

func (t *TicketRepository) RestoreQuota(ticketId int, quantity int) error {
	query := `update tickets set quota=quota+$1, modified_at=$2 where id=$3`

	res, err := database.DBConn.Exec(query, quantity, time.Now(), ticketId)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("gagal mengurangi kuota ticket")
	}

	return nil
}

func (t *TicketRepository) AvailableQuota(ticketId int) (int, error) {
	query := `select quota from tickets where id=$1`

	var quota int
	err := database.DBConn.QueryRow(query, ticketId).Scan(&quota)
	if err != nil {
		return 0, err
	}

	return quota, nil
}
