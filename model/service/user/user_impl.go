package user

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/tiketin-management-api-with-go/helper"
	"github.com/tiketin-management-api-with-go/model/repository/user"
	"github.com/tiketin-management-api-with-go/structs"
)

func NewUserService(repo user.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) RegisterUser(ctx *gin.Context) error {
	var newUser structs.Users
	err := ctx.ShouldBindJSON(&newUser)
	if err != nil {
		result := helper.ValidationCheck(err)
		if result != nil {
			return result
		}
		return err
	}

	exists := helper.IsUserExists(newUser)
	if exists {
		return errors.New("gagal menambahkan user, data sudah ada")
	}

	err = u.repo.RegisterUser(newUser)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) LoginUser(ctx *gin.Context) (structs.Users, error) {
	var loginData structs.LoginRequest
	err := ctx.ShouldBindJSON(&loginData)
	if err != nil {
		result := helper.ValidationCheck(err)
		if result != nil {
			return structs.Users{}, result
		}
		return structs.Users{}, err
	}

	result, err := u.repo.LoginUser(loginData.Email, loginData.Password)
	if err != nil {
		return structs.Users{}, err
	}

	return result, nil
}

func (u *UserService) GetUser(ctx *gin.Context) (structs.Users, error) {
	dataJwt, err := helper.GetJwtData(ctx)
	if err != nil {
		return structs.Users{}, helper.EncodeError(err.Error())
	}

	result, err := u.repo.GetUser(dataJwt.UserId)
	if err != nil {
		return structs.Users{}, err
	}

	return result, nil
}
