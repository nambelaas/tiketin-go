package review

import "github.com/tiketin-management-api-with-go/structs"

type ReviewRepositoryInterface interface {
	CreateReview(review structs.Review) error
	UpdateReview(id int, review structs.Review) error
	GetAllReviewInEvent(eventId int) ([]structs.Review, error)
	HasUserCompletedEvent(userId, eventId int) (bool, error)
}

type ReviewRepository struct {
	
}
