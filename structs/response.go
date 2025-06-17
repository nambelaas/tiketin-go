package structs

type ErrorStruct struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Detail  interface{} `json:"detail"`
}

type SuccessStruct struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type SuccessTokenStruct struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}
