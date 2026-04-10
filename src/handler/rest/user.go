package rest

import (
	"belajar-go/src/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (e *rest) ListUsers(c *gin.Context) {
	users, err := e.svc.User.ListAllDataUser()
	if err != nil {
		util.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	util.ResponseOk(c, "Data Fetch successfully", len(users), users)
}
