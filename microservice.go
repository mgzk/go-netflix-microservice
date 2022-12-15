package main

import (
	"github.com/gin-gonic/gin"
	"go-netflix-microservice/datasource"
	"go-netflix-microservice/film"
	"go-netflix-microservice/ping"
)

func main() {
	datasource.InitDatabase()

	router := gin.Default()
	router.GET("/ping", ping.Ping)
	router.GET("/films", film.Find)
	router.GET("/films/:id", film.FindById)
	router.Run()
}
