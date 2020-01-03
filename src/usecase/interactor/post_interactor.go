package interactor

import (
	"context"
	"time"

	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	"github.com/ezio1119/fishapp-api-gateway/domain/post_grpc"
	"github.com/ezio1119/fishapp-api-gateway/usecase/presenter"
	"github.com/ezio1119/fishapp-api-gateway/usecase/repository"
)

type postInteractor struct {
	postRepository repository.PostRepository
	postPresenter  presenter.PostPresenter
	ctxTimeout     time.Duration
}

func NewPostInteractor(r repository.PostRepository, p presenter.PostPresenter, t time.Duration) PostInteractor {
	return &postInteractor{r, p, t}
}

type PostInteractor interface {
	Post(ctx context.Context, id *post_grpc.ID) (*graphql.Post, error)
	Posts(ctx context.Context, req *post_grpc.ListReq) ([]*graphql.Post, error)
	CreatePost(ctx context.Context, req *post_grpc.CreateReq) (*graphql.Post, error)
	UpdatePost(ctx context.Context, req *post_grpc.UpdateReq) (*graphql.Post, error)
	DeletePost(ctx context.Context, req *post_grpc.DeleteReq) (bool, error)
}

func (i *postInteractor) Post(ctx context.Context, id *post_grpc.ID) (*graphql.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	postProto, err := i.postRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return i.postPresenter.TransformPostGraphQL(postProto)
}

func (i *postInteractor) Posts(ctx context.Context, listReq *post_grpc.ListReq) ([]*graphql.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	listPost, err := i.postRepository.GetList(ctx, listReq)
	if err != nil {
		return nil, err
	}
	return i.postPresenter.TransformListPostGraphQL(listPost.Posts)
}

func (i *postInteractor) CreatePost(ctx context.Context, req *post_grpc.CreateReq) (*graphql.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	postProto, err := i.postRepository.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.postPresenter.TransformPostGraphQL(postProto)
}

func (i *postInteractor) UpdatePost(ctx context.Context, req *post_grpc.UpdateReq) (*graphql.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	postProto, err := i.postRepository.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.postPresenter.TransformPostGraphQL(postProto)
}

func (i *postInteractor) DeletePost(ctx context.Context, req *post_grpc.DeleteReq) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	deleteRes, err := i.postRepository.Delete(ctx, req)
	if err != nil {
		return false, err
	}
	return deleteRes.Deleted, nil
}
