package shipping

type ShippingSerializer interface {
	DecodeShipping(input []byte) (*Shipping, error)
	EncodeShipping(input *Shipping) ([]byte, error)
}
