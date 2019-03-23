package models

import (
	"github.com/jinzhu/gorm"
)

// Category GORM model
type Category struct {
	gorm.Model
	Slug  string `gorm:"unique;" json:"slug,omitempty"`
	Title string `json:"title,omitempty"`
	Posts []Post
}
