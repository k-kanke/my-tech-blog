package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourname/blog-backend/internal/model"
	"gorm.io/gorm"
)

type ArticleHandler struct {
	DB *gorm.DB
}

func NewArticleHandler(db *gorm.DB) *ArticleHandler {
	return &ArticleHandler{DB: db}
}

func (h *ArticleHandler) GetArticles(c *gin.Context) {
	var articles []model.Article
	if err := h.DB.Find(&articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, articles)
}
