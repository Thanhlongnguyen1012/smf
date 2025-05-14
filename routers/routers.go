package routers

import (
	"smf/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/smf")
	{
		api.POST("/sm-context", handlers.CreateSMContextHandler)
	}
	return r
}
