package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nakamurakzz/go-gin-api/config"
	"github.com/nakamurakzz/go-gin-api/db"
	"github.com/nakamurakzz/go-gin-api/site"
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

func getSite(c *gin.Context) {
	var data interface{}

	message := "ok"
	status := http.StatusOK

	// if err := errorFunc(); err != nil {
	// 	message = err.Error()
	// 	data = nil
	// 	status = http.StatusInternalServerError
	// }

	// SiteRepositoryを使ってDBからデータを取得する
	siteRepository := site.SiteRepositoryImpl{}
	sites, err := siteRepository.FindAll()

	if err != nil {
		message = err.Error()
		data = nil
		status = http.StatusInternalServerError
	} else {
		data = sites
	}

	c.IndentedJSON(status, gin.H{
		"message": message,
		"data":    data,
	})
}

// 10回に1度はエラーになる関数
func errorFunc() error {
	if rand.Intn(10) == 0 {
		return nil
	}
	return errors.New("2回に1度発生するERROR")
}

func main() {
	cnf, err := config.New()
	db.DBOpen(cnf.DB_HOST, cnf.DB_PORT, cnf.DB_USER, cnf.DB_PASS, cnf.DB_DATABASE)
	defer db.DBClose()
	r := setupRouter()
	if err != nil {
		panic(err)
	}
	r.Run(fmt.Sprintf(":%d", cnf.Port))
}
