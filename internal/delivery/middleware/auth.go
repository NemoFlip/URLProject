package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func RequireAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		splitedToken := strings.Split(token, " ")
		if len(splitedToken) != 2 || splitedToken[0] != "Bearer" {
			ctx.Writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Println(splitedToken[1])
		ctx.Next()
	}
}
