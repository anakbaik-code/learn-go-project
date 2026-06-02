package repository

import (
	"context"
	"database/sql"
	"go-dbsqlc/db"
	"go-dbsqlc/internal/domain"
)

type UserRepository interface {
	GetById(ctx context.Context, id int64) (domain.User, error)
	Create(ctx context.Context, user domain.User) (domain.User, error)
	List(ctx context.Context) ([]domain.User, error)
	Update(ctx context.Context, user domain.User) error
	Delete(ctx context.Context, id int64) error
	UpdateAvatar(ctx context.Context, user domain.User) error
}

type userRepository struct {
	db *db.Queries
}

func NewUserRepository(db *db.Queries) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetById(ctx context.Context, id int64) (domain.User, error) {
	u, err := r.db.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		AvatarUrl: u.AvatarUrl.String,
	}, nil
}

func (r *userRepository) Create(ctx context.Context, user domain.User) (domain.User, error) {
	result, err := r.db.CreateUser(ctx, db.CreateUserParams{
		Name:  user.Name,
		Email: user.Email,
	})
	if err != nil {
		return domain.User{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return domain.User{}, err
	}

	u, err := r.db.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}, nil
}

func (r *userRepository) List(ctx context.Context) ([]domain.User, error) {
	users, err := r.db.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	var result []domain.User
	for _, user := range users {
		result = append(result, domain.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return result, nil
}

func (r *userRepository) Update(ctx context.Context, user domain.User) error {
	return r.db.UpdateUser(ctx, db.UpdateUserParams{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})
}

func (r *userRepository) Delete(ctx context.Context, id int64) error {
	return r.db.DeleteUser(ctx, id)
}

func (r *userRepository) UpdateAvatar(ctx context.Context, user domain.User) error {
	arg := db.UpdateAvatarParams{
		ID: user.ID,
		AvatarUrl: sql.NullString{
			String: user.AvatarUrl,
			Valid:  user.AvatarUrl != "",
		},
	}
	return r.db.UpdateAvatar(ctx, arg)
}
