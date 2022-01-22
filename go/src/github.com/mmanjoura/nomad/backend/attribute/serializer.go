package attribute

type AttributeSerializer interface {
	DecodeAttribute(input []byte) (*Attribute, error)
	EncodeAttribute(input *Attribute) ([]byte, error)
}
