package main

import (
	"net/http"

	"github.com/go-mall/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/config-read", func(c *gin.Context) {
		database := config.Database
		c.JSON(http.StatusOK, gin.H{
			"type":     database.Type,
			"max_life": database.MaxLifeTime,
		})
	})
	r.Run(":8080") // listen and serve on
}
