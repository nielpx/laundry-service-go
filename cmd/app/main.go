package main

import (
	"golang-gorm-gin/internal/database"
	"golang-gorm-gin/internal/handler"
	"golang-gorm-gin/internal/repository"
	"golang-gorm-gin/internal/usecase"
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db, err := database.ConnectDatabase()
	if err != nil{
		log.Fatalf("Tidak bisa terhubung ke database %v", err)
	}

	layananRepo := repository.NewLayananRepository(db)
	LayananUsecase := usecase.NewLayananUsecase(layananRepo)
	layananHandler := handler.NewLayananHandler(LayananUsecase)

	r.GET("/laundry-services", layananHandler.ListLayanan)
	r.GET("/laundry-services/:id", layananHandler.GetLayanan)
	r.POST("/laundry-services", layananHandler.CreateLayanan)
	r.PUT("/laundry-services/:id", layananHandler.UpdateLayanan)
	r.DELETE("/laundry-services/:id", layananHandler.DeleteLayanan)

	r.Run()
}