package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func PaginationMiddleware(defaultPage int, defaultPageSize int, defaultType string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		page, err := strconv.Atoi(ctx.Query("page"))
		if err != nil {
			page = defaultPage
		}
		pageSize, err := strconv.Atoi(ctx.Query("page_size"))
		if err != nil {
			pageSize = defaultPageSize
		}
		Type, exist := ctx.GetQuery("type")
		if !exist {
			Type = defaultType
		}

		ctx.Set("page", page)
		ctx.Set("page_size", pageSize)
		ctx.Set("type", Type)

		ctx.Next()
	}
}
