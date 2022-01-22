package json

import (
	"encoding/json"

	"github.com/mmanjoura/nomad/backend"
	"github.com/mmanjoura/nomad/backend/order"
)

type Order struct{}

func (r *Order) DecodeOrder(input []byte) (*order.Order, error) {
	attr := &order.Order{}
	if err := json.Unmarshal(input, attr); err != nil {
		return nil, backend.Errorf(backend.EINTERNAL, "serializer.Order.DecodeAttriubte")
	}
	return attr, nil
}

func (r *Order) EncodeOrder(input *order.Order) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, backend.Errorf(backend.EINTERNAL, "serializer.Order.EncodeOrder")
	}
	return rawMsg, nil
}
