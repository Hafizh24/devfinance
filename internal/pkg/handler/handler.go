package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hafizh24/devfinance/internal/pkg/reason"
	"github.com/hafizh24/devfinance/internal/pkg/validator"
)

func ResponseSuccess(ctx *gin.Context, statusCode int, message string, data interface{}) {
	resp := ResponseBody{
		Status:  "success",
		Message: message,
		Data:    data,
	}

	ctx.JSON(statusCode, resp)
}
func ResponseError(ctx *gin.Context, statusCode int, message string) {
	resp := ResponseBody{
		Status:  "error",
		Message: message,
	}

	ctx.JSON(statusCode, resp)
}

func BindAndCheck(ctx *gin.Context, data interface{}) bool {
	err := ctx.ShouldBind(data)
	if err != nil {
		ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return true
	}

	isError := validator.Check(data)
	if isError {
		ResponseError(ctx, http.StatusUnprocessableEntity, reason.RequestFormError)
		return true
	}

	return false
}
