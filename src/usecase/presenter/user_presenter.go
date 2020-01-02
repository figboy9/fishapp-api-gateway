package presenter

import (
	"github.com/ezio1119/fishapp-api-gateway/domain/auth_grpc"
	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	gen "github.com/ezio1119/fishapp-api-gateway/interfaces/resolver/graphql"
)

type UserPresenter interface {
	TransformUserGraphQL(*auth_grpc.User) (*graphql.User, error)
	TransformUserWithTokenGraphQL(*auth_grpc.UserWithToken) (*gen.UserWithToken, error)
	TransformTokenPairGraphQL(*auth_grpc.TokenPair) *graphql.TokenPair
}
