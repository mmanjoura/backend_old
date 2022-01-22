package attribute

import (
	"context"

	"github.com/mmanjoura/nomad/backend"
)

type AttributeService interface {
	Create(ctx context.Context, userId int, attribute *Attribute) error
	FindAll(ctx context.Context, userId int, filter backend.Filter) ([]*Attribute, int, error)
	FindOne(ctx context.Context, userId int) (*Attribute, error)
	Update(ctx context.Context, userId int, attr Attribute) (*Attribute, error)
	Delete(ctx context.Context, userId int) error
}
