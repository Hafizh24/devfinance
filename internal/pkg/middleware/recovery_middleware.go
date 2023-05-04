package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hafizh24/devfinance/internal/pkg/handler"
	"github.com/hafizh24/devfinance/internal/pkg/reason"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				handler.ResponseError(ctx, http.StatusInternalServerError, reason.InternalServerError)
			}
		}()

		ctx.Next()
	}
}
