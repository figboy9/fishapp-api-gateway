// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package graph

import (
	"context"

	"github.com/ezio1119/fishapp-api-gateway/graph/model"
	"github.com/ezio1119/fishapp-api-gateway/grpc/profile_grpc"
)

func (r *mutationResolver) CreateProfile(ctx context.Context, input model.CreateProfileInput) (*model.CreateProfilePayload, error) {
	c, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}
	res, err := r.profileClient.Create(ctx, &profile_grpc.CreateReq{Name: input.Name, UserId: c.UserID})
	if err != nil {
		return nil, err
	}
	p, err := convertProfileGQL(res)
	if err != nil {
		return nil, err
	}
	return &model.CreateProfilePayload{Profile: p}, nil
}

func (r *mutationResolver) DeleteProfile(ctx context.Context) (*model.DeleteProfilePayload, error) {
	c, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}
	res, err := r.profileClient.DeleteByUserID(ctx, &profile_grpc.ID{UserId: c.UserID})
	if err != nil {
		return nil, err
	}
	return &model.DeleteProfilePayload{Success: res.Value}, nil
}

func (r *mutationResolver) UpdateProfile(ctx context.Context, input model.UpdateProfileInput) (*model.UpdateProfilePayload, error) {
	c, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}
	res, err := r.profileClient.UpdateByUserID(ctx, &profile_grpc.UpdateReq{Name: input.Name, UserId: c.UserID})
	if err != nil {
		return nil, err
	}
	p, err := convertProfileGQL(res)
	if err != nil {
		return nil, err
	}
	return &model.UpdateProfilePayload{Profile: p}, nil
}

func (r *queryResolver) Profile(ctx context.Context, userID int64) (*model.Profile, error) {
	res, err := r.profileClient.GetByUserID(ctx, &profile_grpc.ID{UserId: userID})
	if err != nil {
		return nil, err
	}
	return convertProfileGQL(res)
}
