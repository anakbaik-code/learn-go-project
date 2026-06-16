package repository

import (
	"context"
	"go-dbsqlc/db"
	"go-dbsqlc/internal/domain"
)

type ProductRepository interface {
	GetById(ctx context.Context, id int64) (domain.Product, error)
	Create(ctx context.Context, product domain.Product) (domain.Product, error)
	List(ctx context.Context) ([]domain.Product, error)
	Update(ctx context.Context, product domain.Product) error
	Delete(ctx context.Context, id int64) error
}

type productRepository struct {
	db *db.Queries
}

func NewProductRepository(db *db.Queries) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetById(ctx context.Context, id int64) (domain.Product, error) {
	result, err := r.db.GetProduct(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}
	return domain.Product{
		ID:        result.ID,
		Name:      result.Name,
		Price:     result.Price,
		IsActive:  result.IsActive,
		SalePrice: result.SalePrice,
	}, nil
}

func (r *productRepository) Create(ctx context.Context, product domain.Product) (domain.Product, error) {
	result, err := r.db.CreateProduct(ctx, db.CreateProductParams{
		Name:      product.Name,
		Price:     product.Price,
		IsActive:  product.IsActive,
		SalePrice: product.SalePrice,
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
		ID:        p.ID,
		Name:      p.Name,
		Price:     p.Price,
		IsActive:  p.IsActive,
		SalePrice: p.SalePrice,
	}, nil
}

func (r *productRepository) List(ctx context.Context) ([]domain.Product, error) {
	products, err := r.db.ListProducts(ctx)
	if err != nil {
		return nil, err
	}
	var result []domain.Product
	for _, product := range products {
		result = append(result, domain.Product{
			ID:        product.ID,
			Name:      product.Name,
			Price:     product.Price,
			IsActive:  product.IsActive,
			SalePrice: product.SalePrice,
		})
	}
	return result, nil
}

func (r *productRepository) Update(ctx context.Context, product domain.Product) error {
	return r.db.UpdateProduct(ctx, db.UpdateProductParams{
		ID:        product.ID,
		Name:      product.Name,
		Price:     product.Price,
		IsActive:  product.IsActive,
		SalePrice: product.SalePrice,
	})
}

func (r *productRepository) Delete(ctx context.Context, id int64) error {
	return r.db.DeleteProduct(ctx, id)
}
