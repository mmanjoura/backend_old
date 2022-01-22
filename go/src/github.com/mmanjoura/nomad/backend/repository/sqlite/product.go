package sqlite

import (
	"context"
	"strings"

	"github.com/mmanjoura/nomad/backend"
	product "github.com/mmanjoura/nomad/backend/product"
)

var _ product.ProductService = (*ProductService)(nil)

type ProductService struct {
	db *DB
}

func NewProductService(db *DB) *ProductService {
	return &ProductService{db: db}
}

func (s *ProductService) FindOne(ctx context.Context, id int) (*product.Product, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Fetch product and their associated OAuth objects.
	product, err := findProductByID(ctx, tx, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) FindAll(ctx context.Context, productId int, filter backend.Filter) (*product.ProductData, int, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return &product.ProductData{}, 0, err
	}
	defer tx.Rollback()
	return findProducts(ctx, tx, productId, filter)
}

func (s *ProductService) Create(ctx context.Context, productId int, product *product.Product) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Create a new product object and attach associated OAuth objects.
	if err := createProduct(ctx, tx, productId, product); err != nil {
		return err
	}
	return tx.Commit()
}

func (s *ProductService) Update(ctx context.Context, id int, c product.Product) (*product.Product, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Update product & attach associated OAuth objects.
	product, err := updateProduct(ctx, tx, id, c)
	if err != nil {
		return product, err
	} else if err := tx.Commit(); err != nil {
		return product, err
	}
	return product, nil
}

func (s *ProductService) Delete(ctx context.Context, id int) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := deleteProduct(ctx, tx, id); err != nil {
		return err
	}
	return tx.Commit()
}

func findProductByID(ctx context.Context, tx *Tx, productId int) (*product.Product, error) {
	a, _, err := findProducts(ctx, tx, productId, backend.Filter{ID: &productId})
	if err != nil {
		return nil, err
	} else if len(a.Data) == 0 {
		return nil, &backend.Error{Code: backend.ENOTFOUND, Message: "Product not found."}
	}
	return &a.Data[0], nil
}

func findProducts(ctx context.Context, tx *Tx, productId int, filter backend.Filter) (_ *product.ProductData, n int, err error) {

	// Build WHERE clause.
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.ID; v != nil {
		where, args = append(where, "id = ?"), append(args, *v)
	}

	// Execute query to fetch product childeren values rows.
	rows, err := tx.QueryContext(ctx, GetProducts+strings.Join(where, " AND ")+`
		ORDER BY shop.id ASC
		`+FormatLimitOffset(filter.Limit, filter.Offset),
		args...,
	)
	if err != nil {
		return &product.ProductData{}, n, err
	}
	defer rows.Close()

	products := make([]product.Product, 0)
	for rows.Next() {
		var p product.Product
		var cType product.Type
		var setting product.Setting
		var shop product.Shop
		var product_image product.Image
		var logo product.Logo

		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Slug,
			&p.Description,
			&p.TypeID,
			&p.Price,
			&p.ShopID,
			&p.SalePrice,
			&p.Sku,
			&p.IsTaxable,
			&p.ShippingClassID,
			&p.Status,
			&p.ProductType,
			&p.Unit,
			&p.Height,
			&p.Width,
			&p.Length,
			(*NullTime)(&p.DeletedAt),
			(*NullTime)(&p.CreatedAt),
			(*NullTime)(&p.UpdatedAt),
			&p.Video,

			&cType.ID,
			&cType.Name,
			&cType.Slug,
			&cType.Icon,
			(*NullTime)(&cType.CreatedAt),
			(*NullTime)(&cType.UpdatedAt),

			&setting.IsHome,
			&setting.LayoutType,
			&setting.ProductCard,

			&shop.ID,
			&shop.OwnerID,
			&shop.Name,
			&shop.Slug,
			&shop.Description,
			&shop.IsActive,
			(*NullTime)(&shop.CreatedAt),
			(*NullTime)(&shop.UpdatedAt),
			&product_image.ID,
			&product_image.Original,
			&product_image.Thumbnail,

			&logo.ID,
			&logo.Original,
			&logo.Thumbnail,
		); err != nil {
			return &product.ProductData{}, 0, err
		}
		shop.Logo = logo
		shop.CoverImage = product_image
		cType.Setting = setting
		p.Type = cType
		p.Shop = shop
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return &product.ProductData{}, 0, err
	}

	// Get sub structs
	for _, v := range products {
		rows, err = tx.QueryContext(ctx, GetCategories+strings.Join(where, " AND ")+`
			ORDER BY category.id ASC
			`+FormatLimitOffset(filter.Limit, filter.Offset),
			args...,
		)
		if err != nil {
			return &product.ProductData{}, n, err
		}
		defer rows.Close()

		categories := make([]product.Category, 0)
		for rows.Next() {
			var catgr product.Category
			if err := rows.Scan(
				&catgr.ID,
				&catgr.Name,
				&catgr.Slug,
				&catgr.Icon,
				&catgr.TypeID,
				(*NullTime)(&catgr.CreatedAt),
				(*NullTime)(&catgr.UpdatedAt),
				(*NullTime)(&catgr.DeletedAt),
				&catgr.ParentID,
			); err != nil {
				return &product.ProductData{}, 0, err
			}
			categories = append(categories, catgr)
		}

		if err := rows.Err(); err != nil {
			return &product.ProductData{}, 0, err
		}
		v.Categories = categories

		// Execute query to fetch product childeren values rows.
		rows, err = tx.QueryContext(ctx, GetPromotionalSlider+strings.Join(where, " AND ")+`
	ORDER BY id ASC
	`+FormatLimitOffset(filter.Limit, filter.Offset),
			args...,
		)
		if err != nil {
			return &product.ProductData{}, n, err
		}
		defer rows.Close()
		// Deserialize rows into sliders objects.
		promotionalSliders := make([]product.PromotionalSlider, 0)
		for rows.Next() {
			var pSlider product.PromotionalSlider
			if err := rows.Scan(
				&pSlider.ID,
				&pSlider.Original,
				&pSlider.Thumbnail,
			); err != nil {
				return &product.ProductData{}, 0, err
			}

			promotionalSliders = append(promotionalSliders, pSlider)
		}
		if err := rows.Err(); err != nil {
			return &product.ProductData{}, 0, err
		}
		v.Type.PromotionalSlider = promotionalSliders

	}

	productData := product.ProductData{
		Data:         products,
		Total:        10,
		CurrentPage:  1,
		Count:        10,
		LastPage:     1,
		FirstItem:    0,
		LastItem:     9,
		PerPage:      "1000",
		FirstPageURL: "http://localhost:3000/api/categories?search=type.slug:grocery&limit=1000&parent=null&page=1",
		LastPageURL:  "http://localhost:3000/api/categories?search=type.slug:grocery&limit=1000&parent=null&page=1",
		NextPageURL:  "nil",
		PrevPageURL:  "nil",
	}

	productData.Data = products

	return &productData, n, nil
}

func createProduct(ctx context.Context, tx *Tx, productId int, product *product.Product) error {
	// Set timestamps to the current time.
	product.CreatedAt = tx.now
	product.UpdatedAt = product.CreatedAt

	//Get this from Db
	product.ID = productId

	// Perform basic field validation.
	if err := product.Validate(); err != nil {
		return err
	}

	// Execute insertion query.
	result, err := tx.ExecContext(ctx, `
		INSERT INTO product (
			Shop_id,
			slug,
			name,
			created_at,
			updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`,
		product.Slug,
		product.Name,
		product.ID,
		(*NullTime)(&product.CreatedAt),
		(*NullTime)(&product.UpdatedAt),
	)
	if err != nil {
		return FormatError(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	product.ID = int(id)

	return nil
}

func updateProduct(ctx context.Context, tx *Tx, id int, attr product.Product) (*product.Product, error) {
	// Fetch current object state.
	product, err := findProductByID(ctx, tx, id)
	if err != nil {
		return product, err
	} //else if product.ID != product.AttributeIDFromContext(ctx) {
	// 	return nil, nomad.Errorf(nomad.EUNAUTHORIZED, "You are not allowed to update this product.")
	// }

	// Update fields.
	if v := attr.Name; v != "" {
		product.Name = v
	}
	if v := attr.Slug; v != "" {
		product.Slug = v
	}

	// Set last updated date to current time.
	product.UpdatedAt = tx.now

	// Perform basic field validation.
	if err := product.Validate(); err != nil {
		return product, err
	}

	// Execute update query.
	if _, err := tx.ExecContext(ctx, `
		UPDATE product
		SET slug = ?,
		    name = ?,
		    updated_at = ?
		WHERE id = ?
	`,
		product.Slug,
		product.Name,
		(*NullTime)(&product.UpdatedAt),
		id,
	); err != nil {
		return product, FormatError(err)
	}

	return product, nil
}

func deleteProduct(ctx context.Context, tx *Tx, productId int) error {
	// Verify object exists.
	if _, err := findUserByID(ctx, tx, productId); err != nil {
		return err
	} //else if user.ID != user.UserIDFromContext(ctx) {
	// 	return nomad.Errorf(nomad.EUNAUTHORIZED, "You are not allowed to delete this product.")
	// }

	// Remove row from database.
	if _, err := tx.ExecContext(ctx, `DELETE FROM product WHERE id = ?`, productId); err != nil {
		return FormatError(err)
	}
	return nil
}
