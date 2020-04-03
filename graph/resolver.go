package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"cloud.google.com/go/pubsub"
	"github.com/ezio1119/fishapp-api-gateway/graph/generated"
	"github.com/ezio1119/fishapp-api-gateway/grpc/auth_grpc"
	"github.com/ezio1119/fishapp-api-gateway/grpc/chat_grpc"
	"github.com/ezio1119/fishapp-api-gateway/grpc/post_grpc"
	"github.com/ezio1119/fishapp-api-gateway/grpc/profile_grpc"
)

type resolver struct {
	postClient    post_grpc.PostServiceClient
	authClient    auth_grpc.AuthServiceClient
	profileClient profile_grpc.ProfileServiceClient
	chatClient    chat_grpc.ChatServiceClient
	pubsubClient  *pubsub.Client
}

func NewResolver(
	p post_grpc.PostServiceClient,
	a auth_grpc.AuthServiceClient,
	pro profile_grpc.ProfileServiceClient,
	c chat_grpc.ChatServiceClient,
	pubsub *pubsub.Client,
) generated.ResolverRoot {
	return &resolver{p, a, pro, c, pubsub}
}
