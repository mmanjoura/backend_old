package category

type CategorySerializer interface {
	DecodeCategory(input []byte) (*Category, error)
	EncodeCategory(input *Category) ([]byte, error)
	DecodeCategoryType(input []byte) (*CategoryType, error)
	EncodeCategoryType(input *CategoryType) ([]byte, error)
}
