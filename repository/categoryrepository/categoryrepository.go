package categoryrepository

import (
	"api/db"
	"api/models"
)

var category models.Category

// GetByID - get category
func GetByID(id string) models.Category {
	db.DB.Where("id = ?", id).First(&category)
	return category
}

// GetByIDWithPosts - get category with related posts
func GetByIDWithPosts(id string) models.Category {
	db.DB.Where("id = ?", id).Preload("Posts").First(&category)
	return category
}

// Store - store category in db
func Store(category *models.Category) (*models.Category, error) {
	if err := db.DB.Create(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

// UpdateByID - update existed category
func UpdateByID(id string, newCategory *models.Category) (models.Category, error) {
	db.DB.Where("id = ?", id).First(&category)
	category.Slug = newCategory.Slug
	category.Title = newCategory.Title
	if err := db.DB.Save(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

// DeleteByID - delete existed category
func DeleteByID() {

}
