package util

import (
	"belajar-go/src/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseOk(c *gin.Context, message string, totalData int, data any) {
	c.JSON(http.StatusOK, dto.Response{
		Status:    "success",
		Message:   message,
		TotalData: totalData,
		Data:      data,
	})
}

func ResponseError(c *gin.Context, code int, message string) {
	c.JSON(code, dto.ResponseError{
		Status:  "error",
		Message: message,
	})
}
