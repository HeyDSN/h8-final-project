package helpers

import (
	"final-project/models"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, statusCode int, data interface{}, status string, err *models.Error) {
	ctx.JSON(statusCode, models.Response{
		Data:   data,
		Status: status,
		Error:  err,
	})
}
