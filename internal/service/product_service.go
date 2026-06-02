package service

import (
	"context"
	"database/sql"
	"errors"
	"go-dbsqlc/internal/domain"
	"go-dbsqlc/internal/repository"
	"go-dbsqlc/internal/validator"
)

type ProductService interface {
	CreateProduct(ctx context.Context, name,price domain.Product) (domain.Product, error)
	GetProduct(ctx context.Context, id int64) (domain.Product, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) ProductService {
	return &productService{
		repo: r,
	}
}

func (s *productService) CreateProduct(ctx context.Context, name,price domain.Product) (domain.Product, error) {
	if err := validator.ValidateCreateProduct(name,price); err != nil {
		return domain.Product{}, err
	}

	product, err := s.repo.Create(ctx, name,price)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil

}

func (s *productService) GetProduct(ctx context.Context, id int64) (domain.Product, error) {
	if err := validator.ValidateGetProductByID(id); err != nil {
		return domain.Product{}, err
	}

	product, err := s.repo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Product{}, errors.New("Product Not Found")
		}
	}
	return product, nil
}
