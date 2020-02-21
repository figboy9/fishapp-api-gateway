package resolver

import (
	"context"
	"strconv"

	"github.com/ezio1119/fishapp-api-gateway/domain/auth_grpc"
	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	"github.com/ezio1119/fishapp-api-gateway/domain/profile_grpc"
	gen "github.com/ezio1119/fishapp-api-gateway/interfaces/resolver/graphql"
)

func (r *queryResolver) Profile(ctx context.Context, userID string) (*graphql.Profile, error) {
	intUserID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return nil, err
	}
	return r.profileInteractor.Profile(ctx, &profile_grpc.ID{
		UserId: intUserID,
	})
}

func (r *postResolver) Profile(ctx context.Context, obj *graphql.Post) (*graphql.Profile, error) {
	intUserID, err := strconv.ParseInt(obj.UserID, 10, 64)
	if err != nil {
		return nil, err
	}
	return r.profileInteractor.Profile(ctx, &profile_grpc.ID{
		UserId: intUserID,
	})
}

func (r *entryPostResolver) Profile(ctx context.Context, obj *graphql.EntryPost) (*graphql.Profile, error) {
	intUserID, err := strconv.ParseInt(obj.UserID, 10, 64)
	if err != nil {
		return nil, err
	}
	return r.profileInteractor.Profile(ctx, &profile_grpc.ID{
		UserId: intUserID,
	})
}

func (r *mutationResolver) CreateUserProfile(ctx context.Context, in gen.CreateUserProfileInput) (*gen.UserProfileWithToken, error) {
	return r.userInteractor.CreateUserProfile(ctx, &auth_grpc.CreateReq{
		Email:    in.Email,
		Password: in.Password,
	}, &profile_grpc.CreateReq{
		Name: in.Name,
	})
}

func (r *mutationResolver) UpdateProfile(ctx context.Context, in gen.UpdateProfileInput) (*graphql.Profile, error) {
	jwtClaims, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}
	return r.profileInteractor.UpdateProfile(ctx, &profile_grpc.UpdateReq{
		Name:   in.Name,
		UserId: jwtClaims.UserID,
	})
}
