package validator

import (
	"errors"
	"go-dbsqlc/internal/domain"
)


func ValidateCreateProduct(name,price domain.Product) error {
	if name.Name == "" {
		return errors.New("name wajib")
	}
	if price.Price <= 0 {
		return errors.New("price invalid")
	}
	return nil
}

func ValidateGetProductByID(id int64) error {
	if id <= 0 {
		return errors.New("invalid product id")
	}

	return nil
}