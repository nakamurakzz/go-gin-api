package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nakamurakzz/go-gin-api/config"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// UserAgentを取得するミドルウェア
	r.Use(func(c *gin.Context) {
		c.Next()
	})

	api := r.Group("/api/")
	{
		api.GET("/sites", getSite)
		api.POST("/sites", getSite)   // TODO: 修正
		api.PATCH("/sites", getSite)  // TODO: 修正
		api.DELETE("/sites", getSite) // TODO: 修正
	}

	return r
}

type site struct {
	id        int
	name      string
	isEnabled bool
}

func getSite(c *gin.Context) {
	data := []site{
		{
			id:        1,
			name:      "test",
			isEnabled: true,
		},
	}

	message := "ok"
	status := http.StatusOK

	if err := errorFunc(); err != nil {
		message = err.Error()
		data = nil
		status = http.StatusInternalServerError
	}

	c.IndentedJSON(status, gin.H{
		"message": message,
		"data":    data,
	})
}

// 2回に1度はエラーになる関数
func errorFunc() error {
	if rand.Intn(2) == 0 {
		return nil
	}
	return errors.New("2回に1度発生するERROR")
}

func main() {
	r := setupRouter()
	cnf, err := config.New()
	if err != nil {
		panic(err)
	}
	r.Run(fmt.Sprintf(":%d", cnf.Port))
}
