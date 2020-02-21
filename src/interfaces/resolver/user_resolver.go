package resolver

import (
	"context"
	"strconv"

	"github.com/ezio1119/fishapp-api-gateway/domain/auth_grpc"
	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	"github.com/ezio1119/fishapp-api-gateway/domain/profile_grpc"
	gen "github.com/ezio1119/fishapp-api-gateway/interfaces/resolver/graphql"
)

func (r *queryResolver) User(ctx context.Context, id string) (*graphql.User, error) {
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	return r.userInteractor.User(ctx, &auth_grpc.ID{Id: intID})
}

func (r *profileResolver) User(ctx context.Context, obj *graphql.Profile) (*graphql.User, error) {
	intID, err := strconv.ParseInt(obj.UserID, 10, 64)
	if err != nil {
		return nil, err
	}
	return r.userInteractor.User(ctx, &auth_grpc.ID{Id: intID})
}

func (r *mutationResolver) UpdateUser(ctx context.Context, in gen.UpdateUserInput) (*graphql.User, error) {
	token, err := getJwtTokenCtx(ctx)
	if err != nil {
		return nil, err
	}
	return r.userInteractor.UpdateUser(ctx, &auth_grpc.UpdateReq{
		Email:    in.Email,
		Password: in.Password,
	}, token)
}

func (r *mutationResolver) DeleteUserProfile(ctx context.Context) (bool, error) {
	token, err := getJwtTokenCtx(ctx)
	if err != nil {
		return false, err
	}
	jwtClaims, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return false, err
	}
	return r.userInteractor.DeleteUserProfile(ctx, token, &profile_grpc.ID{
		UserId: jwtClaims.UserID,
	})
}

func (r *mutationResolver) Login(ctx context.Context, in gen.LoginInput) (*gen.UserWithToken, error) {
	return r.userInteractor.Login(ctx, &auth_grpc.LoginReq{
		Email:    in.Email,
		Password: in.Password,
	})
}

func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	token, err := getJwtTokenCtx(ctx)
	if err != nil {
		return false, err
	}
	return r.userInteractor.Logout(ctx, token)
}
func (r *mutationResolver) RefreshIDToken(ctx context.Context) (*graphql.TokenPair, error) {
	token, err := getJwtTokenCtx(ctx)
	if err != nil {
		return nil, err
	}
	return r.userInteractor.RefreshIDToken(ctx, token)
}
