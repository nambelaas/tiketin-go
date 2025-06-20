package review

import (
	"errors"
	"fmt"
	"time"

	"github.com/tiketin-management-api-with-go/database"
	"github.com/tiketin-management-api-with-go/structs"
)

func NewReviewRepository() ReviewRepositoryInterface {
	return &ReviewRepository{}
}

func (r *ReviewRepository) CreateReview(review structs.Review) error {
	query := `insert into reviews (user_id, event_id,rating,comment) values ($1,$2,$3,$4)`

	res, err := database.DBConn.Exec(query, review.UserId, review.EventId, review.Rating, review.Comment)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("gagal menambahkan review")
	}

	return nil
}

func (r *ReviewRepository) UpdateReview(id int, review structs.Review) error {
	query := `update reviews set rating=$2, comment=$3, modified_at=$4 where id=$1`

	res, err := database.DBConn.Exec(query, id, review.Rating, review.Comment, time.Now())
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("gagal update review")
	}

	return nil
}

func (r *ReviewRepository) GetAllReviewInEvent(eventId int) ([]structs.Review, error) {
	var result []structs.Review
	query := `select * from reviews where event_id=$1`

	rows, err := database.DBConn.Query(query, eventId)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		var data = structs.Review{}
		var err = rows.Scan(&data.Id, &data.UserId, &data.EventId, &data.Rating, &data.Comment, &data.CreatedAt, &data.ModifiedAt)
		if err != nil {
			return result, err
		}

		result = append(result, data)
	}

	return result, nil
}

func (r *ReviewRepository) HasUserCompletedEvent(userId int, eventId int) (bool, error) {
	var count int
	fmt.Println("userId:", userId, "eventId:", eventId)
	query := `SELECT COUNT(*) FROM orders WHERE user_id = $1 AND event_id = $2 AND status = $3`
	err := database.DBConn.QueryRow(query, userId, eventId, "complete").Scan(&count)
	if err != nil {
		fmt.Println("QUERY ERROR:", err)
		return false, err
	}
	return count > 0, nil
}
