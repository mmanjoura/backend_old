package json

// import (
// 	"encoding/json"

// 	"github.com/mmanjoura/nomad/backend"
// 	"github.com/mmanjoura/nomad/backend/categoryType"
// )

// type CategoryType struct{}

// func (r *CategoryType) DecodeCategoryType(input []byte) (*categoryType.CategoryType, error) {
// 	attr := &categoryType.CategoryType{}
// 	if err := json.Unmarshal(input, attr); err != nil {
// 		return nil, backend.Errorf(backend.EINTERNAL, "serializer.CategoryType.DecodeAttriubte")
// 	}
// 	return attr, nil
// }

// func (r *CategoryType) EncodeCategoryType(input *categoryType.CategoryType) ([]byte, error) {
// 	rawMsg, err := json.Marshal(input)
// 	if err != nil {
// 		return nil, backend.Errorf(backend.EINTERNAL, "serializer.CategoryType.EncodeCategoryType")
// 	}
// 	return rawMsg, nil
// }
