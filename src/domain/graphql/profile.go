package graphql

import "time"

type Profile struct {
	ID        string
	Name      string
	UserID    string 
	CreatedAt time.Time
	UpdatedAt time.Time
}
