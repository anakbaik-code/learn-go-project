package validator

import (
	"errors"
	"go-dbsqlc/internal/handler/dto"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateProductId(v *validator.Validate, id int64) error {
	if err := v.Var(id, "gt=0"); err != nil {
		return errors.New("id must grater than 0")
	}
	return nil
}

func ValidateCreateProduct(v *validator.Validate, req dto.CreateProductNestedRequest) []dto.ValidationError {
	err := v.Struct(req)
	if err == nil {
		return nil
	}

	var validationErrors validator.ValidationErrors
	if !errors.As(err, &validationErrors) {
		return []dto.ValidationError{
			{
				Field:   "general",
				Message: err.Error(),
			},
		}
	}

	var result []dto.ValidationError

	for _, fieldErr := range validationErrors {
		var msg string
		switch fieldErr.Tag() {
		case "required":
			msg = "field is required"
		case "valid_sku":
			msg = "invalid sku format"
		case "gt":
			msg = "must be greater than zero"
		default:
			msg = "invalid value"
		}
		// Namespace menghasilkan: "CreateProductNestedRequest.Items[2].Price"
		fullNamespace := fieldErr.Namespace()
		jsonPath := fullNamespace

		// Potong nama struct utamanya di depan
		if parts := strings.SplitN(fullNamespace, ".", 2); len(parts) > 1 {
			jsonPath = parts[1]
		}

		result = append(result, dto.ValidationError{
			// Kita ubah path-nya jadi lowercase sesuai kebutuhan kamu, misal: "items[2].price"
			Field:   strings.ToLower(jsonPath),
			Message: msg,
		})
	}

	return result

}
