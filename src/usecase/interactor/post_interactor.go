package interactor

import (
	"context"
	"time"

	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	"github.com/ezio1119/fishapp-api-gateway/domain/post_grpc"
	"github.com/ezio1119/fishapp-api-gateway/usecase/presenter"
	"github.com/ezio1119/fishapp-api-gateway/usecase/repository"
)

type PostInteractor struct {
	PostRepository repository.PostRepository
	PostPresenter  presenter.PostPresenter
	ContextTimeout time.Duration
}

type UPostInteractor interface {
	Post(ctx context.Context, id *post_grpc.ID) (*graphql.Post, error)
	Create(ctx context.Context, req *post_grpc.CreateReq) (*graphql.Post, error)
}

func (p *PostInteractor) Post(ctx context.Context, id *post_grpc.ID) (*graphql.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, p.ContextTimeout)
	defer cancel()
	postRPC, err := p.PostRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	post, err := p.PostPresenter.TransformPostGraphQL(postRPC)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (p *PostInteractor) Create(ctx context.Context, req *post_grpc.CreateReq) (*graphql.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, p.ContextTimeout)
	defer cancel()
	postRPC, err := p.PostRepository.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	post, err := p.PostPresenter.TransformPostGraphQL(postRPC)
	if err != nil {
		return nil, err
	}
	return post, nil
}
