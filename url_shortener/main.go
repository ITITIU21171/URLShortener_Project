package main

import (
	"log"

	"github.com/ITITIU21171/url-shortener/database"
	"github.com/ITITIU21171/url-shortener/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	database.ConnectRedis()

	r := gin.Default()

	r.StaticFile("/", "./index.html")

	r.POST("/shorten", routes.ShortenURL)
	r.GET("/shortURL", routes.RedirectURL)

	log.Println("port 8080")
	r.Run(":8080")

}
