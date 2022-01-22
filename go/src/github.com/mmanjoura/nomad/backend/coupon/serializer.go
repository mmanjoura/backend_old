package coupon

type CouponSerializer interface {
	DecodeCoupon(input []byte) (*Coupon, error)
	EncodeCoupon(input *Coupon) ([]byte, error)
}
