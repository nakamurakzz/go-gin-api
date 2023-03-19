package util

import "github.com/gin-gonic/gin"

func HttpErrorResponse(c *gin.Context, statusCode int, err error) {
	c.IndentedJSON(statusCode, gin.H{
		"message": err.Error(),
		"data":    nil,
	})
}
