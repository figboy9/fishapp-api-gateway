package repository

import (
	"context"

	"github.com/ezio1119/fishapp-api-gateway/domain/profile_grpc"
	"github.com/ezio1119/fishapp-api-gateway/usecase/repository"
	"github.com/golang/protobuf/ptypes/wrappers"
)

type profileRepository struct {
	client profile_grpc.ProfileServiceClient
}

func NewProfileRepository(c profile_grpc.ProfileServiceClient) repository.ProfileRepository {
	return &profileRepository{client: c}
}

func (r *profileRepository) Create(ctx context.Context, req *profile_grpc.CreateReq) (*profile_grpc.Profile, error) {
	return r.client.Create(ctx, req)
}
func (r *profileRepository) GetByUserID(ctx context.Context, userID *profile_grpc.ID) (*profile_grpc.Profile, error) {
	return r.client.GetByUserID(ctx, userID)
}
func (r *profileRepository) Update(ctx context.Context, req *profile_grpc.UpdateReq) (*profile_grpc.Profile, error) {
	return r.client.Update(ctx, req)
}
func (r *profileRepository) Delete(ctx context.Context, userID *profile_grpc.ID) (*wrappers.BoolValue, error) {
	return r.client.Delete(ctx, userID)
}
