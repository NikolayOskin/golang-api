package migrations

import (
	"api/db"
	"api/models"
)

// Migrate - migrations
func Migrate() {
	db.DB.DropTableIfExists("posts")
	db.DB.DropTableIfExists("categories")
	db.DB.AutoMigrate(&models.Post{})
	db.DB.AutoMigrate(&models.Category{})
}
