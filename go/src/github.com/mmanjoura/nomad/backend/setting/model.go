package setting

import "time"

type Setting struct {
	ID     int    `json:"id,omitempty"`
	Option Option `json:"options,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
type Seo struct {
	OgImage         *string `json:"ogImage"`
	OgTitle         *string `json:"ogTitle"`
	MetaTags        *string `json:"metaTags"`
	MetaTitle       *string `json:"metaTitle"`
	CanonicalURL    *string `json:"canonicalUrl"`
	OgDescription   *string `json:"ogDescription"`
	TwitterHandle   *string `json:"twitterHandle"`
	MetaDescription *string `json:"metaDescription"`
	TwitterCardType *string `json:"twitterCardType"`
}

type Logo struct {
	ID        int    `json:"id,omitempty"`
	Original  string `json:"original,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

type ContactDetail struct {
	ID      int      `json:"id,omitempty"`
	Contact string   `json:"contact,omitempty"`
	Socials []Social `json:"socials,omitempty"`

	Website  string   `json:"website,omitempty"`
	Location Location `json:"location,omitempty"`
}

type Social struct {
	ID   int    `json:"id,omitempty"`
	URL  string `json:"url,omitempty"`
	Icon string `json:"icon,omitempty"`
}

type Location struct {
	ID               int     `json:"id,omitempty"`
	Lat              float64 `json:"lat,omitempty"`
	Lng              float64 `json:"lng,omitempty"`
	State            string  `json:"state,omitempty"`
	Country          string  `json:"country,omitempty"`
	FormattedAddress string  `json:"formattedAddress,omitempty"`
}

type DeliveryTime struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}
type Option struct {
	ID                 int            `json:"id,omitempty"`
	Seo                Seo            `json:"seo"`
	Logo               Logo           `json:"logo"`
	Currency           string         `json:"currency"`
	TaxClass           int            `json:"taxClass"`
	SiteTitle          string         `json:"siteTitle"`
	DeliveryTime       []DeliveryTime `json:"deliveryTime"`
	SiteSubtitle       string         `json:"siteSubtitle"`
	ShippingClass      int            `json:"shippingClass"`
	ContactDetail      ContactDetail  `json:"contactDetails"`
	MinimumOrderAmount int            `json:"minimumOrderAmount"`
}
