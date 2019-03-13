package db

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Open() (*gorm.DB, error) {
    var err error
    DB, err = gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/golang-api?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        return nil, err
    }
    return DB, nil
}
 
func Close() error {
    return DB.Close()
}