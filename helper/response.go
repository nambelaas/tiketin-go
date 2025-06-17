package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiketin-management-api-with-go/structs"
)

func PrintErrorResponse(ctx *gin.Context, code int, message string) {
	dataError := structs.ErrorStruct{
		Success: false,
		Message: message,
	}

	ctx.AbortWithStatusJSON(code, dataError)
}

func PrintErrorResponseWithDetail(ctx *gin.Context, code int, message string, detail interface{}) {
	dataError := structs.ErrorStruct{
		Success: false,
		Message: message,
		Detail:  detail,
	}

	ctx.AbortWithStatusJSON(code, dataError)
}

func PrintSuccessResponse(ctx *gin.Context, message string) {
	dataSuccess := structs.SuccessStruct{
		Success: true,
		Message: message,
	}

	ctx.JSON(http.StatusOK, dataSuccess)
}

func PrintSuccessResponseWithData(ctx *gin.Context, message string, data interface{}) {
	dataSuccess := structs.SuccessStructWithData{
		Success: true,
		Message: message,
		Data:    data,
	}

	ctx.JSON(http.StatusOK, dataSuccess)
}

func PrintTokenResponse(ctx *gin.Context, token string) {
	dataSuccess := structs.SuccessTokenStruct{
		Success: true,
		Token:   token,
	}

	ctx.JSON(http.StatusOK, dataSuccess)
}
