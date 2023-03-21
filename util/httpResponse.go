package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HttpErrorResponse(c *gin.Context, statusCode int, err error) {
	c.IndentedJSON(statusCode, gin.H{
		"message": err.Error(),
		"data":    nil,
	})
}

func HttpNotFoundResponse(c *gin.Context) {
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "該当するデータが見つかりませんでした",
		"data":    nil,
	})
}
