package eventtype

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/tiketin-management-api-with-go/helper"
	"github.com/tiketin-management-api-with-go/model/repository/eventtype"
	"github.com/tiketin-management-api-with-go/structs"
)

func NewEventTypeService(repo eventtype.EventTypeRepositoryInterface) EventTypeServiceInterface {
	return &EventTypeService{
		repo: repo,
	}
}

func (s *EventTypeService) CreateEventType(ctx *gin.Context) error {
	var newEventType structs.EventType
	err := ctx.ShouldBindJSON(&newEventType)
	if err != nil {
		result := helper.ValidationCheck(err)
		if result != nil {
			return result
		}
		return err
	}

	exists := helper.IsEventTypeExists(newEventType)
	if exists {
		return errors.New("gagal menambahkan event type, data sudah ada")
	}

	err = s.repo.CreateEventType(newEventType)
	if err != nil {
		return err
	}

	return nil
}

func (s *EventTypeService) GetAllEventType(ctx *gin.Context) ([]structs.EventType, error) {
	data, err := s.repo.GetAllEventType()

	if err != nil {
		return []structs.EventType{}, err
	}

	return data, nil
}

func (s *EventTypeService) GetEventByEventTypeId(ctx *gin.Context) ([]structs.Event, error) {
	id := ctx.Param("id")
	data, err := s.repo.GetEventByEventTypeId(id)

	if err != nil {
		return []structs.Event{}, err
	}

	return data, nil
}

func (s *EventTypeService) UpdateEventType(ctx *gin.Context) error {
	var updateEventType structs.EventType
	err := ctx.ShouldBindJSON(&updateEventType)
	if err != nil {
		result := helper.ValidationCheck(err)
		if result != nil {
			return result
		}
		return err
	}

	id := ctx.Param("id")
	err = s.repo.UpdateEventType(id, updateEventType)

	if err != nil {
		return err
	}

	return nil
}

func (s *EventTypeService) DeleteEventType(ctx *gin.Context) error {
	id := ctx.Param("id")
	err := s.repo.DeleteEventType(id)

	if err != nil {
		return err
	}

	return nil
}
