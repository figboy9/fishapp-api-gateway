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
	Posts(ctx context.Context, req *post_grpc.ListReq) ([]*graphql.Post, error)
	CreatePost(ctx context.Context, req *post_grpc.CreateReq) (*graphql.Post, error)
	UpdatePost(ctx context.Context, req *post_grpc.UpdateReq) (*graphql.Post, error)
}

func (p *PostInteractor) Post(ctx context.Context, id *post_grpc.ID) (*graphql.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, p.ContextTimeout)
	defer cancel()
	postRPC, err := p.PostRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return p.PostPresenter.TransformPostGraphQL(postRPC)
}

func (p *PostInteractor) Posts(ctx context.Context, listReq *post_grpc.ListReq) ([]*graphql.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, p.ContextTimeout)
	defer cancel()
	listPost, err := p.PostRepository.GetList(ctx, listReq)
	if err != nil {
		return nil, err
	}
	return p.PostPresenter.TransformListPostGraphQL(listPost.Posts)
}

func (p *PostInteractor) CreatePost(ctx context.Context, req *post_grpc.CreateReq) (*graphql.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, p.ContextTimeout)
	defer cancel()
	postRPC, err := p.PostRepository.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return p.PostPresenter.TransformPostGraphQL(postRPC)
}

func (p *PostInteractor) UpdatePost(ctx context.Context, req *post_grpc.UpdateReq) (*graphql.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, p.ContextTimeout)
	defer cancel()
	postRPC, err := p.PostRepository.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return p.PostPresenter.TransformPostGraphQL(postRPC)
}
