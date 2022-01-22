package sqlite

import (
	"context"
	"strings"

	"github.com/mmanjoura/nomad/backend"
	"github.com/mmanjoura/nomad/backend/setting"
)

var _ setting.SettingService = (*SettingService)(nil)

type SettingService struct {
	db *DB
}

func NewSettingService(db *DB) *SettingService {
	return &SettingService{db: db}
}

func (s *SettingService) FindOne(ctx context.Context, id int) (*setting.Setting, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Fetch setting and their associated OAuth objects.
	setting, err := findSettingByID(ctx, tx, id)
	if err != nil {
		return nil, err
	}
	return setting, nil
}

func (s *SettingService) FindAll(ctx context.Context, settingId int, filter backend.Filter) ([]*setting.Setting, int, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, 0, err
	}
	defer tx.Rollback()
	return findSettings(ctx, tx, settingId, filter)
}

func (s *SettingService) Create(ctx context.Context, settingId int, setting *setting.Setting) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Create a new setting object and attach associated OAuth objects.
	if err := createSetting(ctx, tx, settingId, setting); err != nil {
		return err
	}
	return tx.Commit()
}

func (s *SettingService) Update(ctx context.Context, id int, c setting.Setting) (*setting.Setting, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Update setting & attach associated OAuth objects.
	setting, err := updateSetting(ctx, tx, id, c)
	if err != nil {
		return setting, err
	} else if err := tx.Commit(); err != nil {
		return setting, err
	}
	return setting, nil
}

func (s *SettingService) Delete(ctx context.Context, id int) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := deleteSetting(ctx, tx, id); err != nil {
		return err
	}
	return tx.Commit()
}

func findSettingByID(ctx context.Context, tx *Tx, settingId int) (*setting.Setting, error) {
	a, _, err := findSettings(ctx, tx, settingId, backend.Filter{ID: &settingId})
	if err != nil {
		return nil, err
	} else if len(a) == 0 {
		return nil, &backend.Error{Code: backend.ENOTFOUND, Message: "Setting not found."}
	}
	return a[0], nil
}

func findSettings(ctx context.Context, tx *Tx, settingId int, filter backend.Filter) (_ []*setting.Setting, n int, err error) {

	// Build WHERE clause.
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := &settingId; v != nil {
		where, args = append(where, "setting.id = ?"), append(args, *v)
	}

	// Execute query to fetch setting childeren values rows.
	rows, err := tx.QueryContext(ctx, GetSettings+strings.Join(where, " AND ")+`
		ORDER BY setting.id ASC
		`+FormatLimitOffset(filter.Limit, filter.Offset),
		args...,
	)
	if err != nil {
		return nil, n, err
	}
	defer rows.Close()
	// Deserialize rows into sliders objects.
	settings := make([]*setting.Setting, 0)
	var settg setting.Setting
	var settingOption setting.Option
	var seo setting.Seo
	var logo setting.Logo
	var contctDetail setting.ContactDetail
	var location setting.Location
	for rows.Next() {

		if err := rows.Scan(
			&settg.ID,
			(*NullTime)(&settg.CreatedAt),
			(*NullTime)(&settg.UpdatedAt),

			&settingOption.Currency,
			&settingOption.TaxClass,
			&settingOption.SiteTitle,
			&settingOption.SiteSubtitle,
			&settingOption.ShippingClass,
			&settingOption.MinimumOrderAmount,

			&seo.OgImage,
			&seo.OgTitle,
			&seo.MetaTags,
			&seo.MetaTitle,
			&seo.CanonicalURL,
			&seo.OgDescription,
			&seo.TwitterCardType,
			&seo.TwitterHandle,

			&logo.ID,
			&logo.Original,
			&logo.Thumbnail,

			&contctDetail.Contact,
			&contctDetail.Website,

			&location.Lat,
			&location.Lng,
			&location.State,
			&location.Country,
			&location.FormattedAddress,
		); err != nil {
			return nil, 0, err
		}
		settg.Option = settingOption
		settg.Option.Seo = seo
		settg.Option.Logo = logo
		contctDetail.Location = location

		// Get Socials
		socials, n, err := Socails(ctx, tx, filter, settingId)
		if err != nil {
			return nil, n, err
		}
		contctDetail.Socials = socials
		settg.Option.ContactDetail = contctDetail

		// Get Delivery Time
		deliveryTimes, n, err := DeliveryTimes(ctx, tx, filter, settingId)
		if err != nil {
			return nil, n, err
		}

		settg.Option.DeliveryTime = deliveryTimes
		settings = append(settings, &settg)

	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return settings, n, nil
}

func createSetting(ctx context.Context, tx *Tx, settingId int, setting *setting.Setting) error {
	// Set timestamps to the current time.
	setting.CreatedAt = tx.now
	setting.UpdatedAt = setting.CreatedAt

	//Get this from Db
	setting.ID = settingId

	// Perform basic field validation.
	if err := setting.Validate(); err != nil {
		return err
	}

	// Execute insertion query.
	result, err := tx.ExecContext(ctx, `
		INSERT INTO setting (
			Shop_id,
			slug,
			name,
			created_at,
			updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`,
		setting.Option.Currency,
		setting.Option.Currency,
		setting.ID,
		(*NullTime)(&setting.CreatedAt),
		(*NullTime)(&setting.UpdatedAt),
	)
	if err != nil {
		return FormatError(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	setting.ID = int(id)

	return nil
}

func updateSetting(ctx context.Context, tx *Tx, id int, attr setting.Setting) (*setting.Setting, error) {
	// Fetch current object state.
	setting, err := findSettingByID(ctx, tx, id)
	if err != nil {
		return setting, err
	} //else if setting.ID != setting.AttributeIDFromContext(ctx) {
	// 	return nil, nomad.Errorf(nomad.EUNAUTHORIZED, "You are not allowed to update this setting.")
	// }

	// Update fields.
	if v := attr.Option.Currency; v != "" {
		setting.Option.SiteTitle = v
	}
	if v := attr.Option.Currency; v != "" {
		setting.Option.Currency = v
	}

	// Set last updated date to current time.
	setting.UpdatedAt = tx.now

	// Perform basic field validation.
	if err := setting.Validate(); err != nil {
		return setting, err
	}

	// Execute update query.
	if _, err := tx.ExecContext(ctx, `
		UPDATE setting
		SET slug = ?,
		    name = ?,
		    updated_at = ?
		WHERE id = ?
	`,
		setting.Option.Currency,
		setting.Option.Currency,
		(*NullTime)(&setting.UpdatedAt),
		id,
	); err != nil {
		return setting, FormatError(err)
	}

	return setting, nil
}

func deleteSetting(ctx context.Context, tx *Tx, settingId int) error {
	// Verify object exists.
	if _, err := findUserByID(ctx, tx, settingId); err != nil {
		return err
	} //else if user.ID != user.UserIDFromContext(ctx) {
	// 	return nomad.Errorf(nomad.EUNAUTHORIZED, "You are not allowed to delete this setting.")
	// }

	// Remove row from database.
	if _, err := tx.ExecContext(ctx, `DELETE FROM setting WHERE id = ?`, settingId); err != nil {
		return FormatError(err)
	}
	return nil
}
