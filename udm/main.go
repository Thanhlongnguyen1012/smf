package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SmContextCreateData struct {
	Supi         string `json:"supi" gorm:"column:supi"`
	Dnn          string `json:"dnn" gorm:"column:dnn"`
	SNssai       string `json:"sNssai" gorm:"column:sNssai"`
	PduSessionId int32  `json:"PduSessionId" gorm:"column:pduSessionId"`
}
type NfProfile struct {
	NfType string `json:"nfType"`
}

func (NfProfile) TableName() string {
	return "user_data"
}
func (SmContextCreateData) TableName() string {
	return "user_data"
}
func main() {
	// connect MySQL
	dsn := "root:my-secret-pw@tcp(127.0.0.1:3306)/udm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database")
		return
	}

	r := gin.Default()
	r.GET("/nudm-sdm/v2/imsi-452040916843227/sm-data", func(c *gin.Context) {
		var data SmContextCreateData
		db.First(&data, "supi = ? ", "imsi-452040916843227")
		c.JSON(http.StatusCreated, &data)
	})
	// Endpoint GET /api/user
	// r.GET("/api/user", func(c *gin.Context) {
	// 	imsi := c.Query("imsi")
	// 	dnn := c.Query("dnn")

	// 	query := db.Model(&Subscriber{})

	// 	if imsi != "" {
	// 		query = query.Where("imsi = ?", imsi)
	// 	}
	// 	if dnn != "" {
	// 		query = query.Where("dnn = ?", dnn)
	// 	}

	// 	var users []Subscriber
	// 	if err := query.Find(&users).Error; err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	c.JSON(http.StatusCreated, gin.H{"users": users})
	// })

	r.Run(":9090")
}
