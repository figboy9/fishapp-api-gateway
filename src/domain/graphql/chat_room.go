package graphql

import "time"

type ChatRoom struct {
	ID        string
	PostID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
