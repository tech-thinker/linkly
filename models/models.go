package models

// URL struct information about URL
type URL struct {
	ID        uint64 `json:"id" gorm:"primary_key"`
	URL       string `json:"url"`
	ShortURL  string `json:"short_url"`
	Expires   string `json:"expires"`
	Visits    int    `json:"visits"`
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime"`
}

// Response struct
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"urls"`
}

// // Input get input from request
// type Input struct {
// 	// gorm.Model
// 	URL      string `json:"url"`
// 	ShortURL string `json:"short_url"`
// 	Expires  string `json:"expires"`
// }

// User struct
type User struct {
	ID        uint64 `json:"id" gorm:"primary_key"`
	UserID    string `json:"user_id" gorm:"default:uuid_generate_v3()"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime"`
}
