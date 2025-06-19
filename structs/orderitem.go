package structs

import "time"

type OrderItem struct {
	Id           int       `json:"id"`
	OrderId      int       `json:"order_id"`
	TicketTypeId int       `json:"ticket_type_id"`
	Quantity     int       `json:"quantity"`
	QrCodeUrl    string    `json:"qr_code_url"`
	IsCheckIn    bool      `json:"is_check_in"`
	CreatedAt    time.Time `json:"created_at"`
	ModifiedAt   time.Time `json:"modified_at"`
}
