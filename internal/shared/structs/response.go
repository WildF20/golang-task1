package structs

type SuccessResponse struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Data any `json:"data"`
}

type ErrorResponse struct {
	Status bool `json:"status"`
	Message string `json:"message"`
}