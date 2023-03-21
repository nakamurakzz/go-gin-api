package site

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nakamurakzz/go-gin-api/util"
)

var siteUsecase = SiteUsecase{}

func GetSite(c *gin.Context) {
	var data interface{}

	message := "ok"
	status := http.StatusOK

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

func GetSiteById(c *gin.Context) {
	message := "ok"
	status := http.StatusOK
	var data interface{}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.HttpErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	site, err := siteUsecase.FindByID(id)

	if err != nil {
		util.HttpErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	if site.ID == 0 {
		log.Println("koko")
		util.HttpNotFoundResponse(c)
		return
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
		util.HttpErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	// SiteRepositoryを使ってDBにデータを保存する
	data, err := siteUsecase.Create(site)
	if err != nil {
		util.HttpErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(status, gin.H{
		"message": message,
		"data":    data,
	})
}

func DeleteSite(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.HttpErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	site, err := siteUsecase.FindByID(id)
	if err != nil {
		util.HttpErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	if site.ID == 0 {
		util.HttpNotFoundResponse(c)
		return
	}

	id, err = siteUsecase.Delete(id)
	if err != nil {
		util.HttpErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	message := "ok"
	status := http.StatusNoContent

	c.IndentedJSON(status, gin.H{
		"message": message,
		"data":    nil,
	})
}
