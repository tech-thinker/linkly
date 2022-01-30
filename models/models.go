package models

// URL struct information about URL
type URL struct {
	ID        uint64 `json:"id" gorm:"primary_key"`
	URL       string `json:"url"`
	ShortURL  string `json:"short_url"`
	Expires   string `json:"expires"`
	Visits    int    `json:"visits"`
	CreatedAt string `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime"`
}

// Response struct
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"urls"`
}

// Input get input from request
type Input struct {
	ID       uint64 `json:"id"`
	URL      string `json:"url"`
	ShortURL string `json:"short_url"`
	Expires  string `json:"expires"`
}
