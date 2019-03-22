package categoryrepository

import (
	"api/db"
	"api/models"
)

// GetByIDWithPosts - get category with related posts
func GetByIDWithPosts(id string) models.Category {
	var category models.Category
	db.DB.Where("id = ?", id).Preload("Posts").First(&category)
	return category
}

// Store - store category in db
func Store(category models.Category) (models.Category, error) {
	if err := db.DB.Create(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}
