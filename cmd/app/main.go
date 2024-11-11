package main

import (
	"golang-gorm-gin/internal/database"
	"golang-gorm-gin/internal/handler"
	"golang-gorm-gin/internal/middleware"
	"golang-gorm-gin/internal/repository"
	"golang-gorm-gin/internal/usecase"
	"golang-gorm-gin/pkg"
	"log"

	"github.com/gin-gonic/gin"
)


func main() {
	pkg.LoadEnv()

	r := gin.Default()
	db, err := database.ConnectDatabase()
	if err != nil{
		log.Fatalf("Tidak bisa terhubung ke database %v", err)
	}

	layananRepo := repository.NewLayananRepository(db)
	LayananUsecase := usecase.NewLayananUsecase(layananRepo)
	layananHandler := handler.NewLayananHandler(LayananUsecase)

	r.POST("/signin", handler.SignUp)
	r.POST("/login", handler.LogIn)

	r.GET("/laundry-services", middleware.RequireAuth, layananHandler.ListLayanan)
	r.GET("/laundry-services/:id", middleware.RequireAuth,layananHandler.GetLayanan)
	r.POST("/laundry-services",middleware.RequireAuth, layananHandler.CreateLayanan)
	r.PUT("/laundry-services/:id",middleware.RequireAuth, layananHandler.UpdateLayanan)
	r.DELETE("/laundry-services/:id", middleware.RequireAuth, layananHandler.DeleteLayanan)

	r.Run()
}