package structs

import "time"

// Event
// @Description Fields untuk event
type Event struct {
	// Id dari event
	Id          int        `json:"id"`
	// Id user pembuat event
	UserId      int        `json:"user_id"`
	// nama event
	Title       string     `json:"title" binding:"required"`
	// deskripsi event
	Description *string    `json:"description"`
	// lokasi event
	Location    string     `json:"location" binding:"required"`
	// tanggal event
	EventDate   time.Time  `json:"event_date" binding:"required"`
	// id jenis event
	EventTypeId int        `json:"event_type_id" binding:"required"`
	// status event (open/closed)
	Status      string     `json:"status" binding:"required,oneof=open closed"`
	// waktu pembuatan event
	CreatedAt   time.Time  `json:"created_at"`
	// waktu terakhir event diubah
	ModifiedAt  *time.Time `json:"modified_at"`
}// @name Event
