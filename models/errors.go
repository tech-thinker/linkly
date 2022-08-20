package models

// Error is used as the Response Body
type Error struct {
	Error ServiceError `json:"error"`
}

// ServiceError has fields for Service errors. All fields with no data will
// be omitted
type ServiceError struct {
	Type   string `json:"type,omitempty"`
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
	Status int    `json:"status,omitempty"`
}
