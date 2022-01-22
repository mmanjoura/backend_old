package order

import "time"

type Order struct {
	ID                int             `json:"id,omitempty"`
	TrackingNumber    string          `json:"tracking_number,omitempty"`
	CustomerID        int             `json:"customer_id,omitempty"`
	CustomerContact   string          `json:"customer_contact,omitempty"`
	Status            Status          `json:"status,omitempty"`
	Amount            float64         `json:"amount,omitempty"`
	SalesTax          float64         `json:"sales_tax,omitempty"`
	PaidTotal         float64         `json:"paid_total,omitempty"`
	Total             float64         `json:"total,omitempty"`
	CouponID          int             `json:"coupon_id,omitempty"`
	ParentID          int             `json:"parent_id,omitempty"`
	ShopID            int             `json:"shop_id,omitempty"`
	Discount          int             `json:"discount,omitempty"`
	PaymentID         int             `json:"payment_id,omitempty"`
	PaymentGateway    string          `json:"payment_gateway,omitempty"`
	ShippingAddress   ShippingAddress `json:"shipping_address,omitempty"`
	BillingAddress    BillingAddress  `json:"billing_address,omitempty"`
	LogisticsProvider string          `json:"logistics_provider,omitempty"`
	DeliveryFee       int             `json:"delivery_fee,omitempty"`
	DeliveryTime      time.Time       `json:"delivery_time,omitempty"`
	DeletedAt         time.Time       `json:"deleted_at,omitempty"`
	CreatedAt         time.Time       `json:"created_at,omitempty"`
	UpdatedAt         time.Time       `json:"updated_at,omitempty"`
	Customer          Customer        `json:"customer,omitempty"`
	Products          []Product       `json:"products,omitempty"`
	Children          []Children      `json:"children,omitempty"`
}
type Image struct {
	ID        string `json:"id,omitempty"`
	Original  string `json:"original,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}
type Status struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Serial    int       `json:"serial,omitempty"`
	Color     string    `json:"color,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type ShippingAddress struct {
	ID            int    `json:"id,omitempty"`
	Zip           string `json:"zip,omitempty"`
	City          string `json:"city,omitempty"`
	State         string `json:"state,omitempty"`
	Country       string `json:"country,omitempty"`
	StreetAddress string `json:"street_address,omitempty"`
}
type BillingAddress struct {
	Zip           string `json:"zip,omitempty"`
	City          string `json:"city,omitempty"`
	State         string `json:"state,omitempty"`
	Country       string `json:"country,omitempty"`
	StreetAddress string `json:"street_address,omitempty"`
}

type Customer struct {
	ID              int       `json:"id,omitempty"`
	Name            string    `json:"name,omitempty"`
	Email           string    `json:"email,omitempty"`
	EmailVerifiedAt time.Time `json:"email_verified_at,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	IsActive        int       `json:"is_active,omitempty"`
	ShopID          int       `json:"shop_id,omitempty"`
}

type Gallery struct {
	ID        string `json:"id,omitempty"`
	Original  string `json:"original,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

type Pivot struct {
	ID                string    `json:"id,omitempty"`
	OrderID           int       `json:"order_id,omitempty"`
	ProductID         int       `json:"product_id,omitempty"`
	OrderQuantity     string    `json:"order_quantity,omitempty"`
	UnitPrice         float64   `json:"unit_price,omitempty"`
	Subtotal          float64   `json:"subtotal,omitempty"`
	VariationOptionID int       `json:"variation_option_id,omitempty"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
	UpdatedAt         time.Time `json:"updated_at,omitempty"`
}

type Product struct {
	ID               int           `json:"id,omitempty"`
	Name             string        `json:"name,omitempty"`
	Slug             string        `json:"slug,omitempty"`
	Description      string        `json:"description,omitempty"`
	TypeID           int           `json:"type_id,omitempty"`
	Price            int           `json:"price,omitempty"`
	ShopID           int           `json:"shop_id,omitempty"`
	SalePrice        float64       `json:"sale_price,omitempty"`
	Sku              string        `json:"sku,omitempty"`
	Quantity         int           `json:"quantity,omitempty"`
	InStock          int           `json:"in_stock,omitempty"`
	IsTaxable        int           `json:"is_taxable,omitempty"`
	ShippingClassID  int           `json:"shipping_class_id,omitempty"`
	Status           string        `json:"status,omitempty"`
	ProductType      string        `json:"product_type,omitempty"`
	Unit             string        `json:"unit,omitempty"`
	Height           float64       `json:"height,omitempty"`
	Width            float64       `json:"width,omitempty"`
	Length           float64       `json:"length,omitempty"`
	Image            Image         `json:"image,omitempty"`
	Gallery          []Gallery     `json:"gallery,omitempty"`
	DeletedAt        time.Time     `json:"deleted_at,omitempty"`
	CreatedAt        time.Time     `json:"created_at,omitempty"`
	UpdatedAt        time.Time     `json:"updated_at,omitempty"`
	MaxPrice         float64       `json:"max_price,omitempty"`
	MinPrice         float64       `json:"min_price,omitempty"`
	Pivot            Pivot         `json:"pivot,omitempty"`
	VariationOptions []interface{} `json:"variation_options,omitempty"`
}

type Children struct {
	ID                int             `json:"id,omitempty"`
	TrackingNumber    string          `json:"tracking_number,omitempty"`
	CustomerID        int             `json:"customer_id,omitempty"`
	CustomerContact   string          `json:"customer_contact,omitempty"`
	Status            Status          `json:"status,omitempty"`
	Amount            float64         `json:"amount,omitempty"`
	SalesTax          int             `json:"sales_tax,omitempty"`
	PaidTotal         float64         `json:"paid_total,omitempty"`
	Total             float64         `json:"total,omitempty"`
	CouponID          int             `json:"coupon_id,omitempty"`
	ParentID          int             `json:"parent_id,omitempty"`
	ShopID            int             `json:"shop_id,omitempty"`
	Discount          int             `json:"discount,omitempty"`
	PaymentID         int             `json:"payment_id,omitempty"`
	PaymentGateway    string          `json:"payment_gateway,omitempty"`
	ShippingAddress   ShippingAddress `json:"shipping_address,omitempty"`
	BillingAddress    BillingAddress  `json:"billing_address,omitempty"`
	LogisticsProvider string          `json:"logistics_provider,omitempty"`
	DeliveryFee       int             `json:"delivery_fee,omitempty"`
	DeliveryTime      string          `json:"delivery_time,omitempty"`
	DeletedAt         time.Time       `json:"deleted_at,omitempty"`
	CreatedAt         time.Time       `json:"created_at,omitempty"`
	UpdatedAt         time.Time       `json:"updated_at,omitempty"`
	Customer          Customer        `json:"customer,omitempty"`
	Products          []Product       `json:"products,omitempty"`
}
