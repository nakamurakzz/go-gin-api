package site

import (
	"errors"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nakamurakzz/go-gin-api/util"
)

var siteUsecase = SiteUsecaseImpl{}

func GetSite(c *gin.Context) {
	var data interface{}

	message := "ok"
	status := http.StatusOK

	// if err := errorFunc(); err != nil {
	// 	message = err.Error()
	// 	data = nil
	// 	status = http.StatusInternalServerError
	// }

	// SiteRepositoryを使ってDBからデータを取得する
	sites, err := siteUsecase.FindAll()

	if err != nil {
		util.HttpErrorResponse(c, http.StatusInternalServerError, err)
	}
	data = sites

	// application/jsonでレスポンスを返す
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

func GetSiteById(c *gin.Context) {
	message := "ok"
	status := http.StatusOK
	var data interface{}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.HttpErrorResponse(c, http.StatusInternalServerError, err)
	}

	site, err := siteUsecase.FindByID(id)

	if err != nil {
		util.HttpErrorResponse(c, http.StatusInternalServerError, err)
	}

	data = site

	c.IndentedJSON(status, gin.H{
		"message": message,
		"data":    data,
	})
}

type SiteCreateRequestBody struct {
	Name string `json:"name"`
}

func PostSite(c *gin.Context) {
	message := "ok"
	status := http.StatusOK

	var site SiteCreateRequestBody
	if err := c.BindJSON(&site); err != nil {
		util.HttpErrorResponse(c, http.StatusInternalServerError, err)
	}

	// SiteRepositoryを使ってDBにデータを保存する
	data, err := siteUsecase.Create(site)
	if err != nil {
		util.HttpErrorResponse(c, http.StatusInternalServerError, err)
	}

	c.IndentedJSON(status, gin.H{
		"message": message,
		"data":    data,
	})
}
