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

func IsEventTypeExists(eventType structs.EventType) bool {
	sqlStatement := `SELECT id from event_types where name = $1 returning id`

	var id int
	err := database.DBConn.QueryRow(sqlStatement, eventType.Name).Scan(&id)

	return err == sql.ErrNoRows
}

func IsEventExists(event structs.Event) bool {
	sqlStatement := `SELECT id from events where title = $1 returning id`

	var id int
	err := database.DBConn.QueryRow(sqlStatement, event.Title).Scan(&id)

	return err == sql.ErrNoRows
}
