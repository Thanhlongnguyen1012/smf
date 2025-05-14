package handlers

import (
	"net/http"
	"smf/models"
	"smf/services"

	"github.com/gin-gonic/gin"
)

func CreateSMContextHandler(c *gin.Context) {
	var request models.SMContextCreateData
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	status, response := services.CreateSMContext(request)
	c.JSON(status, response)
}
