package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("views/*")

	r.GET("/", mainPage())

	r.Run(":8080") // Explicitly specify port
}

func mainPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "KaryokeType",
		})
	}
}
