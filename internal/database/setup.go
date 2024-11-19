package database

import (
    "fmt"
    "time"
    "golang-gorm-gin/internal/models"
    "os"
    "github.com/sirupsen/logrus"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() (*gorm.DB, error) {
    var database *gorm.DB
    var err error
    maxRetries := 5
    retryDelay := 5 * time.Second

    for i := 0; i < maxRetries; i++ {
        database, err = gorm.Open(mysql.Open(os.Getenv("ADDRESS")))
        if err == nil {
            break
        }

        logrus.Warnf("Failed to connect to database, retrying... (%d/%d)", i+1, maxRetries)
        time.Sleep(retryDelay)
    }

    if err != nil {
        logrus.Error("Failed to open connection after retries")
        return nil, fmt.Errorf("could not connect to database: %v", err)
    }

    // Run migrations
    database.AutoMigrate(&models.Product{})
    database.AutoMigrate(&models.User{})

    DB = database
    return database, nil
}
