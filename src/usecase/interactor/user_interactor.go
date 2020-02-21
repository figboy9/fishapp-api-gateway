package interactor

import (
	"context"
	"time"

	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	gen "github.com/ezio1119/fishapp-api-gateway/interfaces/resolver/graphql"

	"github.com/ezio1119/fishapp-api-gateway/domain/auth_grpc"
	"github.com/ezio1119/fishapp-api-gateway/domain/profile_grpc"
	"github.com/ezio1119/fishapp-api-gateway/usecase/presenter"
	"github.com/ezio1119/fishapp-api-gateway/usecase/repository"
)

type userInteractor struct {
	userRepository    repository.UserRepository
	userPresenter     presenter.UserPresenter
	profileInteractor ProfileInteractor
	ctxTimeout        time.Duration
}

func NewUserInteractor(r repository.UserRepository, p presenter.UserPresenter, pi ProfileInteractor, t time.Duration) UserInteractor {
	return &userInteractor{r, p, pi, t}
}

type UserInteractor interface {
	User(ctx context.Context, id *auth_grpc.ID) (*graphql.User, error)
	CreateUserProfile(ctx context.Context, userReq *auth_grpc.CreateReq, profileReq *profile_grpc.CreateReq) (*gen.UserProfileWithToken, error)
	DeleteUserProfile(ctx context.Context, token string, userID *profile_grpc.ID) (bool, error)
	UpdateUser(ctx context.Context, req *auth_grpc.UpdateReq, token string) (*graphql.User, error)
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

func (i *userInteractor) CreateUserProfile(ctx context.Context, u *auth_grpc.CreateReq, p *profile_grpc.CreateReq) (*gen.UserProfileWithToken, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	userWithTokenProto, err := i.userRepository.Create(ctx, u)
	if err != nil {
		return nil, err
	}
	p.UserId = userWithTokenProto.User.Id
	profile, err := i.profileInteractor.CreateProfile(ctx, p)
	if err != nil {
		_, _ = i.userRepository.Delete(ctx, userWithTokenProto.TokenPair.IdToken)
		return nil, err
	}
	return i.userPresenter.TransformUserProfileWithTokenGraphQL(userWithTokenProto, profile)
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

func (i *userInteractor) DeleteUserProfile(ctx context.Context, token string, userID *profile_grpc.ID) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	success, err := i.userRepository.Delete(ctx, token)
	if err != nil {
		return false, err
	}
	profileRes, err := i.profileInteractor.DeleteProfile(ctx, userID)
	if err != nil {
		return false, err
	}
	return success.Value && profileRes, nil
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
