package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourname/blog-backend/internal/db"
	"github.com/yourname/blog-backend/internal/model"
)

func main() {
	dbConn := db.InitDB()
	dbConn.AutoMigrate(&model.Article{}) // Articleテーブル自動生成

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	r.Run(":8080")
}
