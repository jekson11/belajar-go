package rest

import (
	"belajar-go/src/service"
	"sync"

	"github.com/gin-gonic/gin"
)

// this singleton pattern to make sure InitRestHandler is called only during execution
var onceRestHandler = &sync.Once{}

type rest struct {
	//router HHPT
	gin  *gin.Engine
	svc  *service.Service
	port string
}

func InitRestHandler(svc *service.Service, port string) {
	onceRestHandler.Do(func() {
		e := &rest{
			gin:  gin.Default(),
			svc:  svc,
			port: port,
		}
		e.Serve()
	})
}

func (e *rest) Serve() {
	e.gin.GET("/users", e.ListUsers)

	if err := e.gin.Run(":" + e.port); err != nil {
		panic(err)
	}
}
