package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "test bro"})
	})

	v1 := router.Group("services/v1/users")
	{
		v1.GET("/test", getAllData)
	}

	router.Run(":3000")

}

type (
	usersModel struct {
	}
)

func getAllData(context *gin.Context) {
	// fmt.Fprintf(w, "Test")
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "No todo found!"})
	return
}
