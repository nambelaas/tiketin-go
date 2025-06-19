package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiketin-management-api-with-go/helper"
	repository "github.com/tiketin-management-api-with-go/model/repository/review"
	service "github.com/tiketin-management-api-with-go/model/service/review"
)

var (
	reviewRepository = repository.NewReviewRepository()
	reviewService    = service.NewReviewService(reviewRepository)
)

func CreateReviewHandle(ctx *gin.Context) {
	err := reviewService.CreateReview(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponse(ctx, "berhasil menambahkan review")
}

func GetAllReviewEventHandle(ctx *gin.Context) {
	data, err := reviewService.GetAllReviewInEvent(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponseWithData(ctx, "berhasil mendapatkan data review", data)
}

func UpdateReviewHandle(ctx *gin.Context) {
	err := reviewService.UpdateReview(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponse(ctx, "berhasil update review")
}
