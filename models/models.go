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

// Response struct
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"urls"`
}

// User struct
type User struct {
	ID        uint64       `json:"id,omitempty" gorm:"primary_key,autoIncrement,not null"`
	UserID    string       `json:"user_id,omitempty" gorm:"default:uuid_generate_v3()"`
	Username  string       `json:"username"  gorm:"unique,not null"`
	Email     string       `json:"email,omitempty"`
	Password  string       `json:"password" gorm:"not null"`
	IPs       []string     `json:"ips,omitempty"`
	Role      string       `json:"role" gorm:"not null"`
	Level     int          `json:"level" gorm:"not null"`
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

// Domain struct
type Domain struct {
	ID uint64 `json:"id"`
}
