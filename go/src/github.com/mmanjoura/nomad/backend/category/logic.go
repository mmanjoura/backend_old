package category

import (
	"context"
	"errors"

	"github.com/mmanjoura/nomad/backend"
)

var (
	ErrCategoryTypeNotFound = errors.New("Category Type Not Found")
	ErrCategoryTypeInvalid  = errors.New("Category Type Invalid")
	ErrCategoryNotFound     = errors.New("Category Not Found")
	ErrCategoryInvalid      = errors.New("Category Invalid")
)

type categoryService struct {
	categoryRepo CategoryRepository
}

func NewCategoryService(categoryRepo CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo,
	}
}

func (r *categoryService) FindOne(ctx context.Context, categoryId int, shopId int) (*Category, error) {
	return r.categoryRepo.FindOne(ctx, categoryId, shopId)
}

func (r *categoryService) FindAll(ctx context.Context, categoryId int, shopId int, filter backend.Filter) (*CategoryData, int, error) {
	return r.categoryRepo.FindAll(ctx, categoryId, shopId, filter)
}

func (r *categoryService) Create(ctx context.Context, categoryId int, shopId int, category *Category) error {
	return r.categoryRepo.Create(ctx, categoryId, shopId, category)
}

func (r *categoryService) Update(ctx context.Context, categoryId int, shopId int, ctg Category) (*Category, error) {
	return r.categoryRepo.Update(ctx, categoryId, shopId, ctg)
}

func (r *categoryService) Delete(ctx context.Context, categoryId int, shopId int) error {
	return r.categoryRepo.Delete(ctx, categoryId, shopId)
}

func (r *categoryService) TypeFindOne(ctx context.Context, categoryTypeId int, shopId int) (*CategoryType, error) {
	return r.categoryRepo.TypeFindOne(ctx, categoryTypeId, shopId)
}

func (r *categoryService) TypeFindAll(ctx context.Context, categoryTypeId int, shopId int, filter backend.Filter) ([]*CategoryType, int, error) {
	return r.categoryRepo.TypeFindAll(ctx, categoryTypeId, shopId, filter)
}

func (r *categoryService) TypeCreate(ctx context.Context, categoryTypeId int, shopId int, categoryType *CategoryType) error {
	return r.categoryRepo.TypeCreate(ctx, categoryTypeId, shopId, categoryType)
}

func (r *categoryService) TypeUpdate(ctx context.Context, categoryTypeId int, shopId int, ctg CategoryType) (*CategoryType, error) {
	return r.categoryRepo.TypeUpdate(ctx, categoryTypeId, shopId, ctg)
}

func (r *categoryService) TypeDelete(ctx context.Context, categoryId int, shopId int) error {
	return r.categoryRepo.TypeDelete(ctx, categoryId, shopId)
}

// Validate returns an error if the category contains invalid fields.
// This only performs basic validation.
func (u *Category) Validate() error {
	if u.Name == "" {
		return errors.New("Category Invalid")
	}
	return nil
}
