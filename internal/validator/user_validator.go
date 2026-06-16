package validator

import (
	"errors"
	"fmt"
	"go-dbsqlc/internal/domain"
	"go-dbsqlc/internal/handler/dto"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ValidateCreateUser(v *validator.Validate, user domain.User) error {
	if err := v.Var(user.Name, "required,min=3,max50"); err != nil {
		return errors.New("name must fill and length must be between 3 and 50 character")
	}
	if err := v.Var(user.Email, "required,email"); err != nil {
		return errors.New("email must fill and must be a valid email format")
	}
	return nil
}

func ValidateGetUserByID(v *validator.Validate, id int64) error {
	if err := v.Var(id, "gt=0"); err != nil {
		return errors.New("id must grater than 0")
	}
	return nil
}

func ValidateUpdateUser(v *validator.Validate, user dto.UpdateUserRequest) error {
	return v.Struct(user)
}

func ValidateImage(
	v *validator.Validate,
	file multipart.File,
	header *multipart.FileHeader,
) error {
	// max size = 5 MB
	maxSize := int64(5 << 20)

	// to string using fmt
	validationTag := fmt.Sprintf("lte=%d", maxSize)

	if err := v.Var(header.Size, validationTag); err != nil {
		return errors.New("file to large, maximum 5 MB")
	}

	buffer := make([]byte, 512)

	_, err := file.Read(buffer)
	if err != nil {
		return err
	}

	contentType := http.DetectContentType(buffer)
	if err := v.Var(contentType, "oneof=image/jpeg image/png"); err != nil {
		return errors.New("invalid file type, only jpeg and png are allowed")
	}

	file.Seek(0, io.SeekStart)

	return nil
}
