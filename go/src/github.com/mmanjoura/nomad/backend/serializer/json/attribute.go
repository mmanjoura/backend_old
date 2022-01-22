package json

import (
	"encoding/json"

	"github.com/mmanjoura/nomad/backend"
	"github.com/mmanjoura/nomad/backend/attribute"
)

type Attribute struct{}

func (r *Attribute) DecodeAttribute(input []byte) (*attribute.Attribute, error) {
	attr := &attribute.Attribute{}
	if err := json.Unmarshal(input, attr); err != nil {
		return nil, backend.Errorf(backend.EINTERNAL, "serializer.Attribute.DecodeAttriubte")
	}
	return attr, nil
}

func (r *Attribute) EncodeAttribute(input *attribute.Attribute) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, backend.Errorf(backend.EINTERNAL, "serializer.Attribute.EncodeAttribute")
	}
	return rawMsg, nil
}
