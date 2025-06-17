package event

import "github.com/tiketin-management-api-with-go/structs"

type EventRepositoryInterface interface {
	CreateEvent(e structs.Event) error
	GetAllEvent() ([]structs.Event, error)
	GetEventById(eventId string) (structs.Event, error)
	GetEventByUser(userId int) ([]structs.Event, error)
	UpdateEvent(eventId string, e structs.Event) error
	DeleteEvent(eventId string) error
}

type EventRepository struct {
}
