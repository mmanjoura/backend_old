package sqlite

import (
	"context"
	"strings"

	"github.com/mmanjoura/nomad/backend"
	category "github.com/mmanjoura/nomad/backend/category"
)

var _ category.CategoryService = (*CategoryService)(nil)

type CategoryService struct {
	db *DB
}

func NewCategoryService(db *DB) *CategoryService {
	return &CategoryService{db: db}
}

func (s *CategoryService) FindOne(ctx context.Context, categoryId int, shopId int) (*category.Category, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Fetch category and their associated OAuth objects.
	category, err := findCategoryByID(ctx, tx, categoryId, shopId)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *CategoryService) FindAll(ctx context.Context, categoryId int, shopId int, filter backend.Filter) (*category.CategoryData, int, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, 0, err
	}
	defer tx.Rollback()
	return findCategories(ctx, tx, categoryId, shopId, filter)
}

func (s *CategoryService) Create(ctx context.Context, categoryId int, shopId int, category *category.Category) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Create a new category object and attach associated OAuth objects.
	if err := createCategory(ctx, tx, categoryId, shopId, category); err != nil {
		return err
	}
	return tx.Commit()
}

func (s *CategoryService) Update(ctx context.Context, categoryId int, shopId int, c category.Category) (*category.Category, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Update category & attach associated OAuth objects.
	category, err := updateCategory(ctx, tx, categoryId, shopId, c)
	if err != nil {
		return category, err
	} else if err := tx.Commit(); err != nil {
		return category, err
	}
	return category, nil
}

func (s *CategoryService) Delete(ctx context.Context, categoryId int, shopId int) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := deleteCategory(ctx, tx, categoryId, shopId); err != nil {
		return err
	}
	return tx.Commit()
}

func findCategoryByID(ctx context.Context, tx *Tx, categoryId int, shopId int) (*category.Category, error) {
	a, _, err := findCategories(ctx, tx, categoryId, shopId, backend.Filter{ID: &categoryId})
	if err != nil {
		return nil, err
	} else if len(a.Data) == 0 {
		return nil, &backend.Error{Code: backend.ENOTFOUND, Message: "Category not found."}
	}
	return &a.Data[0], nil
}

func findCategories(ctx context.Context, tx *Tx, categoryId int, shopId int, filter backend.Filter) (_ *category.CategoryData, n int, err error) {

	// Build WHERE clause.
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.ID; v != nil {
		where, args = append(where, "id = ?"), append(args, *v)
	}

	if v := &categoryId; v != nil {
		where, args = append(where, "category.id = ?"), append(args, *v)
	}

	//We only calling parent Category
	//parent_id := 0
	// if v := &parent_id; v != nil {
	// 	where, args = append(where, "parent_id = ?"), append(args, *v)
	// }

	// Execute query to fetch category values rows.
	rows, err := tx.QueryContext(ctx, GetCategories+strings.Join(where, " AND ")+`
		ORDER BY category.id ASC
		`+FormatLimitOffset(filter.Limit, filter.Offset),
		args...,
	)
	if err != nil {
		return &category.CategoryData{}, n, err
	}
	defer rows.Close()

	// Deserialize rows into Category objects.
	categories := make([]category.Category, 0)
	for rows.Next() {
		var catgr category.Category
		if err := rows.Scan(
			&catgr.ID,
			&catgr.Name,
			&catgr.Slug,
			&catgr.Icon,
			&catgr.TypeID,
			(*NullTime)(&catgr.CreatedAt),
			(*NullTime)(&catgr.UpdatedAt),
			(*NullTime)(&catgr.UpdatedAt),
			&catgr.ParentID,
		); err != nil {
			return &category.CategoryData{}, 0, err
		}

		// get category images
		images, n, err := Images(ctx, tx, filter, catgr.ID)
		if err != nil {
			return &category.CategoryData{}, n, err
		}
		if len(images) == 0 {
			catgr.Image = make([]category.Image, 0)

		} else {
			catgr.Image = images
		}

		// Get Category Type
		cType, n, err := CategoryType(ctx, tx, filter, categoryId, shopId)
		if err != nil {
			return &category.CategoryData{}, n, err
		}
		catgr.CategoryType = cType

		// Get category children
		catChildren, n, err := Childrens(ctx, tx, filter, catgr.ID)
		if err != nil {
			return &category.CategoryData{}, n, err
		}
		catgr.Children = catChildren
		categories = append(categories, catgr)

	}

	if err := rows.Err(); err != nil {
		return &category.CategoryData{}, 0, err
	}

	categorytData := category.CategoryData{
		Data:         categories,
		Total:        30,
		CurrentPage:  1,
		Count:        0,
		LastPage:     1,
		FirstItem:    0,
		LastItem:     29,
		PerPage:      "30",
		FirstPageURL: "http://localhost:3000/api/products?search=type.slug:grocery&limit=30&page=1",
		LastPageURL:  "http://localhost:3000/api/products?search=type.slug:grocery&limit=30&page=9",
		NextPageURL:  "http://localhost:3000/api/products?search=type.slug:grocery&limit=30&page=2",
		PrevPageURL:  "http://localhost:3000/api/products?search=type.slug:grocery&limit=30&page=1",
	}

	return &categorytData, n, nil
}

func createCategory(ctx context.Context, tx *Tx, categoryId int, shopId int, category *category.Category) error {
	// Set timestamps to the current time.
	category.CreatedAt = tx.now
	category.UpdatedAt = category.CreatedAt

	//Get this from Db
	category.ID = categoryId

	// Perform basic field validation.
	if err := category.Validate(); err != nil {
		return err
	}

	// Execute insertion query.
	result, err := tx.ExecContext(ctx, `
		INSERT INTO category (
			Shop_id,
			slug,
			name,
			created_at,
			updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`,
		category.Slug,
		category.Name,
		category.ID,
		(*NullTime)(&category.CreatedAt),
		(*NullTime)(&category.UpdatedAt),
	)
	if err != nil {
		return FormatError(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	category.ID = int(id)

	return nil
}

func updateCategory(ctx context.Context, tx *Tx, categoryId int, shopId int, attr category.Category) (*category.Category, error) {
	// Fetch current object state.
	category, err := findCategoryByID(ctx, tx, categoryId, shopId)
	if err != nil {
		return category, err
	} //else if category.ID != category.AttributeIDFromContext(ctx) {
	// 	return nil, nomad.Errorf(nomad.EUNAUTHORIZED, "You are not allowed to update this category.")
	// }

	// Update fields.
	if v := attr.Name; v != "" {
		category.Name = v
	}
	if v := attr.Slug; v != "" {
		category.Slug = v
	}

	// Set last updated date to current time.
	category.UpdatedAt = tx.now

	// Perform basic field validation.
	if err := category.Validate(); err != nil {
		return category, err
	}

	// Execute update query.
	if _, err := tx.ExecContext(ctx, `
		UPDATE category
		SET slug = ?,
		    name = ?,
		    updated_at = ?
		WHERE id = ?
	`,
		category.Slug,
		category.Name,
		(*NullTime)(&category.UpdatedAt),
		categoryId,
	); err != nil {
		return category, FormatError(err)
	}

	return category, nil
}

func deleteCategory(ctx context.Context, tx *Tx, categoryId int, shopId int) error {
	// Verify object exists.
	if _, err := findUserByID(ctx, tx, categoryId); err != nil {
		return err
	} //else if user.ID != user.UserIDFromContext(ctx) {
	// 	return nomad.Errorf(nomad.EUNAUTHORIZED, "You are not allowed to delete this category.")
	// }

	// Remove row from database.
	if _, err := tx.ExecContext(ctx, `DELETE FROM category WHERE id = ?`, categoryId); err != nil {
		return FormatError(err)
	}
	return nil
}
