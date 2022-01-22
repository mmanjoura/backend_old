package sqlite

import (
	"context"
	"strings"

	"github.com/mmanjoura/nomad/backend"
	category "github.com/mmanjoura/nomad/backend/category"
	// "github.com/mmanjoura/nomad/backend/categoryType"
)

func (s *CategoryService) TypeFindOne(ctx context.Context, categoryTypeId int, shopId int) (*category.CategoryType, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Fetch categoryType and their associated OAuth objects.
	categoryType, err := findCategoryTypeByID(ctx, tx, categoryTypeId, shopId)
	if err != nil {
		return nil, err
	}
	return categoryType, nil
}

func (s *CategoryService) TypeFindAll(ctx context.Context, categoryTypeId int, shopId int, filter backend.Filter) ([]*category.CategoryType, int, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, 0, err
	}
	defer tx.Rollback()
	return findCategoryTypes(ctx, tx, categoryTypeId, shopId, filter)
}

func (s *CategoryService) TypeCreate(ctx context.Context, categoryTypeId int, shopId int, categoryType *category.CategoryType) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Create a new categoryType object and attach associated OAuth objects.
	if err := createCategoryType(ctx, tx, categoryTypeId, categoryType); err != nil {
		return err
	}
	return tx.Commit()
}

func (s *CategoryService) TypeUpdate(ctx context.Context, categoryTypeId int, shopId int, c category.CategoryType) (*category.CategoryType, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Update categoryType & attach associated OAuth objects.
	categoryType, err := updateCategoryType(ctx, tx, categoryTypeId, shopId, c)
	if err != nil {
		return categoryType, err
	} else if err := tx.Commit(); err != nil {
		return categoryType, err
	}
	return categoryType, nil
}

func (s *CategoryService) TypeDelete(ctx context.Context, categoryTypeId int, shopId int) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := deleteCategoryType(ctx, tx, categoryTypeId, shopId); err != nil {
		return err
	}
	return tx.Commit()
}

func findCategoryTypeByID(ctx context.Context, tx *Tx, categoryTypeId int, shopId int) (*category.CategoryType, error) {
	a, _, err := findCategoryTypes(ctx, tx, categoryTypeId, shopId, backend.Filter{ID: &categoryTypeId})
	if err != nil {
		return nil, err
	} else if len(a) == 0 {
		return nil, &backend.Error{Code: backend.ENOTFOUND, Message: "CategoryType not found."}
	}
	return a[0], nil
}

func findCategoryTypes(ctx context.Context, tx *Tx, categoryTypeId int, shopId int, filter backend.Filter) (_ []*category.CategoryType, n int, err error) {

	// Build WHERE clause.
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.ID; v != nil {
		where, args = append(where, "id = ?"), append(args, *v)
	}

	// Execute query to fetch categoryType childeren values rows.
	rows, err := tx.QueryContext(ctx, GetCategory_types+strings.Join(where, " AND ")+`
		ORDER BY id ASC
		`+FormatLimitOffset(filter.Limit, filter.Offset),
		args...,
	)
	if err != nil {
		return nil, n, err
	}
	defer rows.Close()
	// Deserialize rows into sliders objects.
	categoryTypes := make([]*category.CategoryType, 0)
	for rows.Next() {
		var cType category.CategoryType

		if err := rows.Scan(
			&cType.ID,
			&cType.Name,
			&cType.Slug,
			&cType.Icon,
			(*NullTime)(&cType.CreatedAt),
			(*NullTime)(&cType.UpdatedAt),
		); err != nil {
			return nil, 0, err
		}
		// Get setting
		setting, n, err := CategorySettings(ctx, tx, filter, cType.ID, shopId)
		if err != nil {
			return nil, n, err
		}
		cType.Settings = setting

		// Get Promotional Sliders
		sliders, n, err := CategoryTypePromotionalSliders(ctx, tx, filter, cType.ID)

		if err != nil {
			return nil, n, err
		}
		cType.PromotionalSliders = sliders

		// Get Banners
		banners, n, err := Banners(ctx, tx, filter, cType.ID)

		if err != nil {
			return nil, n, err
		}

		cType.Banners = banners

		categoryTypes = append(categoryTypes, &cType)

	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return categoryTypes, n, nil
}

func createCategoryType(ctx context.Context, tx *Tx, categoryTypeId int, categoryType *category.CategoryType) error {
	// Set timestamps to the current time.
	categoryType.CreatedAt = tx.now
	categoryType.UpdatedAt = categoryType.CreatedAt

	//Get this from Db
	categoryType.ID = categoryTypeId

	// Execute insertion query.
	result, err := tx.ExecContext(ctx, `
		INSERT INTO categoryType (
			Shop_id,
			slug,
			name,
			created_at,
			updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`,
		categoryType.Name,
		categoryType.Name,
		categoryType.ID,
		(*NullTime)(&categoryType.CreatedAt),
		(*NullTime)(&categoryType.UpdatedAt),
	)
	if err != nil {
		return FormatError(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	categoryType.ID = int(id)

	return nil
}

func updateCategoryType(ctx context.Context, tx *Tx, categoryTypeId int, shopId int, attr category.CategoryType) (*category.CategoryType, error) {
	// Fetch current object state.
	categoryType, err := findCategoryTypeByID(ctx, tx, categoryTypeId, shopId)
	if err != nil {
		return categoryType, err
	} //else if categoryType.ID != categoryType.AttributeIDFromContext(ctx) {
	// 	return nil, nomad.Errorf(nomad.EUNAUTHORIZED, "You are not allowed to update this categoryType.")
	// }

	// Update fields.
	if v := attr.Name; v != "" {
		categoryType.Name = v
	}
	if v := attr.Name; v != "" {
		categoryType.Name = v
	}

	// Set last updated date to current time.
	categoryType.UpdatedAt = tx.now

	// Execute update query.
	if _, err := tx.ExecContext(ctx, `
		UPDATE categoryType
		SET slug = ?,
		    name = ?,
		    updated_at = ?
		WHERE id = ?
	`,
		categoryType.Name,
		categoryType.Name,
		(*NullTime)(&categoryType.UpdatedAt),
		categoryTypeId,
	); err != nil {
		return categoryType, FormatError(err)
	}

	return categoryType, nil
}

func deleteCategoryType(ctx context.Context, tx *Tx, categoryTypeId int, shopId int) error {
	// Verify object exists.
	if _, err := findUserByID(ctx, tx, categoryTypeId); err != nil {
		return err
	} //else if user.ID != user.UserIDFromContext(ctx) {
	// 	return nomad.Errorf(nomad.EUNAUTHORIZED, "You are not allowed to delete this categoryType.")
	// }

	// Remove row from database.
	if _, err := tx.ExecContext(ctx, `DELETE FROM categoryType WHERE id = ?`, categoryTypeId); err != nil {
		return FormatError(err)
	}
	return nil
}
