package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func CheckAuthorization(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" || !strings.HasPrefix(token, "Bearer") {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you must be authorized for this action",
		})
		return
	}
	ctx.Next()
}
