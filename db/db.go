package db

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {
    db, err := gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/golang-api?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic(err.Error())
    }
    return db
}