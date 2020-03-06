// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package graph

import (
	"context"
	"fmt"

	"github.com/ezio1119/fishapp-api-gateway/graph/generated"
	"github.com/ezio1119/fishapp-api-gateway/graph/model"
	"github.com/ezio1119/fishapp-api-gateway/grpc/post_grpc"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input model.CreatePostInput) (*model.CreatePostPayload, error) {
	c, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}
	res, err := r.postClient.CreatePost(ctx, &post_grpc.CreatePostReq{
		Title:   input.Title,
		Content: input.Content,
		UserId:  c.UserID,
	})
	if err != nil {
		return nil, err
	}
	p, err := convertPostGQL(res.Post)
	if err != nil {
		return nil, err
	}
	return &model.CreatePostPayload{Post: p}, nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, input model.UpdatePostInput) (*model.UpdatePostPayload, error) {
	c, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}
	res, err := r.postClient.UpdatePost(ctx, &post_grpc.UpdatePostReq{
		Id:      input.ID,
		Title:   input.Title,
		Content: input.Content,
		UserId:  c.UserID,
	})
	if err != nil {
		return nil, err
	}
	p, err := convertPostGQL(res.Post)
	if err != nil {
		return nil, err
	}
	return &model.UpdatePostPayload{Post: p}, nil
}

func (r *mutationResolver) DeletePost(ctx context.Context, input model.DeletePostInput) (*model.DeletePostPayload, error) {
	c, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}
	res, err := r.postClient.DeletePost(ctx, &post_grpc.DeletePostReq{
		Id:     input.ID,
		UserId: c.UserID,
	})
	if err != nil {
		return nil, err
	}
	return &model.DeletePostPayload{Success: res.Success}, nil
}

func (r *mutationResolver) CreateApplyPost(ctx context.Context, input model.CreateApplyPostInput) (*model.CreateApplyPostPayload, error) {
	c, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}
	res, err := r.postClient.CreateApplyPost(ctx, &post_grpc.CreateApplyPostReq{
		PostId: input.PostID,
		UserId: c.UserID,
	})
	if err != nil {
		return nil, err
	}
	a, err := convertApplyPostGQL(res.ApplyPost)
	if err != nil {
		return nil, err
	}
	return &model.CreateApplyPostPayload{ApplyPost: a}, nil
}

func (r *mutationResolver) DeleteApplyPost(ctx context.Context, input model.DeleteApplyPostInput) (*model.DeleteApplyPostPayload, error) {
	c, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}
	res, err := r.postClient.DeleteApplyPost(ctx, &post_grpc.DeleteApplyPostReq{Id: input.ApplyPostID, UserId: c.UserID})
	if err != nil {
		return nil, err
	}
	return &model.DeleteApplyPostPayload{Success: res.Success}, nil
}

func (r *queryResolver) Posts(ctx context.Context, first *int64, after *string, input model.PostsInput) (*model.PostConnection, error) {
	res, err := r.postClient.GetListPosts(ctx, &post_grpc.GetListPostsReq{Num: 100})
	if err != nil {
		return nil, err
	}
	posts, err := convertPostsGQL(res.Posts)
	if err != nil {
		return nil, err
	}
	if after != nil {
		id, err := extractIDFromCursor(*after)
		if err != nil {
			return nil, err
		}
		for i, p := range posts {
			if p.ID == id {
				posts = posts[i+1:]
			}
		}
	}
	pageInfo := &model.PageInfo{}

	if first != nil {
		if *first <= 0 {
			return nil, fmt.Errorf("first is 1 or more")
		}
		if len(posts) > int(*first) {
			pageInfo.HasNextPage = true
			posts = posts[:*first]
			cursor := genCursorFromID(posts[len(posts)-1].ID)
			pageInfo.EndCursor = &cursor
		}
	}

	return &model.PostConnection{PageInfo: pageInfo, Nodes: posts}, nil
}

func (r *queryResolver) Post(ctx context.Context, id int64) (*model.Post, error) {
	res, err := r.postClient.GetPostByID(ctx, &post_grpc.GetPostByIDReq{Id: id})
	if err != nil {
		return nil, err
	}
	p, err := convertPostGQL(res.Post)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *resolver }
