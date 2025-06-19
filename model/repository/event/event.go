package event

import "github.com/tiketin-management-api-with-go/structs"

type EventRepositoryInterface interface {
	CreateEvent(e structs.Event) error
	GetAllEvent() ([]structs.Event, error)
	GetEventById(id string) (structs.Event, error)
	GetEventByUser(userId int) ([]structs.Event, error)
	UpdateEvent(id string, e structs.Event) error
	DeleteEvent(id string) error
}

type EventRepository struct {
}
