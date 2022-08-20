package models

import (
	"database/sql"
	"time"
)

// URL struct information about URL
type URL struct {
	ID        uint64       `json:"id,omitempty" gorm:"primary_key,autoIncrement,not null"`
	URL       string       `json:"url,omitempty"`
	ShortURL  string       `json:"short_url,omitempty"`
	Expires   string       `json:"expires,omitempty"`
	Visits    int          `json:"visits,omitempty"`
	UserID    string       `json:"user_id,omitempty"`
	IP        string       `json:"ip,omitempty"`
	CreatedAt time.Time    `json:"created_at,omitempty" gorm:"not null"`
	UpdatedAt time.Time    `json:"updated_at,omitempty" gorm:"not null"`
	DeletedAt sql.NullTime `json:"deleted_at,omitempty" gorm:"index,not null"`
}

// QRCode struct
type QRCode struct {
	ID      uint64 `json:"id" gorm:"primary_key"`
	UserID  string `json:"user_id"`
	Content string `json:"content"`
	Image   []byte `json:"image"`
}

// Message is a struct that holds the message for the response
type Message struct {
	Message string `json:"message"`
}

// Error is used as the Response Body
type Error struct {
	Error ServiceError `json:"error"`
}
