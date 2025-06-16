package user

import (
	"github.com/gin-gonic/gin"
	"github.com/tiketin-management-api-with-go/structs"
	"github.com/tiketin-management-api-with-go/model/repository/user"
)

type UserServiceInterface interface {
	RegisterUser(ctx *gin.Context) error
	LoginUser(ctx *gin.Context) (structs.Users, error)
	GetUser(ctx *gin.Context) (structs.Users, error)
}

type UserService struct {
	repo user.UserRepositoryInterface
}