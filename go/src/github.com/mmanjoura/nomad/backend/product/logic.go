package product

import (
	"context"
	"errors"

	"github.com/mmanjoura/nomad/backend"
)

var (
	ErrProductNotFound = errors.New("Product Not Found")
	ErrProductInvalid  = errors.New("Product Invalid")
)

type productService struct {
	productRepo ProductRepository
}

func NewProductService(productRepo ProductRepository) ProductService {
	return &productService{
		productRepo,
	}
}

func (r *productService) FindOne(ctx context.Context, productId int) (*Product, error) {
	return r.productRepo.FindOne(ctx, productId)
}

func (r *productService) FindAll(ctx context.Context, productId int, filter backend.Filter) (*ProductData, int, error) {
	return r.productRepo.FindAll(ctx, productId, filter)
}

func (r *productService) Create(ctx context.Context, productId int, product *Product) error {
	return r.productRepo.Create(ctx, productId, product)
}

func (r *productService) Update(ctx context.Context, productId int, attr Product) (*Product, error) {
	return r.productRepo.Update(ctx, productId, attr)
}

func (r *productService) Delete(ctx context.Context, productId int) error {
	return r.productRepo.Delete(ctx, productId)
}

// Validate returns an error if the product contains invalid fields.
// This only performs basic validation.
func (u *Product) Validate() error {
	if u.Name == "" {
		return errors.New("Product Invalid")
	}
	return nil
}
