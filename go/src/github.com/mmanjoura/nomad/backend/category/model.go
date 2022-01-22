package category

import "time"

type CategoryData struct {
	Data         []Category `json:"data,omitempty"`
	Count        int        `json:"count"`
	CurrentPage  int        `json:"currentPage"`
	FirstItem    int        `json:"firstItem"`
	FirstPageURL string     `json:"first_page_url"`
	LastItem     int        `json:"lastItem"`
	LastPage     int        `json:"lastPage"`
	LastPageURL  string     `json:"last_page_url"`
	NextPageURL  string     `json:"next_page_url"`
	PerPage      string     `json:"perPage"`
	PrevPageURL  string     `json:"prev_page_url"`
	Total        int        `json:"total"`
}

type Category struct {
	ID           int          `json:"id,omitempty"`
	Name         string       `json:"name,omitempty"`
	Slug         string       `json:"slug,omitempty"`
	Icon         *string      `json:"icon,omitempty"`
	Image        []Image      `json:"image"`
	Details      interface{}  `json:"details"`
	Parent       interface{}  `json:"parent"`
	TypeID       int          `json:"type_id"`
	CreatedAt    time.Time    `json:"created_at,omitempty"`
	UpdatedAt    time.Time    `json:"updated_at,omitempty"`
	DeletedAt    time.Time    `json:"deleted_at,omitempty"`
	ParentID     *string      `json:"parent_id"`
	CategoryType CategoryType `json:"type,omitempty"`
	Children     []Category   `json:"children,omitempty"`
}

type Children struct {
	ID        int         `json:"id,omitempty"`
	Name      string      `json:"name,omitempty"`
	Slug      string      `json:"slug,omitempty"`
	Icon      *string     `json:"icon,omitempty"`
	Image     []Image     `json:"image,omitempty"`
	Details   interface{} `json:"details,omitempty"`
	Parent    interface{} `json:"parent,omitempty"`
	TypeID    int         `json:"type_id,omitempty"`
	CreatedAt time.Time   `json:"created_at,omitempty"`
	UpdatedAt time.Time   `json:"updated_at,omitempty"`
	DeletedAt time.Time   `json:"deleted_at,omitempty"`
	ParentID  int         `json:"parent_id,omitempty"`
	Children  []Children  `json:"children,omitempty"`
}
type PromotionalSlider struct {
	ID        string `json:"id,omitempty"`
	Original  string `json:"original,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

type Setting struct {
	ID          string `json:"id,omitempty"`
	IsHome      bool   `json:"isHome"`
	LayoutType  string `json:"layoutType,omitempty"`
	ProductCard string `json:"productCard,omitempty"`
}

type CategoryType struct {
	ID                 int                 `json:"id,omitempty"`
	Name               string              `json:"name,omitempty"`
	Settings           Setting             `json:"settings,omitempty"`
	Slug               string              `json:"slug,omitempty"`
	Icon               *string             `json:"icon,omitempty"`
	PromotionalSliders []PromotionalSlider `json:"promotional_sliders,omitempty"`
	CreatedAt          time.Time           `json:"created_at,omitempty"`
	UpdatedAt          time.Time           `json:"updated_at,omitempty"`
	Banners            []Banner            `json:"banners,omitempty"`
}

type Image struct {
	ID        *string `json:"id,omitempty"`
	Original  *string `json:"original,omitempty"`
	Thumbnail *string `json:"thumbnail,omitempty"`
}

type Banner struct {
	ID          int       `json:"id,omitempty"`
	TypeID      int       `json:"type_id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Image       Image     `json:"image,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
