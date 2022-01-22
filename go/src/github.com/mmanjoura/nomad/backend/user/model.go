package user

import (
	"time"
)

// User represents a user in the system. Users are typically created via OAuth
// using the AuthService but users can also be created directly for testing.
type User struct {
	ID int `json:"id,omitempty"`

	// User's preferred name & email.
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`

	// Randomly generated API key for use with the CLI.
	APIKey          string `json:"-,omitempty"`
	EmailVerifiedAt string `json:"email_verified_at,omitempty"`
	IsActive        int    `json:"is_active,omitempty"`
	ShopId          int    `json:"shop_Id,omitempty"`

	// Timestamps for user creation & last update.
	CreatedAt time.Time `json:"created_At,omitempty"`
	UpdatedAt time.Time `json:"updated_At,omitempty"`

	// List of associated OAuth authentication objects.
	// Currently only GitHub is supported so there should only be a maximum of one.
	Auths []*Auth `json:"auths,omitempty"`
}

// UserFilter represents a filter passed to FindUsers().
type UserFilter struct {
	// Filtering fields.
	ID     *int    `json:"id,omitempty"`
	Email  *string `json:"email,omitempty"`
	APIKey *string `json:"apiKey,omitempty"`

	// Restrict to subset of results.
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}

// UserUpdate represents a set of fields to be updated via UpdateUser().
type UserUpdate struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

// Auth represents a set of OAuth credentials. These are linked to a User so a
// single user could authenticate through multiple providers.
//
// The authentication system links users by email address, however, some GitHub
// users don't provide their email publicly so we may not be able to link them
// by email address. It's a moot point, however, as we only support GitHub as
// an OAuth provider.
type Auth struct {
	ID int `json:"id,omitempty"`

	// User can have one or more methods of authentication.
	// However, only one per source is allowed per user.
	UserID int   `json:"userID,omitempty"`
	User   *User `json:"user,omitempty"`

	// The authentication source & the source provider's user ID.
	// Source can only be "github" currently.
	Source   string `json:"source,omitempty"`
	SourceID string `json:"sourceID,omitempty"`

	// OAuth fields returned from the authentication provider.
	// GitHub does not use refresh tokens but the field exists for future providers.
	AccessToken  string     `json:"-,omitempty"`
	RefreshToken string     `json:"-,omitempty"`
	Expiry       *time.Time `json:"-,omitempty"`

	// Timestamps of creation & last update.
	CreatedAt time.Time `json:"created_At,omitempty"`
	UpdatedAt time.Time `json:"updated_At,omitempty"`
}

// AuthFilter represents a filter accepted by FindAuths().
type AuthFilter struct {
	// Filtering fields.
	ID       *int    `json:"id,omitempty"`
	UserID   *int    `json:"userID,omitempty"`
	Source   *string `json:"source,omitempty"`
	SourceID *string `json:"sourceID,omitempty"`

	// Restricts results to a subset of the total range.
	// Can be used for pagination.
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}

// A shop is created by a user and can only be edited & deleted by the user who
// created it. Members can be added by sharing an invite link and accepting the
// invitation.
//
// The WTF level for the shop will immediately change when a member's WTF level
// changes and the change will be announced to all other members in real-time.
//
// See the EventService for more information about notifications.
type Shop struct {
	ID int `json:"id,omitempty"`

	// Owner of the shop. Only the owner may delete the shop.
	UserID int   `json:"user_Id,omitempty"`
	User   *User `json:"user,omitempty"`

	// Human-readable name of the shop.
	Name string `json:"name,omitempty"`

	// Code used to share the shop with other users.
	// It allows the creation of a shareable link without explicitly inviting users.
	InviteCode string `json:"inviteCode,omitempty,omitempty"`

	// Aggregate WTF level for the shop. This is a computed field based on the
	// average value of each member's WTF level.
	Value int `json:"value,omitempty"`

	// Timestamps for shop creation & last update.
	CreatedAt time.Time `json:"created_At,omitempty"`
	UpdatedAt time.Time `json:"updated_At,omitempty"`

	// List of associated members and their contributing WTF level.
	// This is only set when returning a single shop.
	Memberships []*ShopMembership `json:"memberships,omitempty,omitempty"`
}

// ShopFilter represents a filter used by FindShops().
type ShopFilter struct {
	// Filtering fields.
	ID         *int    `json:"id,omitempty"`
	InviteCode *string `json:"inviteCode,omitempty"`

	// Restrict to subset of range.
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}

// ShopUpdate represents a set of fields to update on a shop.
type ShopUpdate struct {
	Name *string `json:"name,omitempty"`
}

// ShopValueReport represents a report generated by AverageShopValueReport().
// Each record represents the average value within an interval of time.
type ShopValueReport struct {
	Records []*ShopValueRecord
}

// ShopValueRecord represents an average shop value at a given point in time
// for the ShopValueReport.
type ShopValueRecord struct {
	Value     int       `json:"value,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

// ShopMembership represents a contributor to a Shop. Each membership is
// aggregated to determine the total WTF value of the parent shop.
//
// All members can view all other member's values in the shop. However, only the
// membership owner can edit the membership value.
type ShopMembership struct {
	ID int `json:"id,omitempty"`

	// Parent shop. This shop's WTF level updates when a membership updates.
	ShopID int   `json:"shop_Id,omitempty"`
	Shop   *Shop `json:"shop,omitempty"`

	// Owner of the membership. Only this user can update the membership.
	UserID int   `json:"user_Id,omitempty"`
	User   *User `json:"user,omitempty"`

	// Current WTF level for the user for this shop.
	// Updating this value will cause the parent shop's WTF level to be recomputed.
	Value int `json:"value,omitempty"`

	// Timestamps for membership creation & last update.
	CreatedAt time.Time `json:"created_At,omitempty"`
	UpdatedAt time.Time `json:"updated_At,omitempty"`
}

// ShopMembershipFilter represents a filter used by FindShopMemberships().
type ShopMembershipFilter struct {
	ID     *int `json:"id,omitempty"`
	ShopID *int `json:"shop_Id,omitempty"`
	UserID *int `json:"user_Id,omitempty"`

	// Restricts to a subset of the results.
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`

	// Sorting option for results.
	SortBy string `json:"sortBy,omitempty"`
}

// ShopMembershipUpdate represents a set of fields to update on a membership.
type ShopMembershipUpdate struct {
	Value *int `json:"value,omitempty"`
}
