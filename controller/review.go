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

//	@summary		Create Review
//	@description	Menambahkan data review untuk event
//	@description	* Hanya bisa diakses oleh user
//	@tags			Review
//	@accept			json
//	@produce		json
//	@Param			Authorization	header		string					true	"Bearer token"
//	@Param			review			body		structs.Review			true	"Review Data"
//	@Success		200				{object}	structs.SuccessStruct	"Berhasil menambahkan review"
//	@Failure		400				{object}	structs.ErrorStruct		"Gagal menambahkan review"
//	@Router			/api/events/:event_id/reviews/create [post]
func CreateReviewHandle(ctx *gin.Context) {
	err := reviewService.CreateReview(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponse(ctx, "berhasil menambahkan review")
}

//	@summary		List Review
//	@description	Mendapatkan semua daftar review untuk event
//	@tags			Review
//	@accept			json
//	@produce		json
//	@Success		200	{object}	structs.SuccessStructWithData	"Berhasil mendapatkan data review"
//	@Failure		400	{object}	structs.ErrorStruct				"Gagal mendapatkan data review"
//	@Router			/api/events/:event_id/reviews/list [get]
func GetAllReviewEventHandle(ctx *gin.Context) {
	data, err := reviewService.GetAllReviewInEvent(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponseWithData(ctx, "berhasil mendapatkan data review", data)
}

//	@summary		Update Review
//	@description	Memperbarui data review untuk event
//	@description	* Hanya bisa diakses oleh user
//	@tags			Review
//	@accept			json
//	@produce		json
//	@Param			Authorization	header		string					true	"Bearer token"
//	@Param			review			body		structs.Review			true	"Review Data"
//	@Success		200				{object}	structs.SuccessStruct	"Berhasil memperbarui review"
//	@Failure		400				{object}	structs.ErrorStruct		"Gagal memperbarui review"
//	@Router			/api/events/:event_id/reviews/:review_id/update [put]
func UpdateReviewHandle(ctx *gin.Context) {
	err := reviewService.UpdateReview(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponse(ctx, "berhasil update review")
}
