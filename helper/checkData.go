package helper

import (
	"database/sql"
	"fmt"

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

func IsTicketExists(event structs.Ticket) bool {
	sqlStatement := `SELECT id from tickets where name = $1 returning id`

	var id int
	err := database.DBConn.QueryRow(sqlStatement, event.Name).Scan(&id)

	return err == sql.ErrNoRows
}

func IsReviewExists(review structs.Review) bool {
	sqlStatement := `SELECT id from reviews where user_id = $1 and event_id=$2 returning id`

	var id int
	err := database.DBConn.QueryRow(sqlStatement, review.UserId, review.EventId).Scan(&id)

	return err == sql.ErrNoRows
}

func IsOrderExists(orderId int) bool {
	sqlStatement := `SELECT id from orders where id=$1 returning id`

	var order structs.Order
	var id int
	err := database.DBConn.QueryRow(sqlStatement, order.Id).Scan(&id)

	return err != sql.ErrNoRows
}

func IsOrderAlreadyCheckIn(orderItemId int) (bool, error) {
	sqlStatement := `SELECT id,is_check_in from order_items where id=$1`

	var id int
	var alreadyCheckin bool
	err := database.DBConn.QueryRow(sqlStatement, orderItemId).Scan(&id, &alreadyCheckin)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	if alreadyCheckin {
		return true, nil
	}

	return false, nil
}
