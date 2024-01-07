package dto

import "time"

// Student represents the student entity as stored in the database.
type Student struct {
	ID        string    `bson:"_id,omitempty"`
	FirstName string    `bson:"first_name"`
	LastName  string    `bson:"last_name"`
	Age       int       `bson:"age"`
	Grade     int       `bson:"grade"`
	CreatedAt time.Time `bson:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty"`
}
