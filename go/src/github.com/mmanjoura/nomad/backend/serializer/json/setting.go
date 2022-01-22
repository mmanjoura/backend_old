package json

import (
	"encoding/json"

	"github.com/mmanjoura/nomad/backend"
	"github.com/mmanjoura/nomad/backend/setting"
)

type Setting struct{}

func (r *Setting) DecodeSetting(input []byte) (*setting.Setting, error) {
	attr := &setting.Setting{}
	if err := json.Unmarshal(input, attr); err != nil {
		return nil, backend.Errorf(backend.EINTERNAL, "serializer.Setting.DecodeAttriubte")
	}
	return attr, nil
}

func (r *Setting) EncodeSetting(input *setting.Setting) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, backend.Errorf(backend.EINTERNAL, "serializer.Setting.EncodeSetting")
	}
	return rawMsg, nil
}
