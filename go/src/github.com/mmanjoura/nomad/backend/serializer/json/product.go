package json

import (
	"encoding/json"

	"github.com/mmanjoura/nomad/backend"
	"github.com/mmanjoura/nomad/backend/product"
)

type Product struct{}

func (r *Product) DecodeProduct(input []byte) (*product.Product, error) {
	attr := &product.Product{}
	if err := json.Unmarshal(input, attr); err != nil {
		return nil, backend.Errorf(backend.EINTERNAL, "serializer.Product.DecodeAttriubte")
	}
	return attr, nil
}

func (r *Product) EncodeProduct(input *product.Product) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, backend.Errorf(backend.EINTERNAL, "serializer.Product.EncodeProduct")
	}
	return rawMsg, nil
}
