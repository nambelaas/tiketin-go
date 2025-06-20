package structs

import "time"

// Users
// @Description Fields untuk user
type Users struct {
	// Id dari user
	Id int `json:"id"`
	// Nama user
	Name string `json:"name" binding:"required"`
	// Email user
	Email string `json:"email" binding:"required"`
	// Password user
	Password string `json:"password" binding:"required"`
	// Peran user (admin/user/organizer) jika kosong berarti perannya user
	Role string `json:"role"`
	// Waktu pembuatan user
	CreatedAt time.Time `json:"created_at" `
} // @name Users

// LoginRequest
// @Description Fields untuk request login
type LoginRequest struct {
	// Email user
	Email string `json:"email" binding:"required,email"`
	// Password user
	Password string `json:"password" binding:"required"`
} // @name LoginRequest
