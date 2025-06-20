package model

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Content   string
	IsPublic  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type ZennArticle struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	URL         string
	PublishedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
