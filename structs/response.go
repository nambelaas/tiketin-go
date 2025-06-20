package structs

// ResponseError
// @Description menampilkan response error dengan pesan dan detail,
// @Description bisa berupa error dari binding atau error lainnya
// @Description jika detail tidak ada, maka akan diabaikan
type ErrorStruct struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Detail  interface{} `json:"detail" binding:"omitempty"`
} // @name ResponseError

// ResponseSuccess
// @Description menampilkan response sukses hanya dengan pesan
type SuccessStruct struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
} // @name ResponseSuccess

// ResponseSuccessWithData
// @Description menampilkan response sukses dengan pesan dan data
// @Description bisa berupa struct, slice, atau map
// @Description jika data tidak ada, maka akan diabaikan
type SuccessStructWithData struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
} // @name ResponseSuccessWithData

// ResponseSuccessToken
// @Description menampilkan response sukses dengan token
// @Description biasanya digunakan untuk otentikasi atau otorisasi
type SuccessTokenStruct struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
} // @name ResponseSuccessToken
