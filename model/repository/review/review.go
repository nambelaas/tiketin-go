package review

import "github.com/tiketin-management-api-with-go/structs"

type ReviewRepositoryInterface interface {
	CreateReview(review structs.Review) error
	UpdateReview(id int, review structs.Review) error
	GetAllReviewInEvent(eventId int) ([]structs.Review, error)
	HasUserCompletedEvent(userId, eventId int) (string, error)
}

type ReviewRepository struct {
	
}
