package middleware

import (
	"fmt"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/hafizh24/devfinance/internal/pkg/handler"
	"github.com/hafizh24/devfinance/internal/pkg/reason"
)

func AuthorizationMiddleware(obj string, act string, enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sub, existed := ctx.Get("user_id")
		if !existed {
			handler.ResponseError(ctx, 401, reason.UserNotLogin)
		}

		ok, err := enforcer.Enforce(fmt.Sprint(sub), obj, act)

		if err != nil {
			handler.ResponseError(ctx, http.StatusInternalServerError, reason.ErrAuthorize)
			ctx.Abort()
			return
		}

		if !ok {
			handler.ResponseError(ctx, http.StatusForbidden, reason.NotAuthorized)
			ctx.Abort()
			return
		}
		ctx.Next()

	}
}
