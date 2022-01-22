package category

import (
	"context"

	"github.com/mmanjoura/nomad/backend"
)

type CategoryService interface {
	Create(ctx context.Context, categoryId int, shopId int, category *Category) error
	FindAll(ctx context.Context, categoryId int, shopId int, filter backend.Filter) (*CategoryData, int, error)
	FindOne(ctx context.Context, categoryId int, shopId int) (*Category, error)
	Update(ctx context.Context, categoryId int, shopId int, attr Category) (*Category, error)
	Delete(ctx context.Context, categoryId int, shopId int) error

	TypeCreate(ctx context.Context, categoryTypeId int, shopId int, categoryType *CategoryType) error
	TypeFindAll(ctx context.Context, categoryTypeId int, shopId int, filter backend.Filter) ([]*CategoryType, int, error)
	TypeFindOne(ctx context.Context, categoryTypeId int, shopId int) (*CategoryType, error)
	TypeUpdate(ctx context.Context, categoryTypeId int, shopId int, attr CategoryType) (*CategoryType, error)
	TypeDelete(ctx context.Context, categoryTypeId int, shopId int) error
}
