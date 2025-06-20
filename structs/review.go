package structs

import "time"

// Review
// @Description Fields untuk review event
type Review struct {
	// Id dari review
	Id int `json:"id"`
	// Id user yang membuat review
	UserId int `json:"user_id"`
	// Id event yang direview
	EventId int `json:"event_id"`
	// Rating dari review (1-5)
	Rating int `json:"rating" binding:"required,number,min=1,max=5"`
	// Komentar dari review
	Comment string `json:"comment"`
	// Waktu pembuatan review
	CreatedAt time.Time `json:"created_at"`
	// Waktu terakhir review diubah
	ModifiedAt *time.Time `json:"modified_at"`
} // @name Review
