package coupon

import "time"

type Coupon struct {
	ID          int       `json:"id,omitempty"`
	Code        string    `json:"code,omitempty"`
	Description string    `json:"description,omitempty"`
	Image       Image     `json:"image,omitempty"`
	Type        string    `json:"type,omitempty"`
	Amount      int       `json:"amount,omitempty"`
	ActiveFrom  time.Time `json:"active_from,omitempty"`
	ExpireAt    time.Time `json:"expire_at,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
	IsValid     int       `json:"is_valid,omitempty"`
}

type Image struct {
	ID        int    `json:"id,omitempty"`
	Original  string `json:"original,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}
