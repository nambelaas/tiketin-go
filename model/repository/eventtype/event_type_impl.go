package eventtype

import (
	"errors"
	"time"

	"github.com/tiketin-management-api-with-go/database"
	"github.com/tiketin-management-api-with-go/structs"
)

func NewEventTypeRepository() EventTypeRepositoryInterface {
	return &EventTypeRepository{}
}

func (r *EventTypeRepository) CreateEventType(eventType structs.EventType) error {
	query := `insert into event_types (name) values ($1)`

	res, err := database.DBConn.Exec(query, eventType.Name)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("gagal menambahkan event type")
	}

	return nil
}

func (r *EventTypeRepository) GetAllEventType() ([]structs.EventType, error) {
	var result []structs.EventType
	query := `select * from event_types`

	rows, err := database.DBConn.Query(query)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		var data = structs.EventType{}
		var errs = rows.Scan(&data.Id, &data.Name, &data.CreatedAt, &data.ModifiedAt)
		if errs != nil {
			return result, errs
		}

		result = append(result, data)
	}

	return result, nil
}

func (r *EventTypeRepository) GetEventByEventTypeId(id string) ([]structs.Event, error) {
	var result []structs.Event
	query := `select e.id, e.user_id, e.title, e.description, e.location, e.event_date, e.created_at, e.modified_at from events e 
	join event_types et on e.event_type_id = et.id 
	where et.id = $1`

	rows, err := database.DBConn.Query(query, id)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		var data = structs.Event{}
		var errs = rows.Scan(&data.Id, &data.UserId, &data.Title, &data.Description, &data.Location, &data.EventDate, &data.EventTypeId, &data.CreatedAt, &data.ModifiedAt)
		if errs != nil {
			return result, errs
		}

		result = append(result, data)
	}

	return result, nil
}

func (r *EventTypeRepository) UpdateEventType(id string, eventType structs.EventType) error {
	query := `update event_types set name=$1, modified_at=$2 where id=$3`

	res, err := database.DBConn.Exec(query, eventType.Name, time.Now(), id)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("gagal update event type")
	}

	return nil
}

func (r *EventTypeRepository) DeleteEventType(id string) error {
	query := `delete from event_types where id=$1`

	res, err := database.DBConn.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("gagal menghapus event type")
	}

	return nil
}
