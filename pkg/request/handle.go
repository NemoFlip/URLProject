package request

import (
	"github.com/gin-gonic/gin"
)

func HandleBody[T any](ctx *gin.Context) (*T, error) {
	requestStruct, err := decode[T](ctx)
	if err != nil {
		return nil, err
	}
	if err = isValid[T](requestStruct); err != nil {
		return nil, err
	}

	return &requestStruct, nil
}
