package presenter

import (
	"strconv"

	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	"github.com/ezio1119/fishapp-api-gateway/domain/user_grpc"
	gen "github.com/ezio1119/fishapp-api-gateway/interfaces/resolver/graphql"
	"github.com/golang/protobuf/ptypes"
)

type UserPresenter struct{}

func (p *UserPresenter) TransformUserGraphQL(u *user_grpc.User) (*graphql.User, error) {
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
		Name:      u.Name,
		Email:     u.Email,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}, nil
}

func (p *UserPresenter) TransformUserWithTokenGraphQL(ut *user_grpc.UserWithToken) (*gen.UserWithToken, error) {
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
func (p *UserPresenter) TransformTokenPairGraphQL(tp *user_grpc.TokenPair) *graphql.TokenPair {
	return &graphql.TokenPair{
		IDToken:      tp.IdToken,
		RefreshToken: tp.RefreshToken,
	}
}
