package eventtype

import "github.com/tiketin-management-api-with-go/structs"

type EventTypeRepositoryInterface interface {
	CreateEventType(et structs.EventType) error
	GetAllEventType() ([]structs.EventType, error)
	GetEventByEventTypeId(id string) ([]structs.Event, error)
	UpdateEventType(id string, et structs.EventType) error
	DeleteEventType(id string) error
}

type EventTypeRepository struct {
}
