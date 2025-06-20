package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourname/blog-backend/internal/model"
)

func (h *ArticleHandler) GetZennArticles(c *gin.Context) {
	var zennArticles []model.ZennArticle

	if err := h.DB.Order("published_at desc").Find(&zennArticles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, zennArticles)
}

func (h *ArticleHandler) CreateZennArticle(c *gin.Context) {
	var zennArticle model.ZennArticle

	if err := c.ShouldBindJSON(&zennArticle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, zennArticle)
}
