//go:generate go run github.com/99designs/gqlgen
package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/ezio1119/fishapp-api-gateway/graph/generated"
	"github.com/ezio1119/fishapp-api-gateway/pb"
	"github.com/nats-io/stan.go"
)

type resolver struct {
	postClient  pb.PostServiceClient
	userClient  pb.UserServiceClient
	chatClient  pb.ChatServiceClient
	imageClient pb.ImageServiceClient
	natsConn    stan.Conn
}

func NewResolver(
	p pb.PostServiceClient,
	a pb.UserServiceClient,
	c pb.ChatServiceClient,
	i pb.ImageServiceClient,
	n stan.Conn,
) generated.ResolverRoot {
	return &resolver{p, a, c, i, n}
}
