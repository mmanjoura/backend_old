package attribute

import (
	"context"
	"errors"

	"github.com/mmanjoura/nomad/backend"
)

var (
	ErrAttributeNotFound = errors.New("Attribute Not Found")
	ErrAttributeInvalid  = errors.New("Attribute Invalid")
)

type attributeService struct {
	attributeRepo AttributeRepository
}

func NewAttributeService(attributeRepo AttributeRepository) AttributeService {
	return &attributeService{
		attributeRepo,
	}
}

func (r *attributeService) FindOne(ctx context.Context, userId int) (*Attribute, error) {
	return r.attributeRepo.FindOne(ctx, userId)
}

func (r *attributeService) FindAll(ctx context.Context, userId int, filter backend.Filter) ([]*Attribute, int, error) {
	return r.attributeRepo.FindAll(ctx, userId, filter)
}

func (r *attributeService) Create(ctx context.Context, userId int, attribute *Attribute) error {
	return r.attributeRepo.Create(ctx, userId, attribute)
}

func (r *attributeService) Update(ctx context.Context, userId int, attr Attribute) (*Attribute, error) {
	return r.attributeRepo.Update(ctx, userId, attr)
}

func (r *attributeService) Delete(ctx context.Context, userId int) error {
	return r.attributeRepo.Delete(ctx, userId)
}

// Validate returns an error if the attribute contains invalid fields.
// This only performs basic validation.
func (u *Attribute) Validate() error {
	if u.Name == "" {
		return errors.New("Attribute Invalid")
	}
	return nil
}
