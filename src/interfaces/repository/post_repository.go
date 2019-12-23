package repository

import (
	"context"

	"github.com/ezio1119/fishapp-api-gateway/domain/post_grpc"
)

type PostRepository struct {
	Client post_grpc.PostServiceClient
}

func (p *PostRepository) GetByID(ctx context.Context, id *post_grpc.ID) (*post_grpc.Post, error) {
	return p.Client.GetByID(ctx, id)
}

func (p *PostRepository) GetList(ctx context.Context, listReq *post_grpc.ListReq) (*post_grpc.ListPost, error) {
	return p.Client.GetList(ctx, listReq)
}

func (p *PostRepository) Create(ctx context.Context, createReq *post_grpc.CreateReq) (*post_grpc.Post, error) {
	return p.Client.Create(ctx, createReq)
}
