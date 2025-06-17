package eventtype

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tiketin-management-api-with-go/helper"
	"github.com/tiketin-management-api-with-go/model/repository/event"
	"github.com/tiketin-management-api-with-go/structs"
)

func NewEventService(repo event.EventRepositoryInterface) EventServiceInterface {
	return &EventService{
		repo: repo,
	}
}

func (s *EventService) CreateEvent(ctx *gin.Context) error {
	var newEvent structs.Event
	err := ctx.ShouldBindJSON(&newEvent)
	if err != nil {
		result := helper.ValidationCheck(err)
		if result != nil {
			return result
		}
		return err
	}

	exists := helper.IsEventExists(newEvent)
	if exists {
		return errors.New("gagal menambahkan event, data sudah ada")
	}

	if newEvent.EventDate.Before(time.Now()) {
		return errors.New("waktu mulai event hanya bisa diatas waktu saat ini")
	}

	dataJwt, err := helper.GetJwtData(ctx)
	if err != nil {
		return err
	}

	newEvent.UserId = dataJwt.UserId

	err = s.repo.CreateEvent(newEvent)
	if err != nil {
		return err
	}

	return nil
}

func (s *EventService) GetAllEvent(ctx *gin.Context) ([]structs.Event, error) {
	data, err := s.repo.GetAllEvent()

	if err != nil {
		return []structs.Event{}, err
	}

	return data, nil
}

func (s *EventService) GetEventById(ctx *gin.Context) (structs.Event, error) {
	id := ctx.Param("event_id")
	data, err := s.repo.GetEventById(id)

	if err != nil {
		return structs.Event{}, err
	}

	return data, nil
}

func (s *EventService) GetEventByUser(ctx *gin.Context) ([]structs.Event, error) {
	dataJwt, err := helper.GetJwtData(ctx)
	if err != nil {
		return []structs.Event{}, err
	}

	data, err := s.repo.GetEventByUser(dataJwt.UserId)

	if err != nil {
		return []structs.Event{}, err
	}

	return data, nil
}

func (s *EventService) UpdateEvent(ctx *gin.Context) error {
	var updateEvent structs.Event
	err := ctx.ShouldBindJSON(&updateEvent)
	if err != nil {
		result := helper.ValidationCheck(err)
		if result != nil {
			return result
		}
		return err
	}

	if updateEvent.EventDate.Before(time.Now()) {
		return errors.New("waktu mulai event hanya bisa diatas waktu saat ini")
	}

	dataJwt, err := helper.GetJwtData(ctx)
	if err != nil {
		return err
	}

	updateEvent.UserId = dataJwt.UserId

	id := ctx.Param("event_id")
	err = s.repo.UpdateEvent(id, updateEvent)

	if err != nil {
		return err
	}

	return nil
}

func (s *EventService) DeleteEvent(ctx *gin.Context) error {
	id := ctx.Param("event_id")
	err := s.repo.DeleteEvent(id)

	if err != nil {
		return err
	}

	return nil
}
