package helper

import (
	"database/sql"

	"github.com/tiketin-management-api-with-go/database"
	"github.com/tiketin-management-api-with-go/structs"
)

func IsUserExists(user structs.Users) bool {
	sqlStatement := `SELECT id from users where email = $1 returning id`

	var id int
	err := database.DBConn.QueryRow(sqlStatement, user.Email).Scan(&id)

	return err == sql.ErrNoRows
}
