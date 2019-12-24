package presenter

import (
	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	"github.com/ezio1119/fishapp-api-gateway/domain/user_grpc"
	gen "github.com/ezio1119/fishapp-api-gateway/interfaces/resolver/graphql"
)

type UserPresenter interface {
	TransformUserGraphQL(*user_grpc.User) (*graphql.User, error)
	TransformUserWithTokenGraphQL(*user_grpc.UserWithToken) (*gen.UserWithToken, error)
	TransformTokenPairGraphQL(*user_grpc.TokenPair) *graphql.TokenPair
}
