package product

import "time"

type ProductData struct {
	Data         []Product `json:"data,omitempty"`
	Count        int       `json:"count"`
	CurrentPage  int       `json:"currentPage"`
	FirstItem    int       `json:"firstItem"`
	FirstPageURL string    `json:"first_page_url"`
	LastItem     int       `json:"lastItem"`
	LastPage     int       `json:"lastPage"`
	LastPageURL  string    `json:"last_page_url"`
	NextPageURL  string    `json:"next_page_url"`
	PerPage      string    `json:"perPage"`
	PrevPageURL  string    `json:"prev_page_url"`
	Total        int       `json:"total"`
}

type Product struct {
	ID              int           `json:"id,omitempty"`
	Name            string        `json:"name,omitempty"`
	Slug            string        `json:"slug,omitempty"`
	Description     string        `json:"description,omitempty"`
	TypeID          int           `json:"type_id,omitempty"`
	Price           int           `json:"price,omitempty"`
	ShopID          int           `json:"shop_id,omitempty"`
	SalePrice       float64       `json:"sale_price,omitempty"`
	Sku             string        `json:"sku,omitempty"`
	Quantity        int           `json:"quantity,omitempty"`
	InStock         int           `json:"in_stock,omitempty"`
	IsTaxable       int           `json:"is_taxable,omitempty"`
	ShippingClassID *string       `json:"shipping_class_id,omitempty"`
	Status          string        `json:"status,omitempty"`
	ProductType     string        `json:"product_type,omitempty"`
	Unit            string        `json:"unit,omitempty"`
	Height          float64       `json:"height,omitempty"`
	Width           float64       `json:"width,omitempty"`
	Length          float64       `json:"length,omitempty"`
	Image           Image         `json:"image,omitempty"`
	Gallery         []Gallery     `json:"gallery,omitempty"`
	DeletedAt       time.Time     `json:"deleted_at,omitempty"`
	CreatedAt       time.Time     `json:"created_at,omitempty"`
	UpdatedAt       time.Time     `json:"updated_at,omitempty"`
	MaxPrice        float64       `json:"max_price,omitempty"`
	MinPrice        float64       `json:"min_price,omitempty"`
	Video           string        `json:"video,omitempty"`
	Type            Type          `json:"type,omitempty"`
	Shop            Shop          `json:"shop,omitempty"`
	Categories      []Category    `json:"categories,omitempty"`
	Tags            []interface{} `json:"tags,omitempty"`
	Variations      []interface{} `json:"variations,omitempty"`
}
type Shop struct {
	ID          int       `json:"id,omitempty"`
	OwnerID     int       `json:"owner_id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Slug        string    `json:"slug,omitempty"`
	Description string    `json:"description,omitempty"`
	CoverImage  Image     `json:"cover_image,omitempty"`
	Logo        Logo      `json:"logo,omitempty"`
	IsActive    int       `json:"is_active,omitempty"`
	Address     Address   `json:"address,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
type Logo struct {
	ID        string `json:"id,omitempty"`
	Original  string `json:"original,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

type Location struct {
	ID               int     `json:"id,omitempty"`
	Lat              float64 `json:"lat,omitempty"`
	Lng              float64 `json:"lng,omitempty"`
	City             string  `json:"city,omitempty"`
	State            string  `json:"state,omitempty"`
	Country          string  `json:"country,omitempty"`
	FormattedAddress string  `json:"formattedAddress,omitempty"`
}

type Category struct {
	ID        int         `json:"id,omitempty"`
	Name      string      `json:"name,omitempty"`
	Slug      string      `json:"slug,omitempty"`
	Icon      *string     `json:"icon,omitempty"`
	Image     []Image     `json:"image,omitempty"`
	Details   string      `json:"details,omitempty"`
	Parent    interface{} `json:"parent,omitempty"`
	TypeID    int         `json:"type_id,omitempty"`
	CreatedAt time.Time   `json:"created_at,omitempty"`
	UpdatedAt time.Time   `json:"updated_at,omitempty"`
	DeletedAt time.Time   `json:"deleted_at,omitempty"`
	ParentID  *string     `json:"parent_id,omitempty"`
	Pivot     Pivot       `json:"pivot,omitempty"`
}

type Pivot struct {
	ID         string `json:"id,omitempty"`
	ProductID  int    `json:"product_id,omitempty"`
	CategoryID int    `json:"category_id,omitempty"`
}

type Image struct {
	ID        string `json:"id,omitempty"`
	Original  string `json:"original,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

type Gallery struct {
	ID        string `json:"id,omitempty"`
	Original  string `json:"original,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

type Type struct {
	ID                int                 `json:"id,omitempty"`
	Name              string              `json:"name,omitempty"`
	Setting           Setting             `json:"settings,omitempty"`
	Slug              string              `json:"slug,omitempty"`
	Icon              string              `json:"icon,omitempty"`
	PromotionalSlider []PromotionalSlider `json:"promotional_sliders,omitempty"`
	CreatedAt         time.Time           `json:"created_at,omitempty"`
	UpdatedAt         time.Time           `json:"updated_at,omitempty"`
}

type Setting struct {
	ID          string   `json:"id,omitempty"`
	Contact     string   `json:"contact,omitempty"`
	Socials     []Social `json:"socials,omitempty"`
	Website     string   `json:"website,omitempty"`
	Location    Location `json:"location,omitempty"`
	IsHome      int      `json:"isHome,omitempty"`
	LayoutType  string   `json:"layoutType,omitempty"`
	ProductCard string   `json:"productCard,omitempty"`
}

type PromotionalSlider struct {
	ID        string `json:"id,omitempty"`
	Original  string `json:"original,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

type Address struct {
	ID            string `json:"id,omitempty"`
	Zip           string `json:"zip,omitempty"`
	City          string `json:"city,omitempty"`
	State         string `json:"state,omitempty"`
	Country       string `json:"country,omitempty"`
	StreetAddress string `json:"street_address,omitempty"`
}

type Social struct {
	ID   string `json:"id,omitempty"`
	URL  string `json:"url,omitempty"`
	Icon string `json:"icon,omitempty"`
}
