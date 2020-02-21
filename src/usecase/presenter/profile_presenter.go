package presenter

import (
	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	"github.com/ezio1119/fishapp-api-gateway/domain/profile_grpc"
)

type ProfilePresenter interface {
	TransformProfileGraphQL(*profile_grpc.Profile) (*graphql.Profile, error)
}
