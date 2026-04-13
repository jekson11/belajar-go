package rest

import (
	"net/http"

	"belajar-go/src/dto"
	"belajar-go/src/util"

	"github.com/gin-gonic/gin"
)

func (e *rest) ListUsers(c *gin.Context) {
	var filter dto.UserFilter

	if err := c.ShouldBindQuery(&filter); err != nil {
		util.ResponseError(c, http.StatusBadRequest, "invalid parameter")
		return
	}

	users, err := e.svc.User.ListAllDataUser(filter)
	if err != nil {
		util.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	util.ResponseOk(c, len(users), users)
}
