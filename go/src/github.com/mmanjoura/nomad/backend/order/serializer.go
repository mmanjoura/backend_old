package order

type OrderSerializer interface {
	DecodeOrder(input []byte) (*Order, error)
	EncodeOrder(input *Order) ([]byte, error)
}
