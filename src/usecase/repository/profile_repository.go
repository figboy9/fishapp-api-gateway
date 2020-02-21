package repository

import (
	"context"

	"github.com/ezio1119/fishapp-api-gateway/domain/profile_grpc"
	"github.com/golang/protobuf/ptypes/wrappers"
)

type ProfileRepository interface {
	Create(ctx context.Context, req *profile_grpc.CreateReq) (*profile_grpc.Profile, error)
	GetByUserID(ctx context.Context, userID *profile_grpc.ID) (*profile_grpc.Profile, error)
	Update(ctx context.Context, req *profile_grpc.UpdateReq) (*profile_grpc.Profile, error)
	Delete(ctx context.Context, userID *profile_grpc.ID) (*wrappers.BoolValue, error)
}
