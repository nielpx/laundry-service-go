package main

import (
	"golang-gorm-gin/internal/router"
	"golang-gorm-gin/pkg"
	"github.com/sirupsen/logrus"
)

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
	r.Run()
}