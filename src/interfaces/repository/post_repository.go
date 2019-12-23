package repository

import (
	"context"

	"github.com/ezio1119/fishapp-api-gateway/domain/post_grpc"
)

type PostRepository struct {
	Client post_grpc.PostServiceClient
}

func (r *PostRepository) GetByID(ctx context.Context, id *post_grpc.ID) (*post_grpc.Post, error) {
	return r.Client.GetByID(ctx, id)
}

func (r *PostRepository) GetList(ctx context.Context, listReq *post_grpc.ListReq) (*post_grpc.ListPost, error) {
	return r.Client.GetList(ctx, listReq)
}

func (r *PostRepository) Create(ctx context.Context, createReq *post_grpc.CreateReq) (*post_grpc.Post, error) {
	return r.Client.Create(ctx, createReq)
}

func (r *PostRepository) Update(ctx context.Context, req *post_grpc.UpdateReq) (*post_grpc.Post, error) {
	return r.Client.Update(ctx, req)
}

func (r *PostRepository) Delete(ctx context.Context, req *post_grpc.DeleteReq) (*post_grpc.DeleteRes, error) {
	return r.Client.Delete(ctx, req)
}
