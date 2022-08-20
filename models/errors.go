package models

// ServiceError has fields for Service errors. All fields with no data will
// be omitted
type ServiceError struct {
	Type   string `json:"type,omitempty"`
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
	Status int    `json:"status,omitempty"`
}
