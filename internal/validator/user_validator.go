package validator

import (
	"errors"
	"go-dbsqlc/internal/domain"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

func ValidateCreateUser(user domain.User) error {
	if user.Name == "" {
		return errors.New("name wajib")
	}

	if user.Email == "" {
		return errors.New("email wajib")
	}

	if !strings.Contains(user.Email, "@") {
		return errors.New("email invalid")
	}
	return nil
}

func ValidateGetUserByID(id int64) error {
	if id <= 0 {
		return errors.New("invalid user id")
	}
	return nil
}
func ValidateImage(
	file multipart.File,
	header *multipart.FileHeader,
) error {

	if header.Size > 5<<20 {
		return errors.New("file too large")
	}

	buffer := make([]byte, 512)

	_, err := file.Read(buffer)
	if err != nil {
		return err
	}

	contentType := http.DetectContentType(buffer)

	allowed := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
	}

	if !allowed[contentType] {
		return errors.New("invalid file type")
	}

	file.Seek(0, io.SeekStart)

	return nil
}
