package event

import (
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/tiketin-management-api-with-go/database"
	"github.com/tiketin-management-api-with-go/structs"
)

func NewEventRepository() EventRepositoryInterface {
	return &EventRepository{}
}

func (r *EventRepository) CreateEvent(event structs.Event) error {
	query := `insert into events (user_id,title,description,location,event_date,event_type_id,status) values ($1,$2,$3,$4,$5,$6,$7)`

	res, err := database.DBConn.Exec(query, event.UserId, event.Title, event.Description, event.Location, event.EventDate, event.EventTypeId, event.Status)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("gagal menambahkan event")
	}

	return nil
}

func (r *EventRepository) GetAllEvent() ([]structs.Event, error) {
	var result []structs.Event
	query := `select * from events`

	rows, err := database.DBConn.Query(query)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		var data = structs.Event{}
		var errs = rows.Scan(&data.Id, &data.UserId, &data.Title, &data.Description, &data.Location, &data.EventDate, &data.EventTypeId, &data.Status, &data.CreatedAt, &data.ModifiedAt)
		if errs != nil {
			return result, errs
		}

		result = append(result, data)
	}

	return result, nil
}

func (r *EventRepository) GetEventById(id string) (structs.Event, error) {
	var result structs.Event
	query := `select * from events where id = $1`

	idInt, _ := strconv.Atoi(id)
	err := database.DBConn.QueryRow(query, idInt).Scan(&result.Id, &result.UserId, &result.Title, &result.Description, &result.Location, &result.EventDate, &result.EventTypeId, &result.Status, &result.CreatedAt, &result.ModifiedAt)
	if err == sql.ErrNoRows {
		return result, errors.New("event tidak ditemukan")
	}

	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *EventRepository) GetEventByUser(userId int) ([]structs.Event, error) {
	var result []structs.Event
	query := `select * from events where user_id=$1`

	rows, err := database.DBConn.Query(query, userId)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		var data = structs.Event{}
		var errs = rows.Scan(&data.Id, &data.UserId, &data.Title, &data.Description, &data.Location, &data.EventDate, &data.EventTypeId, &data.Status, &data.CreatedAt, &data.ModifiedAt)
		if errs != nil {
			return result, errs
		}

		result = append(result, data)
	}

	return result, nil
}

func (r *EventRepository) UpdateEvent(id string, event structs.Event) error {
	query := `update events set title=$1, description=$3, location=$4, event_date=$5, event_type_id=$6, status=$7 modified_at=$8  where id=$2`

	res, err := database.DBConn.Exec(query, event.Title, id, event.Description, event.Location, event.EventDate, event.EventTypeId, event.Status, time.Now())
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("gagal update event")
	}

	return nil
}

func (r *EventRepository) DeleteEvent(id string) error {
	query := `delete from events where event_id=$1`

	res, err := database.DBConn.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("gagal menghapus event")
	}

	return nil
}
