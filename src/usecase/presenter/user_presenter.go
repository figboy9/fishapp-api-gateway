package presenter

import (
	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	"github.com/ezio1119/fishapp-api-gateway/domain/user_grpc"
)

type UserPresenter interface {
	TransformUserGraphQL(*user_grpc.User) (*graphql.User, error)
}
