package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	ua := ""

	// UserAgentを取得するミドルウェア
	r.Use(func(c *gin.Context) {
		ua = c.GetHeader("User-Agent")
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"ua": ua,
		})
	})
	return r
}

func main() {
	r:= setupRouter()
	r.Run(":8080")
}