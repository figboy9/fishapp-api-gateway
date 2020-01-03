package repository

import (
	"context"

	"github.com/ezio1119/fishapp-api-gateway/domain/post_grpc"
	"github.com/ezio1119/fishapp-api-gateway/usecase/repository"
)

type postRepository struct {
	client post_grpc.PostServiceClient
}

func NewPostRepository(c post_grpc.PostServiceClient) repository.PostRepository {
	return &postRepository{client: c}
}

func (r *postRepository) GetByID(ctx context.Context, id *post_grpc.ID) (*post_grpc.Post, error) {
	return r.client.GetByID(ctx, id)
}

func (r *postRepository) GetList(ctx context.Context, listReq *post_grpc.ListReq) (*post_grpc.ListPost, error) {
	return r.client.GetList(ctx, listReq)
}

func (r *postRepository) Create(ctx context.Context, createReq *post_grpc.CreateReq) (*post_grpc.Post, error) {
	return r.client.Create(ctx, createReq)
}

func (r *postRepository) Update(ctx context.Context, req *post_grpc.UpdateReq) (*post_grpc.Post, error) {
	return r.client.Update(ctx, req)
}

func (r *postRepository) Delete(ctx context.Context, req *post_grpc.DeleteReq) (*post_grpc.DeleteRes, error) {
	return r.client.Delete(ctx, req)
}
