package repository

import (
	"context"

	"github.com/ezio1119/fishapp-api-gateway/domain/post_grpc"
)

type PostRepository interface {
	GetByID(ctx context.Context, id *post_grpc.ID) (*post_grpc.Post, error)
	Create(ctx context.Context, req *post_grpc.CreateReq) (*post_grpc.Post, error)
	GetList(ctx context.Context, id *post_grpc.ListReq) (*post_grpc.ListPost, error)
}
