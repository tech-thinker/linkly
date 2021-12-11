package models

// URL struct information about URL
type URL struct {
	Id          uint64 `json:"id" redis:"id"`
	OriginalURL string `json:"original_url" redis:"original_url"`
	URL         string `json:"url" redis:"url"`
	Expires     string `json:"expires" redis:"expires"`
	Visits      int    `json:"visits" redis:"visits"`
}

// Response struct
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"shortUrl"`
}

// Input get input from request
var Input struct {
	URL     string `json:"url"`
	Expires string `json:"expires"`
}
