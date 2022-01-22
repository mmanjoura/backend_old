package json

import (
	"encoding/json"

	"github.com/mmanjoura/nomad/backend"
	"github.com/mmanjoura/nomad/backend/shipping"
)

type Shipping struct{}

func (r *Shipping) DecodeShipping(input []byte) (*shipping.Shipping, error) {
	attr := &shipping.Shipping{}
	if err := json.Unmarshal(input, attr); err != nil {
		return nil, backend.Errorf(backend.EINTERNAL, "serializer.Shipping.DecodeAttriubte")
	}
	return attr, nil
}

func (r *Shipping) EncodeShipping(input *shipping.Shipping) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, backend.Errorf(backend.EINTERNAL, "serializer.Shipping.EncodeShipping")
	}
	return rawMsg, nil
}
