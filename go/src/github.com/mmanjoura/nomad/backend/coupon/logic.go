package coupon

import (
	"context"
	"errors"

	"github.com/mmanjoura/nomad/backend"
)

var (
	ErrCouponNotFound = errors.New("Coupon Not Found")
	ErrCouponInvalid  = errors.New("Coupon Invalid")
)

type couponService struct {
	couponRepo CouponRepository
}

func NewCouponService(couponRepo CouponRepository) CouponService {
	return &couponService{
		couponRepo,
	}
}

func (r *couponService) FindOne(ctx context.Context, couponId int) (*Coupon, error) {
	return r.couponRepo.FindOne(ctx, couponId)
}

func (r *couponService) FindAll(ctx context.Context, couponId int, filter backend.Filter) ([]*Coupon, int, error) {
	return r.couponRepo.FindAll(ctx, couponId, filter)
}

func (r *couponService) Create(ctx context.Context, couponId int, coupon *Coupon) error {
	return r.couponRepo.Create(ctx, couponId, coupon)
}

func (r *couponService) Update(ctx context.Context, couponId int, attr Coupon) (*Coupon, error) {
	return r.couponRepo.Update(ctx, couponId, attr)
}

func (r *couponService) Delete(ctx context.Context, couponId int) error {
	return r.couponRepo.Delete(ctx, couponId)
}

// Validate returns an error if the coupon contains invalid fields.
// This only performs basic validation.
func (u *Coupon) Validate() error {
	if u.Code == "" {
		return errors.New("Coupon Invalid")
	}
	return nil
}
