package request

import (
	"github.com/gin-gonic/gin"
)

func decode[T any](ctx *gin.Context) (T, error) {
	var requestStruct T
	if err := ctx.BindJSON(&requestStruct); err != nil {
		return requestStruct, err
	}
	return requestStruct, nil
}
