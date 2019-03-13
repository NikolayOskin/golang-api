package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	//Id string `json:"id,omitempty"`
    Slug string   `json:"slug,omitempty"`
    Title  string   `json:"title,omitempty"`
    Body  string   `json:"body,omitempty"`
}