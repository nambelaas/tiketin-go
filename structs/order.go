package structs

import "time"

type Order struct {
	Id            int         `json:"id"`
	UserId        int         `json:"user_id"`
	EventId       int         `json:"event_id"`
	Status        string      `json:"status"`
	TotalPrice    float32     `json:"total"`
	PaidAt        time.Time   `json:"paid_at" binding:"omitempty"`
	PaymentMethod string      `json:"payment_method"`
	CreatedAt     time.Time   `json:"created_at"`
	ModifiedAt    time.Time   `json:"modified_at" binding:"omitempty"`
	OrderItem     []OrderItem `json:"order_item" binding:"omitempty"`
}

type OrderCreate struct {
	EventId int           `json:"event_id" binding:"required"`
	Ticket  []OrderTicket `json:"ticket" binding:"required,dive"`
}
