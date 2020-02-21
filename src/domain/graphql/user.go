package graphql

import "time"

type User struct {
	ID        string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
