package routers

import (
	"smf/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/nsmf-pdusession")
	{
		api.POST("/v1/sm-contexts/", handlers.CreateSMContextHandler)
	}
	return r
}
