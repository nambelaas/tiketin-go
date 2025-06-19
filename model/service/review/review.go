package review

import (
	"github.com/gin-gonic/gin"
	"github.com/tiketin-management-api-with-go/model/repository/review"
	"github.com/tiketin-management-api-with-go/structs"
)

type ReviewServiceInterface interface {
	CreateReview(ctx *gin.Context) error
	UpdateReview(ctx *gin.Context) error
	GetAllReviewInEvent(ctx *gin.Context) ([]structs.Review, error)
}

type ReviewService struct {
	repo review.ReviewRepositoryInterface
}
