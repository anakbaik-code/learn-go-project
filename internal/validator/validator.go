package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidProductSKU(fl validator.FieldLevel) bool {
	sku := fl.Field().String()

	// Regex: ^[A-Z]{3} (3 huruf besar), - (tanda strip), [0-9]{4,} (minimal 4 digit angka)
	regex := regexp.MustCompile(`^[A-Z]{3}-[0-9]{4,}$`)

	return regex.MatchString(sku)
}

func NewValidator() *validator.Validate {
	v := validator.New()

	// Tag Alias
	v.RegisterAlias("varchar", "required,min=3,max=255")

	// custom validate
	v.RegisterValidation("valid_sku", ValidProductSKU)

	return v
}
