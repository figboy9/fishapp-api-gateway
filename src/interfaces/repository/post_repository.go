package repository

import (
	"context"

	"github.com/ezio1119/fishapp-api-gateway/domain/post_grpc"
)

type PostRepository struct {
	Client post_grpc.PostServiceClient
}

func (p *PostRepository) GetByID(ctx context.Context, id *post_grpc.ID) (*post_grpc.Post, error) {
	res, err := p.Client.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PostRepository) Create(ctx context.Context, req *post_grpc.CreateReq) (*post_grpc.Post, error) {
	res, err := p.Client.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
