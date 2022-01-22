package order

import (
	"context"
	"errors"

	"github.com/mmanjoura/nomad/backend"
)

var (
	ErrOrderNotFound = errors.New("Order Not Found")
	ErrOrderInvalid  = errors.New("Order Invalid")
)

type orderService struct {
	orderRepo OrderRepository
}

func NewOrderService(orderRepo OrderRepository) OrderService {
	return &orderService{
		orderRepo,
	}
}

func (r *orderService) FindOne(ctx context.Context, orderId int) (*Order, error) {
	return r.orderRepo.FindOne(ctx, orderId)
}

func (r *orderService) FindAll(ctx context.Context, orderId int, filter backend.Filter) ([]*Order, int, error) {
	return r.orderRepo.FindAll(ctx, orderId, filter)
}

func (r *orderService) Create(ctx context.Context, orderId int, order *Order) error {
	return r.orderRepo.Create(ctx, orderId, order)
}

func (r *orderService) Update(ctx context.Context, orderId int, attr Order) (*Order, error) {
	return r.orderRepo.Update(ctx, orderId, attr)
}

func (r *orderService) Delete(ctx context.Context, orderId int) error {
	return r.orderRepo.Delete(ctx, orderId)
}

// Validate returns an error if the order contains invalid fields.
// This only performs basic validation.
func (u *Order) Validate() error {
	if u.ShippingAddress.Zip == "" {
		return errors.New("Order Invalid")
	}
	return nil
}
