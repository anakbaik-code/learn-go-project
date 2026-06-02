package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go-dbsqlc/internal/domain"
	"go-dbsqlc/internal/repository"
	"go-dbsqlc/internal/validator"
	"io"
	"mime/multipart"
	"os"
)

type UserService interface {
	GetUser(ctx context.Context, id int64) (domain.User, error)
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	ListUsers(ctx context.Context) ([]domain.User, error)
	UpdateUser(ctx context.Context, id int64, req domain.UpdateUserRequest) error
	DeleteUser(ctx context.Context, id int64) error
	UploadAvatar(ctx context.Context, id int64, file multipart.File, header *multipart.FileHeader) (string, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{
		repo: r,
	}
}

func (s *userService) GetUser(ctx context.Context, id int64) (domain.User, error) {
	if err := validator.ValidateGetUserByID(id); err != nil {
		return domain.User{}, err
	}

	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, errors.New("user not found")
		}

		return domain.User{}, err
	}
	return user, nil
}

func (s *userService) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	if err := validator.ValidateCreateUser(user); err != nil {
		return domain.User{}, err
	}

	user, err := s.repo.Create(ctx, user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *userService) ListUsers(ctx context.Context) ([]domain.User, error) {
	users, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) UpdateUser(ctx context.Context, id int64, req domain.UpdateUserRequest) error {
	user := domain.User{
		ID:    id,
		Name:  req.Name,
		Email: req.Email,
	}
	return s.repo.Update(ctx, user)
}

func (s *userService) DeleteUser(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *userService) UploadAvatar(ctx context.Context, id int64, file multipart.File, header *multipart.FileHeader) (string, error) {
	filename := header.Filename
	if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
		return "", err
	}
	dst, err := os.Create("./uploads/" + filename)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	url := "/uploads/" + filename
	fmt.Println(url)
	
	err = s.repo.UpdateAvatar(ctx, domain.User{
		ID:        id,
		AvatarUrl: url,
	})
	if err != nil {
		return "", err
	}

	return url, nil
}
