package sqlite

import (
	"context"
	"strings"

	"github.com/mmanjoura/nomad/backend"
	"github.com/mmanjoura/nomad/backend/shipping"
)

var _ shipping.ShippingService = (*ShippingService)(nil)

type ShippingService struct {
	db *DB
}

func NewShippingService(db *DB) *ShippingService {
	return &ShippingService{db: db}
}

func (s *ShippingService) FindOne(ctx context.Context, id int) (*shipping.Shipping, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Fetch shipping and their associated OAuth objects.
	shipping, err := findShippingByID(ctx, tx, id)
	if err != nil {
		return nil, err
	}
	return shipping, nil
}

func (s *ShippingService) FindAll(ctx context.Context, shippingId int, filter backend.Filter) ([]*shipping.Shipping, int, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, 0, err
	}
	defer tx.Rollback()
	return findShippings(ctx, tx, shippingId, filter)
}

func (s *ShippingService) Create(ctx context.Context, shippingId int, shipping *shipping.Shipping) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Create a new shipping object and attach associated OAuth objects.
	if err := createShipping(ctx, tx, shippingId, shipping); err != nil {
		return err
	}
	return tx.Commit()
}

func (s *ShippingService) Update(ctx context.Context, id int, c shipping.Shipping) (*shipping.Shipping, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Update shipping & attach associated OAuth objects.
	shipping, err := updateShipping(ctx, tx, id, c)
	if err != nil {
		return shipping, err
	} else if err := tx.Commit(); err != nil {
		return shipping, err
	}
	return shipping, nil
}

func (s *ShippingService) Delete(ctx context.Context, id int) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := deleteShipping(ctx, tx, id); err != nil {
		return err
	}
	return tx.Commit()
}

func findShippingByID(ctx context.Context, tx *Tx, shippingId int) (*shipping.Shipping, error) {
	a, _, err := findShippings(ctx, tx, shippingId, backend.Filter{ID: &shippingId})
	if err != nil {
		return nil, err
	} else if len(a) == 0 {
		return nil, &backend.Error{Code: backend.ENOTFOUND, Message: "Shipping not found."}
	}
	return a[0], nil
}

func findShippings(ctx context.Context, tx *Tx, shippingId int, filter backend.Filter) (_ []*shipping.Shipping, n int, err error) {

	// Build WHERE clause.
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.ID; v != nil {
		where, args = append(where, "id = ?"), append(args, *v)
	}

	// Execute query to fetch shipping childeren values rows.
	rows, err := tx.QueryContext(ctx, GetShippings+strings.Join(where, " AND ")+`
		ORDER BY id ASC
		`+FormatLimitOffset(filter.Limit, filter.Offset),
		args...,
	)
	if err != nil {
		return nil, n, err
	}
	defer rows.Close()

	// Deserialize rows into sliders objects.
	shippings := make([]*shipping.Shipping, 0)
	for rows.Next() {
		var ship shipping.Shipping
		if err := rows.Scan(
			&ship.ID,
			&ship.Name,
			&ship.Amount,
			&ship.IsGlobal,
			&ship.Type,
			(*NullTime)(&ship.CreatedAt),
			(*NullTime)(&ship.UpdatedAt),
		); err != nil {
			return nil, 0, err
		}

		shippings = append(shippings, &ship)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return shippings, n, nil
}

func createShipping(ctx context.Context, tx *Tx, shippingId int, shipping *shipping.Shipping) error {
	// Set timestamps to the current time.
	shipping.CreatedAt = tx.now
	shipping.UpdatedAt = shipping.CreatedAt

	//Get this from Db
	shipping.ID = shippingId

	// Perform basic field validation.
	if err := shipping.Validate(); err != nil {
		return err
	}

	// Execute insertion query.
	result, err := tx.ExecContext(ctx, `
		INSERT INTO shipping (
			Shop_id,
			slug,
			name,
			created_at,
			updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`,
		shipping.Name,
		shipping.Type,
		shipping.ID,
		(*NullTime)(&shipping.CreatedAt),
		(*NullTime)(&shipping.UpdatedAt),
	)
	if err != nil {
		return FormatError(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	shipping.ID = int(id)

	return nil
}

func updateShipping(ctx context.Context, tx *Tx, id int, attr shipping.Shipping) (*shipping.Shipping, error) {
	// Fetch current object state.
	shipping, err := findShippingByID(ctx, tx, id)
	if err != nil {
		return shipping, err
	} //else if shipping.ID != shipping.AttributeIDFromContext(ctx) {
	// 	return nil, nomad.Errorf(nomad.EUNAUTHORIZED, "You are not allowed to update this shipping.")
	// }

	// Update fields.
	if v := attr.Name; v != "" {
		shipping.Name = v
	}
	if v := attr.Type; v != "" {
		shipping.Type = v
	}

	// Set last updated date to current time.
	shipping.UpdatedAt = tx.now

	// Perform basic field validation.
	if err := shipping.Validate(); err != nil {
		return shipping, err
	}

	// Execute update query.
	if _, err := tx.ExecContext(ctx, `
		UPDATE shipping
		SET slug = ?,
		    name = ?,
		    updated_at = ?
		WHERE id = ?
	`,
		shipping.Type,
		shipping.Name,
		(*NullTime)(&shipping.UpdatedAt),
		id,
	); err != nil {
		return shipping, FormatError(err)
	}

	return shipping, nil
}

func deleteShipping(ctx context.Context, tx *Tx, shippingId int) error {
	// Verify object exists.
	if _, err := findUserByID(ctx, tx, shippingId); err != nil {
		return err
	} //else if user.ID != user.UserIDFromContext(ctx) {
	// 	return nomad.Errorf(nomad.EUNAUTHORIZED, "You are not allowed to delete this shipping.")
	// }

	// Remove row from database.
	if _, err := tx.ExecContext(ctx, `DELETE FROM shipping WHERE id = ?`, shippingId); err != nil {
		return FormatError(err)
	}
	return nil
}
