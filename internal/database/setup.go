package database

import (
	"golang-gorm-gin/internal/models"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() (*gorm.DB, error) {
	database, err := gorm.Open(mysql.Open(os.Getenv("ADDRESS")))
	if err != nil{
		logrus.Error("Failed to open connection")
		panic(err)
	}

	database.AutoMigrate(&models.Product{})
	database.AutoMigrate(&models.User{})

	DB = database
	return database, nil
}
