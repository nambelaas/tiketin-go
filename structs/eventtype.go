package structs

import "time"

// EventType
// @Description Fields untuk jenis event
type EventType struct {
	// Id dari jenis event
	Id int `json:"id"`
	// Nama jenis event
	Name string `json:"name" binding:"required"`
	// Waktu pembuatan jenis event
	CreatedAt time.Time `json:"created_at"`
	// Waktu terakhir jenis event diubah
	ModifiedAt *time.Time `json:"modified_at"`
} // @name EventType
