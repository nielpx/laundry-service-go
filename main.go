package main

import (
	"golang-gorm-gin/models"
	"golang-gorm-gin/controllers/prdcontrol"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", prdcontrol.Index)
	r.GET("/api/products/:id", prdcontrol.Show)
	r.POST("/api/products", prdcontrol.Create)
	r.PUT("/api/products/:id", prdcontrol.Update)
	r.DELETE("/api/products/:id", prdcontrol.Delete)

	r.Run()
}