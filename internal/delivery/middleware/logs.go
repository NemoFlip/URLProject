package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logging() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		log.Println(
			ctx.Request.Method,
			ctx.Request.URL.Path,
			fmt.Sprintf("%d miliseconds", time.Since(start).Milliseconds()))
	}
}
