package product

import (
	"context"

	"github.com/mmanjoura/nomad/backend"
)

type ProductRepository interface {
	Create(ctx context.Context, productId int, product *Product) error
	FindAll(ctx context.Context, productId int, filter backend.Filter) (*ProductData, int, error)
	FindOne(ctx context.Context, productId int) (*Product, error)
	Update(ctx context.Context, productId int, cat Product) (*Product, error)
	Delete(ctx context.Context, productId int) error
}
