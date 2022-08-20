package models

import "time"

// Link is a struct that holds the data for a link
type Link struct {
	//ID is the primary key for the link. generates uuid using gorm
	ID string `json:"id" gorm:"primary_key,not null"`
	// Address is the address of the link
	Address *string `json:"address"`
	// Banned is a boolean that determines if the link is banned
	Banned *bool `json:"banned,omitempty"`
	// Link is the unique address that is being stored
	Link *string `json:"link"`
	// Password is the password for the link
	Password *string `json:"password,omitempty"`
	// Target is the target for the link
	Target *string `json:"target"`
	// Description is the description for the link
	Description *string `json:"description"`
	// VisitCount is the number of times the link has been visited
	VisitCount *int `json:"visit_count"`
	// ExipreAt is the time when the link expires
	ExpireAt *time.Time `json:"expire_at,omitempty"`
	// Reusable is a boolean that determines if the link is reusable
	Reusable *bool `json:"reusable,omitempty"`
	// UserID is the user who created the link
	UserID *string `json:"user_id,omitempty"`
	// IP is the ip address of the user who created the link [security,spam]
	IP *string `json:"ip,omitempty"`
	//  CreatedAt is the time the link was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the time the link was last updated
	UpdatedAt time.Time `json:"updated_at"`
}
