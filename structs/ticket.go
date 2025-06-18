package structs

import "time"

type Ticket struct {
	Id         int        `json:"id"`
	TicketId   string     `json:"ticket_id"`
	EventId    int        `json:"event_id"`
	Name       string     `json:"name" binding:"required"`
	Price      float32    `json:"price" binding:"required"`
	Quota      int        `json:"quota" binding:"required"`
	CreatedAt  time.Time  `json:"created_at"`
	ModifiedAt *time.Time `json:"modified_at"`
}
