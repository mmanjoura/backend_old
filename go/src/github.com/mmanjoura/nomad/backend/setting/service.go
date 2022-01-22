package setting

import (
	"context"

	"github.com/mmanjoura/nomad/backend"
)

type SettingService interface {
	Create(ctx context.Context, settingId int, setting *Setting) error
	FindAll(ctx context.Context, settingId int, filter backend.Filter) ([]*Setting, int, error)
	FindOne(ctx context.Context, settingId int) (*Setting, error)
	Update(ctx context.Context, settingId int, attr Setting) (*Setting, error)
	Delete(ctx context.Context, settingId int) error
}
