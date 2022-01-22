package shipping

import (
	"context"

	"github.com/mmanjoura/nomad/backend"
)

type ShippingService interface {
	Create(ctx context.Context, shippingId int, shipping *Shipping) error
	FindAll(ctx context.Context, shippingId int, filter backend.Filter) ([]*Shipping, int, error)
	FindOne(ctx context.Context, shippingId int) (*Shipping, error)
	Update(ctx context.Context, shippingId int, attr Shipping) (*Shipping, error)
	Delete(ctx context.Context, shippingId int) error
}
