package sqlite

import (
	"context"
	"strings"

	"github.com/mmanjoura/nomad/backend"
	"github.com/mmanjoura/nomad/backend/category"
	"github.com/mmanjoura/nomad/backend/setting"
)

func Images(ctx context.Context, tx *Tx, filter backend.Filter, categoryId int) (images []category.Image, n int, err error) {

	// Build WHERE clause.
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := &categoryId; v != nil {
		where, args = append(where, "category_type_id = ?"), append(args, *v)
	}

	rows, err := tx.QueryContext(ctx, GetImages+strings.Join(where, " AND ")+`
		ORDER BY id ASC
		`+FormatLimitOffset(filter.Limit, filter.Offset),
		args...,
	)
	if err != nil {
		return nil, n, err
	}
	defer rows.Close()

	for rows.Next() {
		var image category.Image
		if err := rows.Scan(
			&image.Original,
			&image.Thumbnail,
		); err != nil {
			return nil, 0, err
		}

		images = append(images, image)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return images, n, nil
}

func BannerImages(ctx context.Context, tx *Tx, filter backend.Filter, categoryTypeId int) (bImage category.Image, n int, err error) {

	// Build WHERE clause.
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := &categoryTypeId; v != nil {
		where, args = append(where, "category_type_id = ?"), append(args, *v)
	}

	rows, err := tx.QueryContext(ctx, GetBannerImge+strings.Join(where, " AND ")+`
		ORDER BY id ASC
		`+FormatLimitOffset(filter.Limit, filter.Offset),
		args...,
	)
	if err != nil {
		return category.Image{}, n, err
	}
	defer rows.Close()

	for rows.Next() {
		//var image category.Image
		if err := rows.Scan(
			&bImage.ID,
			&bImage.Original,
			&bImage.Thumbnail,
		); err != nil {
			return category.Image{}, 0, err
		}

	}
	if err := rows.Err(); err != nil {
		return category.Image{}, 0, err
	}

	return bImage, n, nil
}

func CategorySettings(ctx context.Context, tx *Tx, filter backend.Filter, categoryId int, shopId int) (sttrings category.Setting, n int, err error) {

	// Build WHERE clause.
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := &categoryId; v != nil {
		where, args = append(where, "category_Type_id = ?"), append(args, *v)
	}
	// if v := &shopId; v != nil {
	// 	where, args = append(where, "shop_id = ?"), append(args, *v)
	// }

	rows, err := tx.QueryContext(ctx, GetCategoryTypeSetting+strings.Join(where, " AND ")+`
	ORDER BY id ASC
	`+FormatLimitOffset(filter.Limit, filter.Offset),
		args...,
	)
	if err != nil {
		return category.Setting{}, n, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&sttrings.IsHome,
			&sttrings.LayoutType,
			&sttrings.ProductCard,
		); err != nil {
			return category.Setting{}, 0, err
		}

	}

	return sttrings, n, err
}
func CategoryTypeSettings(ctx context.Context, tx *Tx, filter backend.Filter, categoryId int, shopId int) (sttrings category.Setting, n int, err error) {

	// Build WHERE clause.
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := &categoryId; v != nil {
		where, args = append(where, "category_id = ?"), append(args, *v)
	}
	if v := &shopId; v != nil {
		where, args = append(where, "shop_id = ?"), append(args, *v)
	}

	rows, err := tx.QueryContext(ctx, GetCategoryTypeSetting+strings.Join(where, " AND ")+`
	ORDER BY id ASC
	`+FormatLimitOffset(filter.Limit, filter.Offset),
		args...,
	)
	if err != nil {
		return category.Setting{}, n, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&sttrings.IsHome,
			&sttrings.LayoutType,
			&sttrings.ProductCard,
		); err != nil {
			return category.Setting{}, 0, err
		}

	}

	return sttrings, n, err
}

func CategoryType(ctx context.Context, tx *Tx, filter backend.Filter, categoryId int, shopId int) (cType category.CategoryType, n int, err error) {

	where, args := []string{"1 = 1"}, []interface{}{}
	if v := &categoryId; v != nil {
		where, args = append(where, "category_id = ?"), append(args, *v)
	}

	rows, err := tx.QueryContext(ctx, GetCategory_types+strings.Join(where, " AND ")+`
		ORDER BY id ASC
		`+FormatLimitOffset(filter.Limit, filter.Offset),
		args...,
	)
	if err != nil {
		return category.CategoryType{}, n, err
	}
	defer rows.Close()

	for rows.Next() {
		//var categoryType category.CategoryType
		if err := rows.Scan(
			&cType.ID,
			&cType.Name,
			&cType.Slug,
			&cType.Icon,
			(*NullTime)(&cType.CreatedAt),
			(*NullTime)(&cType.UpdatedAt),
		); err != nil {
			return category.CategoryType{}, 0, err
		}

		// Call setting for this category type
		setting, n, err := CategorySettings(ctx, tx, filter, categoryId, shopId)
		if err != nil {
			return category.CategoryType{}, n, err
		}
		cType.Settings = setting

		// Call promotional Sliders for this category type
		sliders, n, err := CategoryTypePromotionalSliders(ctx, tx, filter, cType.ID)
		cType.PromotionalSliders = sliders

	}

	return cType, n, nil
}

func CategoryPromotionalSliders(ctx context.Context, tx *Tx, filter backend.Filter, category_Type_id int) (sliders []category.PromotionalSlider, n int, err error) {

	// Build WHERE clause.
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := &category_Type_id; v != nil {
		where, args = append(where, "category_Type_id = ?"), append(args, *v)
	}

	/// Execute query to fetch category childeren values rows.
	rows, err := tx.QueryContext(ctx, GetPromotionalSlider+strings.Join(where, " AND ")+`
	ORDER BY id ASC
	`+FormatLimitOffset(filter.Limit, filter.Offset),
		args...,
	)
	if err != nil {
		return nil, n, err
	}
	defer rows.Close()

	for rows.Next() {
		var pSlider category.PromotionalSlider
		if err := rows.Scan(
			&pSlider.ID,
			&pSlider.Original,
			&pSlider.Thumbnail,
		); err != nil {
			return nil, 0, err
		}

		sliders = append(sliders, pSlider)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return sliders, n, nil
}
func CategoryTypePromotionalSliders(ctx context.Context, tx *Tx, filter backend.Filter, category_Type_id int) (sliders []category.PromotionalSlider, n int, err error) {

	// Build WHERE clause.
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := &category_Type_id; v != nil {
		where, args = append(where, "category_Type_id = ?"), append(args, *v)
	}

	/// Execute query to fetch category childeren values rows.
	rows, err := tx.QueryContext(ctx, GetPromotionalSlider+strings.Join(where, " AND ")+`
	ORDER BY id ASC
	`+FormatLimitOffset(filter.Limit, filter.Offset),
		args...,
	)
	if err != nil {
		return nil, n, err
	}
	defer rows.Close()

	for rows.Next() {
		var pSlider category.PromotionalSlider
		if err := rows.Scan(
			&pSlider.ID,
			&pSlider.Original,
			&pSlider.Thumbnail,
		); err != nil {
			return nil, 0, err
		}

		sliders = append(sliders, pSlider)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return sliders, n, nil
}

func Childrens(ctx context.Context, tx *Tx, filter backend.Filter, parentId int) (cTypes []category.Category, n int, err error) {

	// Build WHERE clause.
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := &parentId; v != nil {
		where, args = append(where, "parent_id = ?"), append(args, *v)
	}

	rows, err := tx.QueryContext(ctx, GetCategoryChildren+strings.Join(where, " AND ")+`
		ORDER BY category_id ASC
		`+FormatLimitOffset(filter.Limit, filter.Offset),
		args...,
	)
	if err != nil {
		return nil, n, err
	}
	defer rows.Close()

	for rows.Next() {
		var catgr category.Category
		if err := rows.Scan(
			&catgr.ID,
			&catgr.Name,
			&catgr.Slug,
			&catgr.Icon,
			&catgr.ParentID,
			&catgr.TypeID,
			(*NullTime)(&catgr.CreatedAt),
			(*NullTime)(&catgr.UpdatedAt),
			(*NullTime)(&catgr.DeletedAt),
		); err != nil {
			return nil, 0, err
		}

		cTypes = append(cTypes, catgr)

	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return cTypes, n, nil
}

func Banners(ctx context.Context, tx *Tx, filter backend.Filter, category_type_id int) (banners []category.Banner, n int, err error) {

	// Build WHERE clause.
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := &category_type_id; v != nil {
		where, args = append(where, "category_type_id = ?"), append(args, *v)
	}

	/// Execute query to fetch category childeren values rows.
	rows, err := tx.QueryContext(ctx, GetBanners+strings.Join(where, " AND ")+`
	ORDER BY id ASC
	`+FormatLimitOffset(filter.Limit, filter.Offset),
		args...,
	)
	if err != nil {
		return nil, n, err
	}
	defer rows.Close()

	for rows.Next() {
		var banner category.Banner
		if err := rows.Scan(
			&banner.ID,
			&banner.TypeID,
			&banner.Title,
			&banner.Description,
			(*NullTime)(&banner.CreatedAt),
			(*NullTime)(&banner.UpdatedAt),
		); err != nil {
			return nil, 0, err
		}

		// Banner images

		image, n, err := BannerImages(ctx, tx, filter, category_type_id)
		if err != nil {
			return nil, n, err
		}
		banner.Image = image

		banners = append(banners, banner)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return banners, n, nil
}

func Socails(ctx context.Context, tx *Tx, filter backend.Filter, setting_id int) (socials []setting.Social, n int, err error) {

	// Build WHERE clause.
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := &setting_id; v != nil {
		where, args = append(where, "setting_id = ?"), append(args, *v)
	}

	/// Execute query to fetch category childeren values rows.
	rows, err := tx.QueryContext(ctx, GetSocials+strings.Join(where, " AND ")+`
	ORDER BY url ASC
	`+FormatLimitOffset(filter.Limit, filter.Offset),
		args...,
	)
	if err != nil {
		return nil, n, err
	}
	defer rows.Close()

	for rows.Next() {
		var social setting.Social
		if err := rows.Scan(
			&social.URL,
			&social.Icon,
		); err != nil {
			return nil, 0, err
		}

		socials = append(socials, social)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return socials, n, nil
}

func DeliveryTimes(ctx context.Context, tx *Tx, filter backend.Filter, setting_id int) (deliveryTimes []setting.DeliveryTime, n int, err error) {

	// Build WHERE clause.
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := &setting_id; v != nil {
		where, args = append(where, "setting_id = ?"), append(args, *v)
	}

	/// Execute query to fetch category childeren values rows.
	rows, err := tx.QueryContext(ctx, GetDeliveryTime+strings.Join(where, " AND ")+`
	ORDER BY title ASC
	`+FormatLimitOffset(filter.Limit, filter.Offset),
		args...,
	)
	if err != nil {
		return nil, n, err
	}
	defer rows.Close()

	for rows.Next() {
		var deliveryTime setting.DeliveryTime
		if err := rows.Scan(
			&deliveryTime.Title,
			&deliveryTime.Description,
		); err != nil {
			return nil, 0, err
		}

		deliveryTimes = append(deliveryTimes, deliveryTime)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return deliveryTimes, n, nil
}
