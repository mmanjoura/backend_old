package sqlite

import (
	"context"
	"strings"

	"github.com/mmanjoura/nomad/backend"
	"github.com/mmanjoura/nomad/backend/coupon"
)

var _ coupon.CouponService = (*CouponService)(nil)

type CouponService struct {
	db *DB
}

func NewCouponService(db *DB) *CouponService {
	return &CouponService{db: db}
}

func (s *CouponService) FindOne(ctx context.Context, id int) (*coupon.Coupon, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Fetch coupon and their associated OAuth objects.
	coupon, err := findCouponByID(ctx, tx, id)
	if err != nil {
		return nil, err
	}
	return coupon, nil
}

func (s *CouponService) FindAll(ctx context.Context, couponId int, filter backend.Filter) ([]*coupon.Coupon, int, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, 0, err
	}
	defer tx.Rollback()
	return findCoupons(ctx, tx, couponId, filter)
}

func (s *CouponService) Create(ctx context.Context, couponId int, coupon *coupon.Coupon) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Create a new coupon object and attach associated OAuth objects.
	if err := createCoupon(ctx, tx, couponId, coupon); err != nil {
		return err
	}
	return tx.Commit()
}

func (s *CouponService) Update(ctx context.Context, id int, c coupon.Coupon) (*coupon.Coupon, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Update coupon & attach associated OAuth objects.
	coupon, err := updateCoupon(ctx, tx, id, c)
	if err != nil {
		return coupon, err
	} else if err := tx.Commit(); err != nil {
		return coupon, err
	}
	return coupon, nil
}

func (s *CouponService) Delete(ctx context.Context, id int) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := deleteCoupon(ctx, tx, id); err != nil {
		return err
	}
	return tx.Commit()
}

func findCouponByID(ctx context.Context, tx *Tx, couponId int) (*coupon.Coupon, error) {
	a, _, err := findCoupons(ctx, tx, couponId, backend.Filter{ID: &couponId})
	if err != nil {
		return nil, err
	} else if len(a) == 0 {
		return nil, &backend.Error{Code: backend.ENOTFOUND, Message: "Coupon not found."}
	}
	return a[0], nil
}

func findCoupons(ctx context.Context, tx *Tx, couponId int, filter backend.Filter) (_ []*coupon.Coupon, n int, err error) {

	// Build WHERE clause.
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.ID; v != nil {
		where, args = append(where, "id = ?"), append(args, *v)
	}

	// Execute query to fetch coupon childeren values rows.
	rows, err := tx.QueryContext(ctx, GetImages+strings.Join(where, " AND ")+`
		ORDER BY id ASC
		`+FormatLimitOffset(filter.Limit, filter.Offset),
		args...,
	)
	if err != nil {
		return nil, n, err
	}
	defer rows.Close()

	// Deserialize rows into sliders objects.
	cImages := make([]coupon.Image, 0)
	for rows.Next() {
		var image coupon.Image
		if err := rows.Scan(
			&image.Original,
			&image.Thumbnail,
		); err != nil {
			return nil, 0, err
		}

		cImages = append(cImages, image)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	// Execute query to fetch coupon childeren values rows.
	rows, err = tx.QueryContext(ctx, GetCoupons+strings.Join(where, " AND ")+`
		ORDER BY id ASC
		`+FormatLimitOffset(filter.Limit, filter.Offset),
		args...,
	)
	if err != nil {
		return nil, n, err
	}
	defer rows.Close()
	// Deserialize rows into sliders objects.
	coupons := make([]*coupon.Coupon, 0)
	for rows.Next() {
		var c coupon.Coupon
		if err := rows.Scan(
			&c.ID,
			&c.Code,
			&c.Description,
			&c.Type,
			&c.Amount,
			(*NullTime)(&c.ActiveFrom),
			(*NullTime)(&c.ExpireAt),
			(*NullTime)(&c.CreatedAt),
			(*NullTime)(&c.UpdatedAt),
			(*NullTime)(&c.DeletedAt),
			&c.IsValid,
		); err != nil {
			return nil, 0, err
		}
		//fix
		c.Image = cImages[0]
		coupons = append(coupons, &c)

	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return coupons, n, nil
}

func createCoupon(ctx context.Context, tx *Tx, couponId int, coupon *coupon.Coupon) error {
	// Set timestamps to the current time.
	coupon.CreatedAt = tx.now
	coupon.UpdatedAt = coupon.CreatedAt

	//Get this from Db
	coupon.ID = couponId

	// Perform basic field validation.
	if err := coupon.Validate(); err != nil {
		return err
	}

	// Execute insertion query.
	result, err := tx.ExecContext(ctx, `
		INSERT INTO coupon (
			Shop_id,
			slug,
			name,
			created_at,
			updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`,
		coupon.Code,
		coupon.Type,
		coupon.ID,
		(*NullTime)(&coupon.CreatedAt),
		(*NullTime)(&coupon.UpdatedAt),
	)
	if err != nil {
		return FormatError(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	coupon.ID = int(id)

	return nil
}

func updateCoupon(ctx context.Context, tx *Tx, id int, attr coupon.Coupon) (*coupon.Coupon, error) {
	// Fetch current object state.
	coupon, err := findCouponByID(ctx, tx, id)
	if err != nil {
		return coupon, err
	} //else if coupon.ID != coupon.AttributeIDFromContext(ctx) {
	// 	return nil, nomad.Errorf(nomad.EUNAUTHORIZED, "You are not allowed to update this coupon.")
	// }

	// Update fields.
	if v := attr.Code; v != "" {
		coupon.Code = v
	}
	if v := attr.Type; v != "" {
		coupon.Type = v
	}

	// Set last updated date to current time.
	coupon.UpdatedAt = tx.now

	// Perform basic field validation.
	if err := coupon.Validate(); err != nil {
		return coupon, err
	}

	// Execute update query.
	if _, err := tx.ExecContext(ctx, `
		UPDATE coupon
		SET slug = ?,
		    name = ?,
		    updated_at = ?
		WHERE id = ?
	`,
		coupon.Type,
		coupon.Code,
		(*NullTime)(&coupon.UpdatedAt),
		id,
	); err != nil {
		return coupon, FormatError(err)
	}

	return coupon, nil
}

func deleteCoupon(ctx context.Context, tx *Tx, couponId int) error {
	// Verify object exists.
	if _, err := findUserByID(ctx, tx, couponId); err != nil {
		return err
	} //else if user.ID != user.UserIDFromContext(ctx) {
	// 	return nomad.Errorf(nomad.EUNAUTHORIZED, "You are not allowed to delete this coupon.")
	// }

	// Remove row from database.
	if _, err := tx.ExecContext(ctx, `DELETE FROM coupon WHERE id = ?`, couponId); err != nil {
		return FormatError(err)
	}
	return nil
}
