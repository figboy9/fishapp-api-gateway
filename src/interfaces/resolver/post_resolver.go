package resolver

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ezio1119/fishapp-api-gateway/domain/entry_post_grpc"
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
	return r.postInteractor.Post(ctx, &post_grpc.ID{Id: intID})
}

func (r *entryPostResolver) Post(ctx context.Context, obj *graphql.EntryPost) (*graphql.Post, error) {
	intID, err := strconv.ParseInt(obj.PostID, 10, 64)
	if err != nil {
		return nil, err
	}
	return r.postInteractor.Post(ctx, &post_grpc.ID{Id: intID})
}

func (r *postResolver) Entries(ctx context.Context, obj *graphql.Post) ([]*graphql.EntryPost, error) {
	intID, err := strconv.ParseInt(obj.ID, 10, 64)
	if err != nil {
		return nil, err
	}
	return r.postInteractor.Entries(ctx, &entry_post_grpc.ID{PostId: intID})
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
	return r.postInteractor.Posts(ctx, listReq)
}

func (r *mutationResolver) CreatePost(ctx context.Context, in gen.CreatePostInput) (*graphql.Post, error) {
	jwtClaims, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}
	return r.postInteractor.CreatePost(ctx, &post_grpc.CreateReq{
		Title:   in.Title,
		Content: in.Content,
		UserId:  jwtClaims.UserID,
	})
}

func (r *mutationResolver) UpdatePost(ctx context.Context, in gen.UpdatePostInput) (*graphql.Post, error) {
	jwtClaims, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}
	intID, err := strconv.ParseInt(in.ID, 10, 64)
	return r.postInteractor.UpdatePost(ctx, &post_grpc.UpdateReq{
		Id:      intID,
		Title:   in.Title,
		Content: in.Content,
		UserId:  jwtClaims.UserID,
	})
}

func (r *mutationResolver) DeletePost(ctx context.Context, id string) (bool, error) {
	jwtClaims, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return false, err
	}
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return false, err
	}
	return r.postInteractor.DeletePost(ctx, &post_grpc.DeleteReq{
		Id:     intID,
		UserId: jwtClaims.UserID,
	})
}

func (r *mutationResolver) CreateEntryPost(ctx context.Context, postID string) (*graphql.EntryPost, error) {
	fmt.Println("aa")
	jwtClaims, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}
	intPostID, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		return nil, err
	}
	return r.postInteractor.CreateEntryPost(ctx, &entry_post_grpc.CreateReq{
		PostId: intPostID,
		UserId: jwtClaims.UserID,
	})
}

func (r *mutationResolver) DeleteEntryPost(ctx context.Context, id string) (bool, error) {
	jwtClaims, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return false, err
	}
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return false, err
	}
	return r.postInteractor.DeleteEntryPost(ctx, &entry_post_grpc.DeleteReq{
		Id:     intID,
		UserId: jwtClaims.UserID,
	})
}
