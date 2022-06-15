package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CheckAccessRequest struct {
	ID       string `json:"id"`
	Resource string `json:"resource"`
	Action   string `json:"action"`
}

type CheckAccessResp struct {
	HasAccess bool `json:"has_access"`
}

func CheckAccess(resource, action string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("X-API-Key")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, "")
			ctx.Abort()
			return
		}

		ca := &CheckAccessRequest{
			ID:       token,
			Resource: resource,
			Action:   action,
		}

		respBody, _ := json.Marshal(ca)
		bytes.NewReader(respBody)

		req, _ := http.NewRequest("POST", "http://authz:8081/v1/authz/check_access", bytes.NewReader(respBody))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Tenant", ctx.GetHeader("X-Tenant"))

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "")
			ctx.Abort()
			return
		}

		defer resp.Body.Close()
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "")
			ctx.Abort()
			return
		}

		cr := &CheckAccessResp{}
		err = json.Unmarshal(b, cr)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "")
			ctx.Abort()
			return
		}

		if !cr.HasAccess {
			ctx.JSON(http.StatusUnauthorized, "")
			ctx.Abort()
			return
		}
		ctx.Next()
		return
	}
}
