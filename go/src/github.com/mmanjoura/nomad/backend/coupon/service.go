package coupon

import (
	"context"

	"github.com/mmanjoura/nomad/backend"
)

type CouponService interface {
	Create(ctx context.Context, couponId int, coupon *Coupon) error
	FindAll(ctx context.Context, couponId int, filter backend.Filter) ([]*Coupon, int, error)
	FindOne(ctx context.Context, couponId int) (*Coupon, error)
	Update(ctx context.Context, couponId int, attr Coupon) (*Coupon, error)
	Delete(ctx context.Context, couponId int) error
}
