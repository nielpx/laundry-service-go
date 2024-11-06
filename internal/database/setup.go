package database

import (
	"golang-gorm-gin/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() (*gorm.DB, error) {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_crud"))
	if err != nil{
		panic(err)
	}

	database.AutoMigrate(&models.Product{})

	DB = database
	return database, nil
}