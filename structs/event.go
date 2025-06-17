package structs

import "time"

type Event struct {
	Id          int        `json:"id"`
	EventId     *string    `json:"event_id"`
	UserId      int        `json:"user_id" binding:"required"`
	Title       string     `json:"title" binding:"required"`
	Description *string    `json:"description"`
	Location    string     `json:"location" binding:"required"`
	EventDate   time.Time  `json:"event_date" binding:"required"`
	EventTypeId int        `json:"event_type_id" binding:"required"`
	CreatedAt   time.Time  `json:"created_at"`
	ModifiedAt  *time.Time `json:"modified_at"`
}
