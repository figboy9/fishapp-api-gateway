package graphql

import "time"

type EntryPost struct {
	ID        string
	UserID    string
	PostID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
