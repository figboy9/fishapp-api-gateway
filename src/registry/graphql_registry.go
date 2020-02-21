package registry

import (
	"github.com/ezio1119/fishapp-api-gateway/interfaces/presenter"
	"github.com/ezio1119/fishapp-api-gateway/interfaces/repository"
	"github.com/ezio1119/fishapp-api-gateway/interfaces/resolver"
	"github.com/ezio1119/fishapp-api-gateway/interfaces/resolver/graphql"
	"github.com/ezio1119/fishapp-api-gateway/usecase/interactor"
)

func (r *registry) NewResolver() graphql.ResolverRoot {
	return resolver.NewResolver(
		interactor.NewUserInteractor(
			repository.NewUserRepository(r.authClient),
			presenter.NewUserPresenter(),
			interactor.NewProfileInteractor(
				repository.NewProfileRepository(r.profileClient),
				presenter.NewProfilePresenter(),
				r.timeout,
			),
			r.timeout,
		),
		interactor.NewProfileInteractor(
			repository.NewProfileRepository(r.profileClient),
			presenter.NewProfilePresenter(),
			r.timeout,
		),
		interactor.NewPostInteractor(
			repository.NewPostRepository(r.postClient),
			repository.NewEntryRepository(r.entryPostClient),
			presenter.NewPostPresenter(),
			r.timeout,
		),
		interactor.NewChatInteractor(
			repository.NewChatRepository(r.chatClient),
			presenter.NewChatPresenter(),
			r.timeout,
		),
	)
}
