package service

import (
	"context"
	"database/sql"
	"errors"
	"go-dbsqlc/db"
	"go-dbsqlc/internal/domain"
	"go-dbsqlc/internal/repository"
	"log/slog"
)

type ProductService interface {
	GetProduct(ctx context.Context, id int64) (domain.Product, error)
	CreateProduct(ctx context.Context, product domain.Product) (domain.Product, error)
	ListProduct(ctx context.Context) ([]domain.Product, error)
	UpdateProduct(ctx context.Context, id int64, product db.UpdateProductParams) error
	DeleteProduct(ctx context.Context, id int64) error
}
type productService struct {
	repo repository.ProductRepository
}

func NewProductService(logger *slog.Logger, r repository.ProductRepository) ProductService {
	return &productService{
		repo: r,
	}
}

func (s *productService) GetProduct(ctx context.Context, id int64) (domain.Product, error) {
	product, err := s.repo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Product{}, errors.New("product not found")
		}
		return domain.Product{}, err
	}
	return product, nil
}
func (s *productService) CreateProduct(ctx context.Context, product domain.Product) (domain.Product, error) {
	product, err := s.repo.Create(ctx, product)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Product{}, errors.New("product not found")
		}
		return domain.Product{}, err
	}
	return product, nil
}
func (s *productService) ListProduct(ctx context.Context) ([]domain.Product, error) {
	products, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}
func (s *productService) UpdateProduct(ctx context.Context, id int64, product db.UpdateProductParams) error {
	products := domain.Product{
		ID:    id,
		Name:  product.Name,
		Price: product.Price,
	}

	return s.repo.Update(ctx, products)
}
func (s *productService) DeleteProduct(ctx context.Context, id int64) error {
	return s.DeleteProduct(ctx, id)
}
