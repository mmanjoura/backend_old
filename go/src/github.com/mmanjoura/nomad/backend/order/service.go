package order

import (
	"context"

	"github.com/mmanjoura/nomad/backend"
)

type OrderService interface {
	Create(ctx context.Context, orderId int, order *Order) error
	FindAll(ctx context.Context, orderId int, filter backend.Filter) ([]*Order, int, error)
	FindOne(ctx context.Context, orderId int) (*Order, error)
	Update(ctx context.Context, orderId int, attr Order) (*Order, error)
	Delete(ctx context.Context, orderId int) error
}
