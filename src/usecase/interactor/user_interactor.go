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
	UpdateUser(ctx context.Context, req *user_grpc.UpdateReq) (*graphql.User, error)
	DeleteUser(ctx context.Context, req *user_grpc.ID) (bool, error)
	Login(ctx context.Context, req *user_grpc.LoginReq) (*gen.UserWithToken, error)
	Logout(ctx context.Context, req *user_grpc.AddBlackListReq) (bool, error)
	RefreshToken(ctx context.Context, req *user_grpc.CheckBlackListReq) (*graphql.TokenPair, error)
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
	userWithTokenRPC, err := i.UserRepository.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.UserPresenter.TransformUserWithTokenGraphQL(userWithTokenRPC)
}

func (i *UserInteractor) UpdateUser(ctx context.Context, req *user_grpc.UpdateReq) (*graphql.User, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ContextTimeout)
	defer cancel()
	userRPC, err := i.UserRepository.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.UserPresenter.TransformUserGraphQL(userRPC)
}

func (i *UserInteractor) DeleteUser(ctx context.Context, req *user_grpc.ID) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ContextTimeout)
	defer cancel()
	boolValue, err := i.UserRepository.Delete(ctx, req)
	if err != nil {
		return false, err
	}
	return boolValue.Value, nil
}

func (i *UserInteractor) Login(ctx context.Context, req *user_grpc.LoginReq) (*gen.UserWithToken, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ContextTimeout)
	defer cancel()
	userWithTokenRPC, err := i.UserRepository.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.UserPresenter.TransformUserWithTokenGraphQL(userWithTokenRPC)
}

func (i *UserInteractor) Logout(ctx context.Context, req *user_grpc.AddBlackListReq) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ContextTimeout)
	defer cancel()
	boolValue, err := i.UserRepository.AddBlackList(ctx, req)
	if err != nil {
		return false, err
	}
	return boolValue.Value, nil
}
func (i *UserInteractor) RefreshToken(ctx context.Context, req *user_grpc.CheckBlackListReq) (*graphql.TokenPair, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ContextTimeout)
	defer cancel()
	tokenPairRPC, err := i.UserRepository.CheckBlackListAndGenToken(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.UserPresenter.TransformTokenPairGraphQL(tokenPairRPC), nil
}
