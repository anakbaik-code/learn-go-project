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
	GetProduct(ctx context.Context, id int64) (domain.Product, error)
}
type productService struct {
	log  *slog.Logger
	repo repository.ProductRepository
}

func NewProductService(logger *slog.Logger, r repository.ProductRepository) ProductService {
	return &productService{
		log:  logger.With("component", "product_service"),
		repo: r,
	}
}

func (s *productService) GetProduct(ctx context.Context, id int64) (domain.Product, error) {
	product, err := s.repo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Product{}, errors.New("product not found")
		}

		slog.Error("database error occurred while fetching product",
			"error", err,
			"product_id", id)
		return domain.Product{}, err
	}
	return product, nil
}
