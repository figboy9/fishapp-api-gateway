package interactor

import (
	"context"
	"time"

	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	gen "github.com/ezio1119/fishapp-api-gateway/interfaces/resolver/graphql"

	"github.com/ezio1119/fishapp-api-gateway/domain/user_grpc"
	"github.com/ezio1119/fishapp-api-gateway/usecase/presenter"
	"github.com/ezio1119/fishapp-api-gateway/usecase/repository"
)

type UserInteractor struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
	ContextTimeout time.Duration
}

type UUserInteractor interface {
	User(ctx context.Context, id *user_grpc.ID) (*graphql.User, error)
	CreateUser(ctx context.Context, req *user_grpc.CreateReq) (*gen.UserWithToken, error)
}

func (i *UserInteractor) User(ctx context.Context, id *user_grpc.ID) (*graphql.User, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ContextTimeout)
	defer cancel()
	userRPC, err := i.UserRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return i.UserPresenter.TransformUserGraphQL(userRPC)
}

func (i *UserInteractor) CreateUser(ctx context.Context, req *user_grpc.CreateReq) (*gen.UserWithToken, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ContextTimeout)
	defer cancel()
	userWithToken, err := i.UserRepository.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.UserPresenter.TransformUserWithTokenGraphQL(userWithToken)
}
