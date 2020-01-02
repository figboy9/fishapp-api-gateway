package registry

import (
	"time"

	"github.com/ezio1119/fishapp-api-gateway/domain/auth_grpc"
	"github.com/ezio1119/fishapp-api-gateway/domain/post_grpc"
	"github.com/ezio1119/fishapp-api-gateway/interfaces/presenter"
	"github.com/ezio1119/fishapp-api-gateway/interfaces/repository"
	"github.com/ezio1119/fishapp-api-gateway/interfaces/resolver"
	"github.com/ezio1119/fishapp-api-gateway/usecase/interactor"
)

func NewGraphQLResolver(t time.Duration, postClient post_grpc.PostServiceClient, authClient auth_grpc.AuthServiceClient) *resolver.Resolver {
	return &resolver.Resolver{
		PostInteractor: &interactor.PostInteractor{
			PostRepository: &repository.PostRepository{
				Client: postClient,
			},
			PostPresenter:  &presenter.PostPresenter{},
			ContextTimeout: t,
		},
		UserInteractor: &interactor.UserInteractor{
			UserRepository: &repository.UserRepository{
				Client: authClient,
			},
			UserPresenter:  &presenter.UserPresenter{},
			ContextTimeout: t,
		},
	}
}
