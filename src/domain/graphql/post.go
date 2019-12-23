package graphql

import "time"

type Post struct {
	ID        string
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    string
}
