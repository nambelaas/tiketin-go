package structs

import "time"

// OrderItem
// @Description Fields untuk item dalam order
type OrderItem struct {
	// Id dari order item
	Id           int       `json:"id"`
	// Id dari order yang berisi item ini
	OrderId      int       `json:"order_id"`
	// Id dari tiket yang dipesan
	TicketTypeId int       `json:"ticket_type_id"`
	// Jumlah tiket yang dipesan
	Quantity     int       `json:"quantity"`
	// Barcode unik untuk order item yang digunakan untuk check-in
	QrCodeUrl    string    `json:"qr_code_url"`
	// Status check-in untuk order item (true jika sudah check-in)
	IsCheckIn    bool      `json:"is_check_in"`
	// Waktu pembuatan order item
	CreatedAt    time.Time `json:"created_at"`
	// Waktu terakhir order item diubah
	ModifiedAt   time.Time `json:"modified_at"`
}// @name OrderItem
