package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nakamurakzz/go-gin-api/config"
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
			"ua":      ua,
		})
	})
	return r
}

func main() {
	r := setupRouter()
	cnf, err := config.New()
	if err != nil {
		panic(err)
	}
	r.Run(fmt.Sprintf(":%d", cnf.Port))
}
