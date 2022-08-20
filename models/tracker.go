package models

import (
	"time"
)

// Tracker is a struct that holds the data for a tracker
type Tracker struct {
	ID          string    `json:"id" gorm:"primary_key"`
	Link        string    `json:"link"`
	Description *string   `json:"description,omitempty"`
	VisitCount  int       `json:"visit_count"`
	Image       []byte    `json:"image"`
	UserID      *string   `json:"user_id,omitempty"`
	IP          *string   `json:"ip,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TrackerStatus is the status of the tracker
type TrackerStatus struct {
	Message string `json:"message"`
}
