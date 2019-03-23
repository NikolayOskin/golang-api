package models

import (
	"github.com/jinzhu/gorm"
)

// Post GORM model
type Post struct {
	gorm.Model
	//Id string `json:"id,omitempty"`
	Slug       string `gorm:"unique;" json:"slug,omitempty"`
	Title      string `json:"title,omitempty"`
	Body       string `json:"body,omitempty"`
	CategoryID int    `json:"category_id,omitempty"`
	Category   Category
}
