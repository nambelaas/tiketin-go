package user

import "github.com/tiketin-management-api-with-go/structs"

type UserRepositoryInterface interface {
	RegisterUser(user structs.Users) error
	LoginUser(email string, password string) (structs.Users, error)
	GetUser(id int) (structs.Users, error)
}

type UserRepository struct {
}
