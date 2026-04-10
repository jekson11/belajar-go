package dto

type Response struct {
	Data      any `json:"data" extensions:"x-order=0"`
	TotalData int `json:"totalData" extensions:"x-order=1"`
}

type ResponseError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
