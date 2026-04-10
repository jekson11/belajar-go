package dto

type Response struct {
	Status    string `json:"status" extensions:"x-order=0"`
	Message   string `json:"message" extensions:"x-order=1"`
	TotalData int    `json:"total_data" extensions:"x-order=2"`
	Data      any    `json:"data" extensions:"x-order=3"`
}

type ResponseError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
