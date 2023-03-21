package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
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

	r.Use(cors.Default())

	api := r.Group("/api/")
	{
		api.GET("/sites", site.GetSite)
		api.GET("/sites/:id", site.GetSiteById)
		api.POST("/sites", site.PostSite)
		api.PATCH("/sites", site.GetSite) // TODO: 修正
		api.DELETE("/sites/:id", site.DeleteSite)
	}

	return r
}

func main() {
	cnf, err := config.New()
	db.DBOpen()
	defer db.DBClose()
	r := setupRouter()
	if err != nil {
		panic(err)
	}
	r.Run(fmt.Sprintf(":%d", cnf.Port))
}
