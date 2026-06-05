package repository

import (
	"context"
	"go-dbsqlc/db"
	"go-dbsqlc/internal/domain"
)

type ProductRepository interface {
	GetById(ctx context.Context,id int64) (domain.Product,error)
}

type productRepository struct {
	db *db.Queries
}

func NewProductRepository(db *db.Queries) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetById (ctx context.Context,id int64)(domain.Product,error){
	result ,err := r.db.GetProduct(ctx,id)
	if err != nil {
		return domain.Product{},err
	}
	return domain.Product{
		ID: result.ID,
		Name: result.Name,
		Price: result.Price,
	},nil
}