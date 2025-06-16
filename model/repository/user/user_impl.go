package user

import (
	"database/sql"
	"errors"

	"github.com/tiketin-management-api-with-go/database"
	"github.com/tiketin-management-api-with-go/structs"
	"golang.org/x/crypto/bcrypt"
)

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

func (u *UserRepository) RegisterUser(user structs.Users) error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	var res sql.Result

	if user.Role != "" {
		query := `Insert into users(name,email,password,role) values ($1,$2,$3,$4)`
		res, err = database.DBConn.Exec(query, user.Name, user.Email, password, user.Role)
	} else {
		query := `Insert into users(name,email,password) values ($1,$2,$3)`
		res, err = database.DBConn.Exec(query, user.Name, user.Email, password)
	}

	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("gagal menambahkan data user, data tidak valid")
	}

	return nil
}

func (u *UserRepository) LoginUser(email string, password string) (structs.Users, error) {
	var dataUser structs.Users
	query := `Select id, name, password, role from users where email = $1`

	err := database.DBConn.QueryRow(query, email).Scan(&dataUser.Id, &dataUser.Name, &dataUser.Password, &dataUser.Role)
	if err == sql.ErrNoRows {
		return dataUser, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dataUser.Password), []byte(password))
	if err != nil {
		return dataUser, errors.New("password tidak cocok dengan username ini")
	}

	return dataUser, nil
}

func (u *UserRepository) GetUser(id int) (structs.Users, error) {
	var dataUser structs.Users
	query := `Select * from users where id = $1`

	err := database.DBConn.QueryRow(query, id).Scan(&dataUser.Id, &dataUser.Name, &dataUser.Email, &dataUser.Password, &dataUser.Role, &dataUser.CreatedAt)
	if err == sql.ErrNoRows {
		return dataUser, err
	}

	return dataUser, nil
}
