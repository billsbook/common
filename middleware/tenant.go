package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func TenantHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := ctx.GetHeader("X-Tenant")
		if t == "" {
			ctx.JSON(http.StatusBadRequest, "")
			ctx.Abort()
			return
		}

		s := strings.Split(t, ":")
		if len(s) != 2 {
			ctx.JSON(http.StatusBadRequest, "")
			ctx.Abort()
			return
		}

		c := ctx.Request.Context()
		c = context.WithValue(c, "tenant_id", s[0])
		c = context.WithValue(c, "db_user", s[0])
		c = context.WithValue(c, "db_pass", s[1])

		ctx.Request = ctx.Request.Clone(c)
		ctx.Next()
		return

	}
}
