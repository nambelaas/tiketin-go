package event

import (
	"errors"
	"time"

	"github.com/tiketin-management-api-with-go/database"
	"github.com/tiketin-management-api-with-go/helper"
	"github.com/tiketin-management-api-with-go/structs"
)

func NewEventRepository() EventRepositoryInterface {
	return &EventRepository{}
}

func (r *EventRepository) CreateEvent(event structs.Event) error {
	eventId := helper.GenerateRandomString(7)

	query := `insert into events (event_id,user_id,title,description,location,event_date,event_type_id) values ($1,$2,$3,$4,$5,$6,$7)`

	res, err := database.DBConn.Exec(query, eventId, event.UserId, event.Title, event.Description, event.Location, event.EventDate, event.EventTypeId)
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
		var errs = rows.Scan(&data.Id, &data.EventId, &data.UserId, &data.Title, &data.Description, &data.Location, &data.EventDate, &data.EventTypeId, &data.CreatedAt, &data.ModifiedAt)
		if errs != nil {
			return result, errs
		}

		result = append(result, data)
	}

	return result, nil
}

func (r *EventRepository) GetEventById(eventId string) (structs.Event, error) {
	var result structs.Event
	query := `select * from events where event_id = $1`

	err := database.DBConn.QueryRow(query, eventId).Scan(&result.Id, &result.EventId, &result.UserId, &result.Title, &result.Description, &result.Location, &result.EventDate, &result.EventTypeId, &result.CreatedAt, &result.ModifiedAt)
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
		var errs = rows.Scan(&data.Id, &data.EventId, &data.UserId, &data.Title, &data.Description, &data.Location, &data.EventDate, &data.EventTypeId, &data.CreatedAt, &data.ModifiedAt)
		if errs != nil {
			return result, errs
		}

		result = append(result, data)
	}

	return result, nil
}

func (r *EventRepository) UpdateEvent(eventId string, event structs.Event) error {
	query := `update events set title=$1, description=$3, location=$4, event_date=$5, event_type_id=$6, modified_at=$7  where event_id=$2`

	res, err := database.DBConn.Exec(query, event.Title, eventId, event.Description, event.Location, event.EventDate, event.EventTypeId, time.Now())
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("gagal update event")
	}

	return nil
}

func (r *EventRepository) DeleteEvent(eventId string) error {
	query := `delete from events where event_id=$1`

	res, err := database.DBConn.Exec(query, eventId)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("gagal menghapus event")
	}

	return nil
}
