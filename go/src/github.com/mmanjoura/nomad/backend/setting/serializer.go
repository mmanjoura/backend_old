package setting

type SettingSerializer interface {
	DecodeSetting(input []byte) (*Setting, error)
	EncodeSetting(input *Setting) ([]byte, error)
}
