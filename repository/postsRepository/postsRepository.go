package postsRepository

import (
    "api/db"
    "api/models"
)

func GetAll() []models.Post {
    db := db.Connect()
    defer db.Close()
    var posts []models.Post
    db.Find(&posts)
    return posts
}

func GetById(id string) models.Post {
    db := db.Connect()
    defer db.Close()
    var post models.Post
    db.Where("id = ?", id).First(&post)
    return post
}

func Create(post models.Post) (models.Post, error) {
    db := db.Connect()
    defer db.Close()
    if err := db.Create(&post).Error; err != nil {
        return post, err
    }
    return post, nil
}

func Delete(id string) bool {
    db := db.Connect()
    defer db.Close()
    var post models.Post
    db.Where("id = ?", id).First(&post)
    if post.ID != 0 {
        db.Unscoped().Delete(&post)
        return true
    }    
    return false
}