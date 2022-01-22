package attribute

import "time"

type Attribute struct {
	ID        int       `json:"id,omitempty"`
	ShopID    int       `json:"shop_id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Slug      string    `json:"slug,omitempty"`
	Value     []Value   `json:"values,omitempty"`
	Shop      Shop      `json:"shop,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Shop struct {
	ID          int        `json:"id,omitempty"`
	OwnerID     int        `json:"owner_id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Slug        string     `json:"slug,omitempty"`
	Description string     `json:"description,omitempty"`
	CoverImage  CoverImage `json:"cover_image,omitempty"`
	Logo        Logo       `json:"logo,omitempty"`
	IsActive    int        `json:"is_active,omitempty"`
	Address     Address    `json:"address,omitempty"`
	Setting     Setting    `json:"settings,omitempty"`
	//Owner         Owner      `json:"owner,omitempty"`
	OrdersCount   int       `json:"orders_count,omitempty"`
	ProductsCount int       `json:"products_count,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

type Value struct {
	ID          int       `json:"id,omitempty"`
	AttributeID int       `json:"attribute_id,omitempty"`
	Value       string    `json:"value,omitempty"`
	Meta        string    `json:"meta,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type CoverImage struct {
	ID        string `json:"id,omitempty"`
	Original  string `json:"original,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

type Logo struct {
	ID        string `json:"id,omitempty"`
	Original  string `json:"original,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

type Address struct {
	ID            int    `json:"id,omitempty"`
	Zip           string `json:"zip,omitempty"`
	City          string `json:"city,omitempty"`
	State         string `json:"state,omitempty"`
	Country       string `json:"country,omitempty"`
	StreetAddress string `json:"street_address,omitempty"`
}

type Setting struct {
	ID       int      `json:"id,omitempty"`
	Contact  string   `json:"contact,omitempty"`
	Social   []Social `json:"social,omitempty"`
	Website  string   `json:"website,omitempty"`
	Location Location `json:"location,omitempty"`
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

type Social struct {
	ID   int    `json:"id,omitempty"`
	URL  string `json:"url,omitempty"`
	Icon string `json:"icon,omitempty"`
}

type Owner struct {
	ID              int       `json:"id,omitempty"`
	Name            string    `json:"name,omitempty"`
	Email           string    `json:"email,omitempty"`
	EmailVerifiedAt time.Time `json:"email_verified_at,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	IsActive        int       `json:"is_active,omitempty"`
	ShopID          int       `json:"shop_id,omitempty"`
	Profile         Profile   `json:"profile,omitempty"`
}

type Avatar struct {
	ID        string `json:"id,omitempty"`
	Original  string `json:"original,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

type Profile struct {
	ID         int       `json:"id,omitempty"`
	AvatarID   int       `json:"avatar_Id,omitempty"`
	Avatar     Avatar    `json:"avatar,omitempty"`
	Bio        string    `json:"bio,omitempty"`
	Social     []Social  `json:"socials,omitempty"`
	Contact    string    `json:"contact,omitempty"`
	CustomerID int       `json:"customer_id,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}
