package repository

import (
	"context"

	"github.com/ezio1119/fishapp-api-gateway/domain/post_grpc"
	"github.com/golang/protobuf/ptypes/wrappers"
)

type PostRepository interface {
	GetByID(ctx context.Context, id *post_grpc.ID) (*post_grpc.Post, error)
	Create(ctx context.Context, req *post_grpc.CreateReq) (*post_grpc.Post, error)
	GetList(ctx context.Context, req *post_grpc.ListReq) (*post_grpc.ListPost, error)
	Update(ctx context.Context, req *post_grpc.UpdateReq) (*post_grpc.Post, error)
	Delete(ctx context.Context, req *post_grpc.DeleteReq) (*wrappers.BoolValue, error)
}
