package backend

// ShopFilter represents a filter used by FindShops().
type Filter struct {
	// Filtering fields.
	ID *int `json:"id,omitempty"`

	// Restrict to subset of range.
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}
