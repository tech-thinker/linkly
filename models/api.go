package models

// QRCode struct
type QRCode struct {
	ID      uint64 `json:"id" gorm:"primary_key"`
	UserID  string `json:"user_id"`
	Content string `json:"content"`
	Image   []byte `json:"image"`
}

// Message is a struct that holds the message for the response
type Message struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error is used as the Response Body
type Error struct {
	Error ServiceError `json:"error"`
}

// LinkBody is the body of the link
type LinkBody struct {
	CustomURL   string `json:"customurl"`
	Target      string `json:"target"`
	Description string `json:"description"`
	Reusable    *bool  `json:"reusable"`
	Password    string `json:"password"`
	ExpireIn    string `json:"expire_in"`
	Domain      string `json:"domain"`
}
