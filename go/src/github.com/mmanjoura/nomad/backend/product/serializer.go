package product

type ProductSerializer interface {
	DecodeProduct(input []byte) (*Product, error)
	EncodeProduct(input *Product) ([]byte, error)
}
