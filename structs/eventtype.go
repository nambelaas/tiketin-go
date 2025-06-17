package structs

import "time"

type EventType struct {
	Id         int        `json:"id"`
	Name       string     `json:"name" binding:"required"`
	CreatedAt  time.Time  `json:"created_at"`
	ModifiedAt *time.Time `json:"modified_at"`
}
