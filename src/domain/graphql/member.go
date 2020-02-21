package graphql

import "time"

type Member struct {
	ID        string
	RoomID    string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
