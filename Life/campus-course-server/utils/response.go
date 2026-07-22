package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(context *gin.Context, message string, data interface{}) {
	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": message,
		"data":    data,
	})
}

func Fail(context *gin.Context, statusCode int, message string) {
	context.JSON(statusCode, gin.H{
		"code":    statusCode,
		"message": message,
		"data":    nil,
	})
}
