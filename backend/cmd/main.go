package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourname/blog-backend/internal/db"
)

func main() {
	dbConn := db.InitDB()
	if dbConn != nil {
		log.Println("[aaaa]")
	}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	r.Run(":8080")
}
