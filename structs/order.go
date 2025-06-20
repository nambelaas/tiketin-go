package structs

import "time"

// Order
// @Description Fields untuk order
type Order struct {
	// Id dari order
	Id int `json:"id"`
	// Id user yang membuat order
	UserId int `json:"user_id"`
	// Id event yang dipesan
	EventId int `json:"event_id"`
	// Status dari order (new, paid, complete, cancelled)
	Status string `json:"status"`
	// Total harga dari order
	TotalPrice float32 `json:"total"`
	// Waktu pembayaran order
	PaidAt *time.Time `json:"paid_at" binding:"omitempty"`
	// Metode pembayaran order (cash, transfer, etc)
	PaymentMethod string `json:"payment_method"`
	// Waktu pembuatan order
	CreatedAt time.Time `json:"created_at"`
	// Waktu terakhir order diubah
	ModifiedAt *time.Time `json:"modified_at" binding:"omitempty"`
	// Daftar item dalam order
	OrderItem []OrderItem `json:"order_item" binding:"omitempty"`
} // @name Order

// OrderCreate
// @Description Fields untuk membuat order
type OrderCreate struct {
	// Id event yang dipesan
	EventId int `json:"event_id" binding:"required"`
	// Daftar tiket yang dipesan
	Ticket []OrderTicket `json:"ticket" binding:"required,dive"`
} // @name OrderCreate
