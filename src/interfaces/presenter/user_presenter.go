package presenter

import (
	"strconv"

	"github.com/ezio1119/fishapp-api-gateway/domain/auth_grpc"
	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	gen "github.com/ezio1119/fishapp-api-gateway/interfaces/resolver/graphql"
	"github.com/ezio1119/fishapp-api-gateway/usecase/presenter"
	"github.com/golang/protobuf/ptypes"
)

type userPresenter struct{}

func NewUserPresenter() presenter.UserPresenter {
	return &userPresenter{}
}

func (p *userPresenter) TransformUserGraphQL(u *auth_grpc.User) (*graphql.User, error) {
	id := strconv.FormatInt(u.Id, 10)
	updatedAt, err := ptypes.Timestamp(u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	createdAt, err := ptypes.Timestamp(u.CreatedAt)
	if err != nil {
		return nil, err
	}
	updatedAt = updatedAt.In(location)
	createdAt = createdAt.In(location)
	return &graphql.User{
		ID:        id,
		Email:     u.Email,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}, nil
}

func (p *userPresenter) TransformUserWithTokenGraphQL(ut *auth_grpc.UserWithToken) (*gen.UserWithToken, error) {
	user, err := p.TransformUserGraphQL(ut.User)
	if err != nil {
		return nil, err
	}
	tokenPair := p.TransformTokenPairGraphQL(ut.TokenPair)
	return &gen.UserWithToken{
		User:      user,
		TokenPair: tokenPair,
	}, nil
}

func (p *userPresenter) TransformUserProfileWithTokenGraphQL(tp *auth_grpc.UserWithToken, profile *graphql.Profile) (*gen.UserProfileWithToken, error) {
	userWithToken, err := p.TransformUserWithTokenGraphQL(tp)
	if err != nil {
		return nil, err
	}
	return &gen.UserProfileWithToken{
		User:      userWithToken.User,
		Profile:   profile,
		TokenPair: userWithToken.TokenPair,
	}, nil
}

func (p *userPresenter) TransformTokenPairGraphQL(tp *auth_grpc.TokenPair) *graphql.TokenPair {
	return &graphql.TokenPair{
		IDToken:      tp.IdToken,
		RefreshToken: tp.RefreshToken,
	}
}
