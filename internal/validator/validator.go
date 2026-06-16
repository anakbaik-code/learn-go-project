package validator

import (
	"github.com/go-playground/validator/v10"
)

func NewValidator() *validator.Validate {
	r := validator.New()
	return r
}
