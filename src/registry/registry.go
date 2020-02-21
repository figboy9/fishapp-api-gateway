package registry

import (
	"time"

	"github.com/ezio1119/fishapp-api-gateway/domain/auth_grpc"
	"github.com/ezio1119/fishapp-api-gateway/domain/chat_grpc"
	"github.com/ezio1119/fishapp-api-gateway/domain/entry_post_grpc"
	"github.com/ezio1119/fishapp-api-gateway/domain/post_grpc"
	"github.com/ezio1119/fishapp-api-gateway/domain/profile_grpc"
	"github.com/ezio1119/fishapp-api-gateway/interfaces/resolver/graphql"
)

type registry struct {
	postClient      post_grpc.PostServiceClient
	entryPostClient entry_post_grpc.EntryServiceClient
	authClient      auth_grpc.AuthServiceClient
	profileClient   profile_grpc.ProfileServiceClient
	chatClient      chat_grpc.ChatServiceClient
	timeout         time.Duration
}

type Registry interface {
	NewResolver() graphql.ResolverRoot
}

func NewRegistry(
	pc post_grpc.PostServiceClient,
	ec entry_post_grpc.EntryServiceClient,
	ac auth_grpc.AuthServiceClient,
	prc profile_grpc.ProfileServiceClient,
	cc chat_grpc.ChatServiceClient,
	t time.Duration,
) Registry {
	return &registry{pc, ec, ac, prc, cc, t}
}
