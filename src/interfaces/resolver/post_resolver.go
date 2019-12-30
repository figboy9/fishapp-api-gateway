package resolver

import (
	"context"
	"strconv"

	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	"github.com/ezio1119/fishapp-api-gateway/domain/post_grpc"
	gen "github.com/ezio1119/fishapp-api-gateway/interfaces/resolver/graphql"
	"github.com/golang/protobuf/ptypes"
)

func (r *queryResolver) Post(ctx context.Context, id string) (*graphql.Post, error) {
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	return r.PostInteractor.Post(ctx, &post_grpc.ID{Id: intID})
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
	jwtClaims, err := getJwtCtx(ctx)
	if err != nil {
		return nil, err
	}
	createReq := &post_grpc.CreateReq{
		Title:   in.Title,
		Content: in.Content,
		UserId:  jwtClaims.UserID,
	}
	return r.PostInteractor.CreatePost(ctx, createReq)
}

func (r *mutationResolver) UpdatePost(ctx context.Context, in gen.UpdatePostInput) (*graphql.Post, error) {
	jwtClaims, err := getJwtCtx(ctx)
	if err != nil {
		return nil, err
	}
	intID, err := strconv.ParseInt(in.ID, 10, 64)
	updateReq := &post_grpc.UpdateReq{
		Id:      intID,
		Title:   in.Title,
		Content: in.Content,
		UserId:  jwtClaims.UserID,
	}
	return r.PostInteractor.UpdatePost(ctx, updateReq)
}

func (r *mutationResolver) DeletePost(ctx context.Context, id string) (bool, error) {
	jwtClaims, err := getJwtCtx(ctx)
	if err != nil {
		return false, err
	}
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return false, err
	}
	deleteReq := &post_grpc.DeleteReq{
		Id:     intID,
		UserId: jwtClaims.UserID,
	}
	return r.PostInteractor.DeletePost(ctx, deleteReq)
}
