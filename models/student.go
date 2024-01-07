package models

import "time"

// StudentRequest represents the student entity as received in the request body.
type StudentRequest struct {
	FirstName string `json:"first_name" validate:"required,alpha"`
	LastName  string `json:"last_name" validate:"required,alpha"`
	Age       int    `json:"age" validate:"required,gte=1,lte=100"`
	Grade     int    `json:"grade" validate:"required,gte=1,lte=12"`
}

// StudentResponse represents the student entity as returned in the response body.
type StudentResponse struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Age       int       `json:"age"`
	Grade     int       `json:"grade"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
