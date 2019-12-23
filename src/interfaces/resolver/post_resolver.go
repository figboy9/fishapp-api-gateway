package resolver

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	"github.com/ezio1119/fishapp-api-gateway/domain/post_grpc"
	gen "github.com/ezio1119/fishapp-api-gateway/interfaces/resolver/graphql"
)

func (r *queryResolver) Post(ctx context.Context, id string) (*graphql.Post, error) {
	fmt.Println("始まり")
	n, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	post, err := r.PostInteractor.Post(ctx, &post_grpc.ID{Id: n})
	if err != nil {
		return nil, err
	}
	fmt.Println("終わり")
	return post, nil
}

func (r *mutationResolver) CreatePost(ctx context.Context, in gen.CreatePostInput) (*graphql.Post, error) {
	userID := ctx.Value("userID").(int64)
	createReq := &post_grpc.CreateReq{
		Title:   in.Title,
		Content: in.Content,
	}
	post, err := r.PostInteractor.Create(ctx, createReq)
	if err != nil {
		return nil, err
	}
	return post, nil
}
