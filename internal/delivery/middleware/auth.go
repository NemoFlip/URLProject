package middleware

import (
	"URLProject/configs"
	"URLProject/pkg/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	ContextEmailKey string = "email"
)

func RequireAuthorization(config *configs.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		fmt.Println(token)
		isValid, data := jwt.NewJWT(config.Auth.SecretKey).Parse(token)
		if !isValid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		fmt.Println(token)
		ctx.Set(ContextEmailKey, data.Email)
		ctx.Next()
	}
}
