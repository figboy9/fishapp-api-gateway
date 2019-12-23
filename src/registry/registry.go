package registry

import (
	"time"

	"github.com/ezio1119/fishapp-api-gateway/domain/post_grpc"
	"github.com/ezio1119/fishapp-api-gateway/interfaces/presenter"
	"github.com/ezio1119/fishapp-api-gateway/interfaces/repository"
	"github.com/ezio1119/fishapp-api-gateway/interfaces/resolver"
	"github.com/ezio1119/fishapp-api-gateway/usecase/interactor"
)

func NewGraphQLResolver(t time.Duration, postClient post_grpc.PostServiceClient) *resolver.Resolver {
	return &resolver.Resolver{
		PostInteractor: &interactor.PostInteractor{
			PostRepository: &repository.PostRepository{
				Client: postClient,
			},
			PostPresenter:  &presenter.PostPresenter{},
			ContextTimeout: t,
		},
	}
}
