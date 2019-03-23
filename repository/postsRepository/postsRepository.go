package postsRepository

import (
	"api/db"
	"api/models"
)

var post models.Post

func GetAll() []models.Post {
	var posts []models.Post
	db.DB.Find(&posts)
	return posts
}

func GetById(id string) models.Post {
	db.DB.Where("id = ?", id).First(&post)
	return post
}

func GetByIdWithCategory(id string) models.Post {
	db.DB.Where("id = ?", id).Preload("Category").First(&post)
	return post
}

func Create(post models.Post) (models.Post, error) {
	if err := db.DB.Create(&post).Error; err != nil {
		return post, err
	}
	return post, nil
}

func Delete(id string) bool {
	db.DB.Where("id = ?", id).First(&post)
	if post.ID != 0 {
		db.DB.Unscoped().Delete(&post)
		return true
	}
	return false
}
