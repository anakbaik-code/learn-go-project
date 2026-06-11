package validator

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"go-dbsqlc/internal/domain"
)

func ValidateCreateProduct(v *validator.Validate, product domain.Product) error {
	if err := v.Var(product.Name, "required"); err != nil {
		return errors.New("name must fill")
	}
	if err := v.Var(product.Price, "gt=0"); err != nil {
		return errors.New("price must not null")
	}
	return nil
}
