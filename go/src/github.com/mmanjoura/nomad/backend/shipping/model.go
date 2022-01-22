package shipping

import "time"

type Shipping struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Amount    int       `json:"amount,omitempty"`
	IsGlobal  bool      `json:"is_global,omitempty"`
	Type      string    `json:"type,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
