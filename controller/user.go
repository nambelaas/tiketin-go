package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiketin-management-api-with-go/helper"
	"github.com/tiketin-management-api-with-go/middleware"
	repository "github.com/tiketin-management-api-with-go/model/repository/user"
	service "github.com/tiketin-management-api-with-go/model/service/user"
)

var (
	userRepo    = repository.NewUserRepository()
	userService = service.NewUserService(userRepo)
)

func RegisterUserHandle(ctx *gin.Context) {
	err := userService.RegisterUser(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponse(ctx, "berhasil mendaftarkan user")
}

func LoginUserHandle(ctx *gin.Context) {
	user, err := userService.LoginUser(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token := middleware.GenerateJwtToken(user)

	helper.PrintTokenResponse(ctx, token)
}

func GetUserHandle(ctx *gin.Context) {
	user,err := userService.GetUser(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponseWithData(ctx, "berhasil mendapatkan user", user)
}