package categoryType

// import (
// 	"context"
// 	"errors"

// 	"github.com/mmanjoura/nomad/backend"
// )

// var (
// 	ErrCategoryTypeNotFound = errors.New("CategoryType Not Found")
// 	ErrCategoryTypeInvalid  = errors.New("CategoryType Invalid")
// )

// type categoryTypeService struct {
// 	categoryTypeRepo CategoryTypeRepository
// }

// func NewCategoryTypeService(categoryTypeRepo CategoryTypeRepository) CategoryTypeService {
// 	return &categoryTypeService{
// 		categoryTypeRepo,
// 	}
// }

// func (r *categoryTypeService) FindOne(ctx context.Context, categoryTypeId int, shopId int) (*CategoryType, error) {
// 	return r.categoryTypeRepo.FindOne(ctx, categoryTypeId, shopId)
// }

// func (r *categoryTypeService) FindAll(ctx context.Context, categoryTypeId int, shopId int, filter backend.Filter) ([]*CategoryType, int, error) {
// 	return r.categoryTypeRepo.FindAll(ctx, categoryTypeId, shopId, filter)
// }

// func (r *categoryTypeService) Create(ctx context.Context, categoryTypeId int, shopId int, categoryType *CategoryType) error {
// 	return r.categoryTypeRepo.Create(ctx, categoryTypeId, shopId, categoryType)
// }

// func (r *categoryTypeService) Update(ctx context.Context, categoryTypeId int, shopId int, attr CategoryType) (*CategoryType, error) {
// 	return r.categoryTypeRepo.Update(ctx, categoryTypeId, shopId, attr)
// }

// func (r *categoryTypeService) Delete(ctx context.Context, categoryTypeId int, shopId int) error {
// 	return r.categoryTypeRepo.Delete(ctx, categoryTypeId, shopId)
// }

// // Validate returns an error if the categoryType contains invalid fields.
// // This only performs basic validation.
// func (u *CategoryType) Validate() error {
// 	if u.Name == "" {
// 		return errors.New("CategoryType Invalid")
// 	}
// 	return nil
// }
