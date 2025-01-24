package request

import (
	"github.com/go-playground/validator/v10"
)

func isValid[T any](requestStruct T) error {
	validate := validator.New()
	return validate.Struct(requestStruct)
}
