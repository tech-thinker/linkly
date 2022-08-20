package models

// Domain struct is a struct that holds the data for a domain
type Domain struct {
	// ID is the primary key for the domain. generates uuid using gorm
	ID *string `json:"id" gorm:"primary_key,not null"`
	// Address is the address of the domain
	Address *string `json:"address"`
	// Banned is a boolean that determines if the domain is banned
	Banned *bool `json:"banned,omitempty"`
	// Homepage is the homepage of the domain
	Homepage *string `json:"homepage"`
	// CreatedAt is the time the domain was created
	CreatedAt *string `json:"created_at"`
	// UpdatedAt is the time the domain was updated
	UpdatedAt *string `json:"updated_at"`
}
