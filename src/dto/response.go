package dto

type Response struct {
	Data      any `json:"data" extensions:"x-order=0"`
	TotalData int `json:"totalData" extensions:"x-order=1"`
}

type ResponseError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type UserFilter struct {
	Name     string `form:"name"`
	Username string `form:"username"`
	Email    string `form:"email"`
	SortBy   string `form:"sort_by"`
	SortDir  string `form:"sort_dir"`
	Page     int    `form:"page" binding:"min=1"`
	Limit    int    `form:"limit" binding:"min=1"`
}
