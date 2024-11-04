package main

import (
	"golang-gorm-gin/internal/controllers"
	"golang-gorm-gin/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDatabase()


	r.GET("/laundry-services", controllers.Index)
	r.GET("/laundry-services/:id", controllers.Show)
	r.POST("/laundry-services", controllers.Create)
	r.PUT("/laundry-services/:id", controllers.Update)
	r.DELETE("/laundry-services/:id", controllers.Delete)

	r.Run()
}