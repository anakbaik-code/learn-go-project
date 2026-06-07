package service

import (
	"context"
	"database/sql"
	"errors"
	"go-dbsqlc/internal/domain"
	"go-dbsqlc/internal/repository"
	"log/slog"
)

type ProductService interface {
	GetProduct(ctx context.Context,id int64)(domain.Product,error)
}
type productService struct {
	log *slog.Logger
	repo repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) ProductService {
	return  &productService{
		
		repo: r,
	}
}

func (s *productService) GetProduct(ctx context.Context,id int64)(domain.Product,error){

	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Product{}, errors.New("user not found")
		}

		return domain.Product{}, err
	}
	return user, nil
}