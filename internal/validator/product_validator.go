package validator

import (
	"errors"
	"go-dbsqlc/internal/handler/dto"

	"github.com/go-playground/validator/v10"
)

func ValidateProductId(v *validator.Validate, id int64) error {
	if err := v.Var(id, "gt=0"); err != nil {
		return errors.New("id must grater than 0")
	}
	return nil
}

// func ValidateCreateProduct(v *validator.Validate, product domain.Product) error {
// 	if err := v.Var(product.Name, "required"); err != nil {
// 		return errors.New("name must fill")
// 	}
// 	if err := v.Var(product.Price, "gt=0"); err != nil {
// 		return errors.New("price must not null")
// 	}
// 	return nil
// }

func ValidateCreateProductNested(v *validator.Validate, input dto.CreateProductNestedRequest) error {
	r := v.Struct(input)
	return r
}
