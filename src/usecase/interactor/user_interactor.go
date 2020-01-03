package interactor

import (
	"context"
	"time"

	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	gen "github.com/ezio1119/fishapp-api-gateway/interfaces/resolver/graphql"

	"github.com/ezio1119/fishapp-api-gateway/domain/auth_grpc"
	"github.com/ezio1119/fishapp-api-gateway/usecase/presenter"
	"github.com/ezio1119/fishapp-api-gateway/usecase/repository"
)

type userInteractor struct {
	userRepository repository.UserRepository
	userPresenter  presenter.UserPresenter
	ctxTimeout     time.Duration
}

func NewUserInteractor(r repository.UserRepository, p presenter.UserPresenter, t time.Duration) UserInteractor {
	return &userInteractor{r, p, t}
}

type UserInteractor interface {
	User(ctx context.Context, id *auth_grpc.ID) (*graphql.User, error)
	CreateUser(ctx context.Context, req *auth_grpc.CreateReq) (*gen.UserWithToken, error)
	UpdateUser(ctx context.Context, req *auth_grpc.UpdateReq, token string) (*graphql.User, error)
	DeleteUser(ctx context.Context, token string) (bool, error)
	Login(ctx context.Context, req *auth_grpc.LoginReq) (*gen.UserWithToken, error)
	Logout(ctx context.Context, token string) (bool, error)
	RefreshIDToken(ctx context.Context, token string) (*graphql.TokenPair, error)
}

func (i *userInteractor) User(ctx context.Context, id *auth_grpc.ID) (*graphql.User, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	userProto, err := i.userRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return i.userPresenter.TransformUserGraphQL(userProto)
}

func (i *userInteractor) CreateUser(ctx context.Context, req *auth_grpc.CreateReq) (*gen.UserWithToken, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	userWithTokenProto, err := i.userRepository.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.userPresenter.TransformUserWithTokenGraphQL(userWithTokenProto)
}

func (i *userInteractor) UpdateUser(ctx context.Context, req *auth_grpc.UpdateReq, token string) (*graphql.User, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	userProto, err := i.userRepository.Update(ctx, req, token)
	if err != nil {
		return nil, err
	}
	return i.userPresenter.TransformUserGraphQL(userProto)
}

func (i *userInteractor) DeleteUser(ctx context.Context, token string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	boolValue, err := i.userRepository.Delete(ctx, token)
	if err != nil {
		return false, err
	}
	return boolValue.Value, nil
}

func (i *userInteractor) Login(ctx context.Context, req *auth_grpc.LoginReq) (*gen.UserWithToken, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	userWithTokenProto, err := i.userRepository.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.userPresenter.TransformUserWithTokenGraphQL(userWithTokenProto)
}

func (i *userInteractor) Logout(ctx context.Context, token string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	boolValue, err := i.userRepository.Logout(ctx, token)
	if err != nil {
		return false, err
	}
	return boolValue.Value, nil
}
func (i *userInteractor) RefreshIDToken(ctx context.Context, token string) (*graphql.TokenPair, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	tokenPairProto, err := i.userRepository.RefreshIDToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return i.userPresenter.TransformTokenPairGraphQL(tokenPairProto), nil
}
