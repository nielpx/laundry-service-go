package main

import (
	"golang-gorm-gin/internal/router"
	"golang-gorm-gin/pkg"

	"github.com/sirupsen/logrus"
	_ "golang-gorm-gin/cmd/app/docs"
)


// @title Golang CRUD API
// @version 1.0
// @description This is a sample server for managing products.
// @host localhost:8080
// @BasePath /
func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
		FullTimestamp: true,
	})
	pkg.LoadEnv()
}


func main(){
	r := router.InitializeRouter()
	logrus.Info("Succesfully connected")
	r.Run(":8080")
}