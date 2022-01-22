package sqlite

import (
	"context"
	"strings"

	"github.com/mmanjoura/nomad/backend"
	attribute "github.com/mmanjoura/nomad/backend/attribute"
	"github.com/mmanjoura/nomad/backend/setting"
)

// Ensure service implements interface.
var _ attribute.AttributeService = (*AttributeService)(nil)

// AttributeService represents a service for managing attributes.
type AttributeService struct {
	db *DB
}

// NewAttributeService returns a new instance of AttributeService.
func NewAttributeService(db *DB) *AttributeService {
	return &AttributeService{db: db}
}

// FindAttributeByID retrieves a attribute by ID along with their associated auth objects.
// Returns ENOTFOUND if attribute does not exist.
func (s *AttributeService) FindOne(ctx context.Context, id int) (*attribute.Attribute, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Fetch attribute and their associated OAuth objects.
	attribute, err := findAttributeByID(ctx, tx, id)
	if err != nil {
		return nil, err
	}
	return attribute, nil
}

// FindAttributes retrieves a list of attributes by filter. Also returns total count of
// matching attributes which may differ from returned results if filter.Limit is specified.
func (s *AttributeService) FindAll(ctx context.Context, userId int, filter backend.Filter) ([]*attribute.Attribute, int, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, 0, err
	}
	defer tx.Rollback()
	return findAttributes(ctx, tx, userId, filter)
}

// CreateAttribute creates a new attribute. This is only used for testing since attributes are
// typically created during the OAuth creation process in AuthService.CreateAuth().
func (s *AttributeService) Create(ctx context.Context, userId int, attribute *attribute.Attribute) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Create a new attribute object and attach associated OAuth objects.
	if err := createAttribute(ctx, tx, userId, attribute); err != nil {
		return err
	}
	return tx.Commit()
}

// UpdateAttribute updates a attribute object. Returns EUNAUTHORIZED if current attribute is
// not the attribute that is being updated. Returns ENOTFOUND if attribute does not exist.
func (s *AttributeService) Update(ctx context.Context, id int, upd attribute.Attribute) (*attribute.Attribute, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Update attribute & attach associated OAuth objects.
	attribute, err := updateAttribute(ctx, tx, id, upd)
	if err != nil {
		return attribute, err
	} else if err := tx.Commit(); err != nil {
		return attribute, err
	}
	return attribute, nil
}

// DeleteAttribute permanently deletes a attribute and all owned shops.
// Returns EUNAUTHORIZED if current attribute is not the attribute being deleted.
// Returns ENOTFOUND if attribute does not exist.
func (s *AttributeService) Delete(ctx context.Context, id int) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := deleteAttribute(ctx, tx, id); err != nil {
		return err
	}
	return tx.Commit()
}

// findAttributeByID is a helper function to fetch a attribute by ID.
// Returns ENOTFOUND if attribute does not exist.
func findAttributeByID(ctx context.Context, tx *Tx, userId int) (*attribute.Attribute, error) {
	a, _, err := findAttributes(ctx, tx, userId, backend.Filter{ID: &userId})
	if err != nil {
		return nil, err
	} else if len(a) == 0 {
		return nil, &backend.Error{Code: backend.ENOTFOUND, Message: "Attribute not found."}
	}
	return a[0], nil
}

// findAttributes returns a list of attributes matching a filter. Also returns a count of
// total matching attributes which may differ if filter.Limit is set.
func findAttributes(ctx context.Context, tx *Tx, userId int, filter backend.Filter) (_ []*attribute.Attribute, n int, err error) {
	// Build WHERE clause.
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.ID; v != nil {
		where, args = append(where, "id = ?"), append(args, *v)
	}

	// Execute query to fetch attribute values rows.
	rows, err := tx.QueryContext(ctx, GetAttributeValues+strings.Join(where, " AND ")+`
		ORDER BY attribute.id ASC
		`+FormatLimitOffset(filter.Limit, filter.Offset),
		args...,
	)
	if err != nil {
		return nil, n, err
	}
	defer rows.Close()

	// Deserialize rows into Attribute objects.
	attributeValues := make([]attribute.Value, 0)
	for rows.Next() {
		var attr attribute.Value
		if err := rows.Scan(
			&attr.ID,
			&attr.AttributeID,
			&attr.Value,
			&attr.Meta,
			(*NullTime)(&attr.CreatedAt),
			(*NullTime)(&attr.UpdatedAt),
		); err != nil {
			return nil, 0, err
		}

		attributeValues = append(attributeValues, attr)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	// Execute query to fetch attribute values rows.
	rows, err = tx.QueryContext(ctx, GetSocials+strings.Join(where, " AND ")+`
		ORDER BY id ASC
		`+FormatLimitOffset(filter.Limit, filter.Offset),
		args...,
	)
	if err != nil {
		return nil, n, err
	}
	defer rows.Close()

	// Deserialize rows into Attribute objects.
	shopSocials := make([]attribute.Social, 0)
	for rows.Next() {
		var social attribute.Social
		if err := rows.Scan(
			&social.URL,
			&social.Icon,
		); err != nil {
			return nil, 0, err
		}

		shopSocials = append(shopSocials, social)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	// Execute query to fetch attribute rows.
	rows, err = tx.QueryContext(ctx, GetAttributes+strings.Join(where, " AND ")+`
		ORDER BY attribute.id ASC
		`+FormatLimitOffset(filter.Limit, filter.Offset),
		args...,
	)
	if err != nil {
		return nil, n, err
	}
	defer rows.Close()

	// Deserialize rows into Attribute objects.
	attributes := make([]*attribute.Attribute, 0)
	for rows.Next() {
		var attr attribute.Attribute
		// var value attribute.Value
		var shop attribute.Shop
		var coverImage attribute.CoverImage
		var logo attribute.Logo
		var address attribute.Address
		var contactDetail setting.ContactDetail
		var location attribute.Location

		if err := rows.Scan(
			&attr.ID,
			&attr.Slug,
			&attr.Name,
			&attr.ShopID,
			(*NullTime)(&attr.CreatedAt),
			(*NullTime)(&attr.UpdatedAt),

			&shop.ID,
			&shop.OwnerID,
			&shop.Name,
			&shop.Slug,
			&shop.Description,
			&shop.IsActive,
			(*NullTime)(&shop.CreatedAt),
			(*NullTime)(&shop.UpdatedAt),

			&coverImage.ID,
			&coverImage.Original,
			&coverImage.Thumbnail,

			&logo.ID,
			&logo.Original,
			&logo.Thumbnail,

			&address.Zip,
			&address.City,
			&address.State,
			&address.Country,
			&address.StreetAddress,

			&contactDetail.Contact,
			&contactDetail.Website,

			&location.Lat,
			&location.Lng,
			&location.City,
			&location.State,
			&location.Country,
			&location.FormattedAddress,
		); err != nil {
			return nil, 0, err
		}
		// contactDetail.Location = location
		// contactDetail.Socials = shopSocials
		// c= contactDetail
		shop.Address = address
		shop.CoverImage = coverImage
		shop.Logo = logo
		attr.Shop = shop
		attr.Value = attributeValues

		attributes = append(attributes, &attr)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return attributes, n, nil
}

// createAttribute creates a new attribute. Sets the new database ID to attribute.ID and sets
// the timestamps to the current time.
func createAttribute(ctx context.Context, tx *Tx, userId int, attribute *attribute.Attribute) error {
	// Set timestamps to the current time.
	attribute.CreatedAt = tx.now
	attribute.UpdatedAt = attribute.CreatedAt

	//Get this from Db
	attribute.ShopID = userId

	// Perform basic field validation.
	if err := attribute.Validate(); err != nil {
		return err
	}

	// Execute insertion query.
	result, err := tx.ExecContext(ctx, `
		INSERT INTO attribute (
			Shop_id,
			slug,
			name,
			created_at,
			updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`,
		attribute.Slug,
		attribute.Name,
		attribute.ShopID,
		(*NullTime)(&attribute.CreatedAt),
		(*NullTime)(&attribute.UpdatedAt),
	)
	if err != nil {
		return FormatError(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	attribute.ID = int(id)

	return nil
}

// updateAttribute updates fields on a attribute object. Returns EUNAUTHORIZED if current
// attribute is not the attribute being updated.
func updateAttribute(ctx context.Context, tx *Tx, id int, attr attribute.Attribute) (*attribute.Attribute, error) {
	// Fetch current object state.
	attribute, err := findAttributeByID(ctx, tx, id)
	if err != nil {
		return attribute, err
	} //else if attribute.ID != attribute.AttributeIDFromContext(ctx) {
	// 	return nil, nomad.Errorf(nomad.EUNAUTHORIZED, "You are not allowed to update this attribute.")
	// }

	// Update fields.
	if v := attr.Name; v != "" {
		attribute.Name = v
	}
	if v := attr.Slug; v != "" {
		attribute.Slug = v
	}

	// Set last updated date to current time.
	attribute.UpdatedAt = tx.now

	// Perform basic field validation.
	if err := attribute.Validate(); err != nil {
		return attribute, err
	}

	// Execute update query.
	if _, err := tx.ExecContext(ctx, `
		UPDATE attribute
		SET slug = ?,
		    name = ?,
		    updated_at = ?
		WHERE id = ?
	`,
		attribute.Slug,
		attribute.Name,
		(*NullTime)(&attribute.UpdatedAt),
		id,
	); err != nil {
		return attribute, FormatError(err)
	}

	return attribute, nil
}

// deleteAttribute permanently removes a attribute by ID. Returns EUNAUTHORIZED if current
// attribute is not the one being deleted.
func deleteAttribute(ctx context.Context, tx *Tx, userId int) error {
	// Verify object exists.
	if _, err := findUserByID(ctx, tx, userId); err != nil {
		return err
	} //else if user.ID != user.UserIDFromContext(ctx) {
	// 	return nomad.Errorf(nomad.EUNAUTHORIZED, "You are not allowed to delete this attribute.")
	// }

	// Remove row from database.
	if _, err := tx.ExecContext(ctx, `DELETE FROM attribute WHERE id = ?`, userId); err != nil {
		return FormatError(err)
	}
	return nil
}
