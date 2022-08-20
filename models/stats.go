package models

import "time"

// Stat is a struct that holds the data for a link
type Stat struct {
	//  AllTime is the number of times the link has been visited
	AllTime StatItem `json:"all_time"`
	// LastDay is the number of times the link has been visited in the last day
	LastDay StatItem `json:"last_day"`
	// LastWeek is the number of times the link has been visited in the last week
	LastWeek StatItem `json:"last_week"`
	// LastMonth is the number of times the link has been visited in the last month
	LastMonth StatItem `json:"last_month"`
	// Address is the address of the link
	Address string `json:"address"`
	// Banned is a boolean that determines if the link is banned
	Banned bool `json:"banned,omitempty"`
	// Link is the unique address that is being stored
	Link string `json:"link"`
	// Password is the password for the link
	Password string `json:"password,omitempty"`
	// Target is the target for the link
	Target string `json:"target"`
	// Description is the description for the link
	Description string `json:"description"`
	// ExipreAt is the time when the link expires
	ExpireAt time.Time `json:"expire_at,omitempty"`
	// Reusable is a boolean that determines if the link is reusable
	Reusable *bool `json:"reusable,omitempty"`
	// UserID is the user who created the link
	UserID string `json:"user_id,omitempty"`
	// IP is the ip address of the user who created the link [security,spam]
	IP string `json:"ip,omitempty"`
	//  CreatedAt is the time the link was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the time the link was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// allTime is the number of times the link has been visited
type StatItem struct {
	// Browser is the browser the link was visited from
	Browser Browser `json:"browser"`
	// OS is the operating system the link was visited from
	OS OS `json:"os"`
	// Country is the country the link was visited from
	Country Country `json:"country"`
	// Referrer is the referrer the link was visited from
	Referrer Referrer `json:"referrer"`
	Views    int      `json:"views"`
}

// Browser is the browser the link was visited from
type Browser struct {
	// Name is the name of the browser
	Name string `json:"name"`
	// Value is the number of times the browser was visited
	Value int `json:"value"`
}

// OS is the operating system the link was visited from
type OS struct {
	// Name is the name of the operating system
	Name string `json:"name"`
	// Value is the number of times the operating system was visited
	Value int `json:"value"`
}

// Country is the country the link was visited from
type Country struct {
	// Name is the name of the country
	Name string `json:"name"`
	// Value is the number of times the country was visited
	Value int `json:"value"`
}

// Referrer is the referrer the link was visited from
type Referrer struct {
	// Name is the name of the referrer
	Name string `json:"name"`
	// Value is the number of times the referrer was visited
	Value int `json:"value"`
}
