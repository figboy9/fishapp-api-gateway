package resolver

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	"github.com/ezio1119/fishapp-api-gateway/domain/post_grpc"
	gen "github.com/ezio1119/fishapp-api-gateway/interfaces/resolver/graphql"
)

func getUserIDCtx(ctx context.Context) (int64, error) {
	v := ctx.Value(UserIDCtxKey)
	userID, ok := v.(int64)
	if !ok {
		return 0, fmt.Errorf("userID not found")
	}

	return userID, nil
}

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

func (r *queryResolver) Posts(ctx context.Context, in *gen.GetPostListInput) ([]*graphql.Post, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreatePost(ctx context.Context, in gen.CreatePostInput) (*graphql.Post, error) {
	userID, err := getUserIDCtx(ctx)
	if err != nil {
		return nil, err
	}
	createReq := &post_grpc.CreateReq{
		Title:   in.Title,
		Content: in.Content,
		UserId:  userID,
	}
	post, err := r.PostInteractor.Create(ctx, createReq)
	if err != nil {
		return nil, err
	}
	return post, nil
}
