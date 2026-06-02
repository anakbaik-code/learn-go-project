package repository

import (
	"context"
	"go-dbsqlc/db"
	"go-dbsqlc/internal/domain"
)

type ProductRepository interface {
	GetById(ctx context.Context, id int64) (domain.Product, error)
	Create(ctx context.Context, name,email domain.Product) (domain.Product, error)
}

type productRepository struct {
	db *db.Queries
}

func NewProductRepository(db *db.Queries) ProductRepository {
	return &productRepository{db: db}
}

func (r productRepository) Create(ctx context.Context, name,price domain.Product) (domain.Product, error) {
	result, err := r.db.CreateProduct(ctx, db.CreateProductParams{
		Name:  name.Name,
		Price: price.Price,
	})
	if err != nil {
		return domain.Product{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Product{}, err
	}

	p, err := r.db.GetProduct(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}

	return domain.Product{
		ID:    p.ID,
		Name:  p.Name,
		Price: p.Price,
	}, nil

}

func (r productRepository) GetById(ctx context.Context, id int64) (domain.Product, error) {
	p, err := r.db.GetProduct(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}
	return domain.Product{
		ID:    p.ID,
		Name:  p.Name,
		Price: p.Price,
	}, nil
}
