package main

import (
	"golang-gorm-gin/internal/controllers"
	"golang-gorm-gin/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDatabase()


	r.GET("/api/products", api.Index)
	r.GET("/api/products/:id", api.Show)
	r.POST("/api/products", api.Create)
	r.PUT("/api/products/:id", api.Update)
	r.DELETE("/api/products/:id", api.Delete)

	r.Run()
}