package shipping

import (
	"context"
	"errors"

	"github.com/mmanjoura/nomad/backend"
)

var (
	ErrShippingNotFound = errors.New("Shipping Not Found")
	ErrShippingInvalid  = errors.New("Shipping Invalid")
)

type shippingService struct {
	shippingRepo ShippingRepository
}

func NewShippingService(shippingRepo ShippingRepository) ShippingService {
	return &shippingService{
		shippingRepo,
	}
}

func (r *shippingService) FindOne(ctx context.Context, shippingId int) (*Shipping, error) {
	return r.shippingRepo.FindOne(ctx, shippingId)
}

func (r *shippingService) FindAll(ctx context.Context, shippingId int, filter backend.Filter) ([]*Shipping, int, error) {
	return r.shippingRepo.FindAll(ctx, shippingId, filter)
}

func (r *shippingService) Create(ctx context.Context, shippingId int, shipping *Shipping) error {
	return r.shippingRepo.Create(ctx, shippingId, shipping)
}

func (r *shippingService) Update(ctx context.Context, shippingId int, attr Shipping) (*Shipping, error) {
	return r.shippingRepo.Update(ctx, shippingId, attr)
}

func (r *shippingService) Delete(ctx context.Context, shippingId int) error {
	return r.shippingRepo.Delete(ctx, shippingId)
}

// Validate returns an error if the shipping contains invalid fields.
// This only performs basic validation.
func (u *Shipping) Validate() error {
	if u.Name == "" {
		return errors.New("Shipping Invalid")
	}
	return nil
}
