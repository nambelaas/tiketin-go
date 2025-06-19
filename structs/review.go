package structs

import "time"

type Review struct {
	Id         int        `json:"id"`
	UserId     int        `json:"user_id"`
	EventId    int        `json:"event_id"`
	Rating     int        `json:"rating" binding:"required,number,min=1,max=5"`
	Comment    string     `json:"comment"`
	CreatedAt  time.Time  `json:"created_at"`
	ModifiedAt *time.Time `json:"modified_at"`
}
