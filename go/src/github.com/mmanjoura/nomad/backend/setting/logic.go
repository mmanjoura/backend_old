package setting

import (
	"context"
	"errors"

	"github.com/mmanjoura/nomad/backend"
)

var (
	ErrSettingNotFound = errors.New("Setting Not Found")
	ErrSettingInvalid  = errors.New("Setting Invalid")
)

type settingService struct {
	settingRepo SettingRepository
}

func NewSettingService(settingRepo SettingRepository) SettingService {
	return &settingService{
		settingRepo,
	}
}

func (r *settingService) FindOne(ctx context.Context, settingId int) (*Setting, error) {
	return r.settingRepo.FindOne(ctx, settingId)
}

func (r *settingService) FindAll(ctx context.Context, settingId int, filter backend.Filter) ([]*Setting, int, error) {
	return r.settingRepo.FindAll(ctx, settingId, filter)
}

func (r *settingService) Create(ctx context.Context, settingId int, setting *Setting) error {
	return r.settingRepo.Create(ctx, settingId, setting)
}

func (r *settingService) Update(ctx context.Context, settingId int, attr Setting) (*Setting, error) {
	return r.settingRepo.Update(ctx, settingId, attr)
}

func (r *settingService) Delete(ctx context.Context, settingId int) error {
	return r.settingRepo.Delete(ctx, settingId)
}

// Validate returns an error if the setting contains invalid fields.
// This only performs basic validation.
func (u *Setting) Validate() error {
	if u.Option.Currency == "" {
		return errors.New("Setting Invalid")
	}
	return nil
}
