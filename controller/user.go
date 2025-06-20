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

//	@summary		Register User
//	@description	Endpoint untuk mendaftarkan user baru
//	@tags			User
//	@accept			json
//	@produce		json
//	@Param			user	body		structs.Users			true	"User Data"
//	@Success		200		{object}	structs.SuccessStruct	"Berhasil mendaftarkan user"
//	@Failure		400		{object}	structs.ErrorStruct		"Gagal mendaftarkan user"
//	@Router			/api/users/register [post]
func RegisterUserHandle(ctx *gin.Context) {
	err := userService.RegisterUser(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponse(ctx, "berhasil mendaftarkan user")
}

//	@summary		Login User
//	@description	Endpoint untuk login user
//	@tags			User
//	@accept			json
//	@produce		json
//	@Param			user	body		structs.LoginRequest		true	"User Data"
//	@Success		200		{object}	structs.SuccessTokenStruct	"Berhasil login"
//	@Failure		400		{object}	structs.ErrorStruct			"Gagal login"
//	@Router			/api/users/login [post]
func LoginUserHandle(ctx *gin.Context) {
	user, err := userService.LoginUser(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token := middleware.GenerateJwtToken(user)

	helper.PrintTokenResponse(ctx, token)
}

//	@summary		Get User
//	@description	Mengambil data user yang sedang login
//	@tags			User
//	@accept			json
//	@produce		json
//	@Security		BearerAuth
//	@Param			Authorization	header		string							true	"Bearer Token"
//	@Success		200				{object}	structs.SuccessStructWithData	"Berhasil mendapatkan user"
//	@Failure		400				{object}	structs.ErrorStruct				"Gagal mendapatkan user"
//	@Router			/api/users/me [get]
func GetUserHandle(ctx *gin.Context) {
	user, err := userService.GetUser(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.PrintSuccessResponseWithData(ctx, "berhasil mendapatkan user", user)
}
