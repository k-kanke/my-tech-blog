package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourname/blog-backend/internal/db"
	"github.com/yourname/blog-backend/internal/handler"
	"github.com/yourname/blog-backend/internal/model"
)

func main() {
	dbConn := db.InitDB()
	dbConn.AutoMigrate(&model.Article{}) // Articleテーブル自動生成

	articleHandler := handler.NewArticleHandler(dbConn)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	r.GET("/articles", articleHandler.GetArticles)
	r.POST("/articles", articleHandler.CreateArticle)
	r.GET("/articles/:id", articleHandler.GetArticleByID)
	r.PUT("/articles/:id", articleHandler.UpdateArticle)
	r.DELETE("/articles/:id", articleHandler.DeleteArticle)

	r.Run(":8080")
}
