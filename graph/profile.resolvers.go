package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/ezio1119/fishapp-api-gateway/graph/model"
	"github.com/ezio1119/fishapp-api-gateway/grpc/profile_grpc"
)

func (r *mutationResolver) CreateProfile(ctx context.Context, input model.CreateProfileInput) (*model.CreateProfilePayload, error) {
	c, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}
	uID, err := strconv.ParseInt(c.User.ID, 10, 64)
	if err != nil {
		return nil, err
	}
	res, err := r.profileClient.CreateProfile(ctx, &profile_grpc.CreateProfileReq{
		Name:         input.Name,
		Introduction: input.Introduction,
		Sex:          input.Sex,
		UserId:       uID,
	})
	if err != nil {
		return nil, err
	}
	return &model.CreateProfilePayload{Profile: res}, nil
}

func (r *mutationResolver) UpdateProfile(ctx context.Context, input model.UpdateProfileInput) (*model.UpdateProfilePayload, error) {
	c, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}
	uID, err := strconv.ParseInt(c.User.ID, 10, 64)
	if err != nil {
		return nil, err
	}
	res, err := r.profileClient.UpdateProfile(ctx, &profile_grpc.UpdateProfileReq{
		Name:         input.Name,
		Introduction: input.Introduction,
		UserId:       uID,
	})
	if err != nil {
		return nil, err
	}
	return &model.UpdateProfilePayload{Profile: res}, nil
}
