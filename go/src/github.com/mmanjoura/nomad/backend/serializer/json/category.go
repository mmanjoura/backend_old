package json

import (
	"encoding/json"

	"github.com/mmanjoura/nomad/backend"
	"github.com/mmanjoura/nomad/backend/category"
)

type Category struct{}

func (r *Category) DecodeCategory(input []byte) (*category.Category, error) {
	ctg := &category.Category{}
	if err := json.Unmarshal(input, ctg); err != nil {
		return nil, backend.Errorf(backend.EINTERNAL, "serializer.Category.DecodeCategory")
	}
	return ctg, nil
}

func (r *Category) EncodeCategory(input *category.Category) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, backend.Errorf(backend.EINTERNAL, "serializer.Category.EncodeCategory")
	}
	return rawMsg, nil
}

func (r *Category) DecodeCategoryType(input []byte) (*category.CategoryType, error) {
	ctgType := &category.CategoryType{}
	if err := json.Unmarshal(input, ctgType); err != nil {
		return nil, backend.Errorf(backend.EINTERNAL, "serializer.Category.DecodeAttriubte")
	}
	return ctgType, nil
}

func (r *Category) EncodeCategoryType(input *category.CategoryType) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, backend.Errorf(backend.EINTERNAL, "serializer.Category.EncodeCategoryType")
	}
	return rawMsg, nil
}
