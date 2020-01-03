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

func (*profileRepository) Create(ctx context.Context, req *profile_grpc.CreateReq) (*profile_grpc.Profile, error) {
	panic("not implement")
}
func (*profileRepository) GetByUserID(ctx context.Context, id *profile_grpc.ID) (*profile_grpc.Profile, error) {
	panic("not implement")
}
func (*profileRepository) Update(ctx context.Context, id *profile_grpc.ID) (*profile_grpc.Profile, error) {
	panic("not implement")
}
func (*profileRepository) Delete(ctx context.Context, id *profile_grpc.ID) (*wrappers.BoolValue, error) {
	panic("not implement")
}
