package structs

import "time"

// Ticket
// @Description Fields untuk tiket event
type Ticket struct {
	// Id dari tiket
	Id         int        `json:"id"`
	// Id event yang tiket ini miliki
	EventId    int        `json:"event_id"`
	// Nama dari tiket
	Name       string     `json:"name" binding:"required"`
	// Harga dari tiket
	Price      float32    `json:"price" binding:"required"`
	// Jumlah tiket yang tersedia
	Quota      int        `json:"quota" binding:"required"`
	// Waktu pembuatan tiket
	CreatedAt  time.Time  `json:"created_at"`
	// Waktu terakhir tiket diubah
	ModifiedAt *time.Time `json:"modified_at"`
}// @name Ticket

// OrderTicket
// @Description Fields untuk tiket dalam order
type OrderTicket struct {
	// Id dari tiket yang dipesan
	TicketTypeId int `json:"ticket_type_id" binding:"required,number"`
	// Jumlah tiket yang dipesan
	Quantity     int `json:"quantity" binding:"required,number,min=1"`
}// @name OrderTicket
