package graphql

import "time"

type Message struct {
	ID        string
	Body      string
	RoomID    string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
