package registry

import (
	"time"

	"github.com/ezio1119/fishapp-api-gateway/domain/auth_grpc"
	"github.com/ezio1119/fishapp-api-gateway/domain/post_grpc"
	"github.com/ezio1119/fishapp-api-gateway/interfaces/resolver/graphql"
)

type registry struct {
	postClient post_grpc.PostServiceClient
	authClient auth_grpc.AuthServiceClient
	timeout    time.Duration
}

type Registry interface {
	NewResolver() graphql.ResolverRoot
}

func NewRegistry(postClient post_grpc.PostServiceClient, authClient auth_grpc.AuthServiceClient, timeout time.Duration) Registry {
	return &registry{postClient, authClient, timeout}
}
