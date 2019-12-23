package resolver

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	"github.com/ezio1119/fishapp-api-gateway/domain/post_grpc"
	gen "github.com/ezio1119/fishapp-api-gateway/interfaces/resolver/graphql"
	"github.com/golang/protobuf/ptypes"
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
	// fmt.Println("始まり")
	n, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	return r.PostInteractor.Post(ctx, &post_grpc.ID{Id: n})
}

func (r *queryResolver) Posts(ctx context.Context, in *gen.GetPostListInput) ([]*graphql.Post, error) {
	datetime, err := ptypes.TimestampProto(in.Datetime)
	if err != nil {
		return nil, err
	}
	listReq := &post_grpc.ListReq{
		Datetime: datetime,
		Num:      int64(in.Num),
	}
	return r.PostInteractor.Posts(ctx, listReq)
}

func (r *mutationResolver) CreatePost(ctx context.Context, in gen.CreatePostInput) (*graphql.Post, error) {
	userID, err := getUserIDCtx(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println("useridだよ！！", userID)
	createReq := &post_grpc.CreateReq{
		Title:   in.Title,
		Content: in.Content,
		UserId:  userID,
	}
	return r.PostInteractor.CreatePost(ctx, createReq)
}

func (r *mutationResolver) UpdatePost(ctx context.Context, in gen.UpdatePostInput) (*graphql.Post, error) {
	userID, err := getUserIDCtx(ctx)
	if err != nil {
		return nil, err
	}
	n, err := strconv.ParseInt(in.ID, 10, 64)
	updateReq := &post_grpc.UpdateReq{
		Id:      n,
		Title:   in.Title,
		Content: in.Content,
		UserId:  userID,
	}
	return r.PostInteractor.UpdatePost(ctx, updateReq)
}
