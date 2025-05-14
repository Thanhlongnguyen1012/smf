package handlers

import (
	"net/http"
	"smf/models"
	"smf/services"

	"github.com/gin-gonic/gin"
)

// gán enponit và hàm xử lý
func CreateSMContextHandler(c *gin.Context) {
	// khởi tạo biến request gửi bản tin đi
	var request models.SMContextCreateData
	//lấy file JSON ở thân POST để gán vào struct request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	// xử lý logic
	status, response := services.CreateSMContext(request)
	//trả phần response về cho client để người dùng có thể biết kết quả request
	c.JSON(status, response)
}
