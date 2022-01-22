package json

import (
	"encoding/json"

	"github.com/mmanjoura/nomad/backend"
	"github.com/mmanjoura/nomad/backend/coupon"
)

type Coupon struct{}

func (r *Coupon) DecodeCoupon(input []byte) (*coupon.Coupon, error) {
	attr := &coupon.Coupon{}
	if err := json.Unmarshal(input, attr); err != nil {
		return nil, backend.Errorf(backend.EINTERNAL, "serializer.Coupon.DecodeAttriubte")
	}
	return attr, nil
}

func (r *Coupon) EncodeCoupon(input *coupon.Coupon) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, backend.Errorf(backend.EINTERNAL, "serializer.Coupon.EncodeCoupon")
	}
	return rawMsg, nil
}
