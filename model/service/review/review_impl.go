package review

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tiketin-management-api-with-go/helper"
	"github.com/tiketin-management-api-with-go/model/repository/review"
	"github.com/tiketin-management-api-with-go/structs"
)

func NewReviewService(repo review.ReviewRepositoryInterface) ReviewServiceInterface {
	return &ReviewService{
		repo: repo,
	}
}

func (s *ReviewService) CreateReview(ctx *gin.Context) error {
	var newReview structs.Review
	err := ctx.ShouldBindJSON(&newReview)
	if err != nil {
		result := helper.ValidationCheck(err)
		if result != nil {
			return result
		}
		return err
	}

	dataJwt, err := helper.GetJwtData(ctx)
	if err != nil {
		return err
	}

	newReview.UserId = dataJwt.UserId

	eventId, _ := strconv.Atoi(ctx.Param("event_id"))
	newReview.EventId = eventId

	exists := helper.IsReviewExists(newReview)
	if exists {
		return errors.New("gagal menambahkan review, data sudah ada")
	}

	status, err := s.repo.HasUserCompletedEvent(newReview.UserId, newReview.EventId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("kamu belum pernah membeli tiket event ini, tidak bisa membuat review")
		}
		return err
	}

	if status != "complete" {
		return errors.New("kamu belum menyelesaikan event ini, tidak bisa membuat review")
	}

	err = s.repo.CreateReview(newReview)
	if err != nil {
		return err
	}

	return nil
}

func (s *ReviewService) UpdateReview(ctx *gin.Context) error {
	var updateReview structs.Review
	err := ctx.ShouldBindJSON(&updateReview)
	if err != nil {
		result := helper.ValidationCheck(err)
		if result != nil {
			return result
		}
		return err
	}

	dataJwt, err := helper.GetJwtData(ctx)
	if err != nil {
		return err
	}

	updateReview.UserId = dataJwt.UserId

	eventId, _ := strconv.Atoi(ctx.Param("event_id"))
	updateReview.EventId = eventId

	reviewId, _ := strconv.Atoi(ctx.Param("review_id"))

	err = s.repo.UpdateReview(reviewId, updateReview)
	if err != nil {
		return err
	}

	return nil
}

func (s *ReviewService) GetAllReviewInEvent(ctx *gin.Context) ([]structs.Review, error) {
	eventId, _ := strconv.Atoi(ctx.Param("event_id"))

	data, err := s.repo.GetAllReviewInEvent(eventId)
	if err != nil {
		return []structs.Review{}, err
	}

	return data, nil
}
