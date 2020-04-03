package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ezio1119/fishapp-api-gateway/graph/dataloader"
	"github.com/ezio1119/fishapp-api-gateway/graph/generated"
	"github.com/ezio1119/fishapp-api-gateway/graph/gqlerr"
	"github.com/ezio1119/fishapp-api-gateway/graph/model"
	"github.com/ezio1119/fishapp-api-gateway/grpc/post_grpc"
	"github.com/ezio1119/fishapp-api-gateway/grpc/profile_grpc"
)

func (r *applyPostResolver) Profile(ctx context.Context, obj *post_grpc.ApplyPost) (*profile_grpc.Profile, error) {
	return dataloader.For(ctx).ProfileByUserID.Load(obj.UserId)
}

func (r *applyPostResolver) Post(ctx context.Context, obj *post_grpc.ApplyPost) (*post_grpc.Post, error) {
	return r.postClient.GetPost(ctx, &post_grpc.GetPostReq{Id: obj.PostId})
}

func (r *mutationResolver) CreatePost(ctx context.Context, input model.CreatePostInput) (*model.CreatePostPayload, error) {
	c, err := getClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	uID, err := strconv.ParseInt(c.User.ID, 10, 64)
	if err != nil {
		return nil, err
	}
	fmt.Println("キトや")

	p, err := r.postClient.CreatePost(ctx, &post_grpc.CreatePostReq{
		Title:             input.Title,
		Content:           input.Content,
		FishingSpotTypeId: input.FishingSpotTypeID,
		FishTypeIds:       input.FishTypeIds,
		PrefectureId:      input.PrefectureID,
		MeetingPlaceId:    input.MeetingPlaceID,
		MeetingAt:         &input.MeetingAt,
		MaxApply:          input.MaxApply,
		UserId:            uID,
	})
	if err != nil {
		return nil, err
	}
	return &model.CreatePostPayload{Post: p}, nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, input model.UpdatePostInput) (*model.UpdatePostPayload, error) {
	c, err := getClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	uID, err := strconv.ParseInt(c.User.ID, 10, 64)
	if err != nil {
		return nil, err
	}
	res, err := r.postClient.GetPost(ctx, &post_grpc.GetPostReq{Id: input.ID})
	if err != nil {
		return nil, err
	}
	if res.UserId != uID {
		return nil, gqlerr.ForbiddenError("user_id=%d does not have permission to update post_id=%d", uID, input.ID)
	}
	p, err := r.postClient.UpdatePost(ctx, &post_grpc.UpdatePostReq{
		Id:                input.ID,
		Title:             input.Title,
		Content:           input.Content,
		FishingSpotTypeId: input.FishingSpotTypeID,
		FishTypeIds:       input.FishTypeIds,
		PrefectureId:      input.PrefectureID,
		MeetingPlaceId:    input.MeetingPlaceID,
		MeetingAt:         &input.MeetingAt,
		MaxApply:          input.MaxApply,
	})
	if err != nil {
		return nil, err
	}
	return &model.UpdatePostPayload{Post: p}, nil
}

func (r *mutationResolver) DeletePost(ctx context.Context, input model.DeletePostInput) (*model.DeletePostPayload, error) {
	c, err := getClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	uID, err := strconv.ParseInt(c.User.ID, 10, 64)
	if err != nil {
		return nil, err
	}
	res, err := r.postClient.GetPost(ctx, &post_grpc.GetPostReq{Id: input.ID})
	if err != nil {
		return nil, err
	}
	if res.UserId != uID {
		return nil, gqlerr.ForbiddenError("user_id=%d does not have permission to delete post_id=%d", uID, input.ID)
	}
	if _, err := r.postClient.DeletePost(ctx, &post_grpc.DeletePostReq{Id: input.ID}); err != nil {
		return nil, err
	}
	return &model.DeletePostPayload{Success: true}, nil
}

func (r *mutationResolver) CreateApplyPost(ctx context.Context, input model.CreateApplyPostInput) (*model.CreateApplyPostPayload, error) {
	c, err := getClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	uID, err := strconv.ParseInt(c.User.ID, 10, 64)
	if err != nil {
		return nil, err
	}
	res, err := r.postClient.GetPost(ctx, &post_grpc.GetPostReq{Id: input.PostID})
	if err != nil {
		return nil, err
	}
	if res.UserId == uID {
		return nil, gqlerr.ForbiddenError("cannot apply your own post")
	}
	a, err := r.postClient.CreateApplyPost(ctx, &post_grpc.CreateApplyPostReq{
		PostId: input.PostID,
		UserId: uID,
	})
	if err != nil {
		return nil, err
	}
	return &model.CreateApplyPostPayload{ApplyPost: a}, nil
}

func (r *mutationResolver) DeleteApplyPost(ctx context.Context, input model.DeleteApplyPostInput) (*model.DeleteApplyPostPayload, error) {
	c, err := getClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	uID, err := strconv.ParseInt(c.User.ID, 10, 64)
	if err != nil {
		return nil, err
	}
	res, err := r.postClient.GetApplyPost(ctx, &post_grpc.GetApplyPostReq{Id: input.ID})
	if err != nil {
		return nil, err
	}
	if res.UserId != uID {
		return nil, gqlerr.ForbiddenError("user_id=%d does not have permission to delete apply_post_id=%d: %s", uID, input.ID, err.Error())
	}
	if _, err := r.postClient.DeleteApplyPost(ctx, &post_grpc.DeleteApplyPostReq{Id: input.ID}); err != nil {
		return nil, err
	}
	return &model.DeleteApplyPostPayload{Success: true}, nil
}

func (r *postResolver) ApplyPosts(ctx context.Context, obj *post_grpc.Post) ([]*post_grpc.ApplyPost, error) {
	return dataloader.For(ctx).ApplyPostsByPostIDs.Load(obj.Id)
}

func (r *postResolver) Profile(ctx context.Context, obj *post_grpc.Post) (*profile_grpc.Profile, error) {
	return dataloader.For(ctx).ProfileByUserID.Load(obj.UserId)
}

func (r *queryResolver) Posts(ctx context.Context, first *int64, after *string, input model.PostsInput) (*model.PostConnection, error) {
	req := &post_grpc.ListPostsReq{Filter: &post_grpc.ListPostsReq_Filter{
		MeetingAtFrom: input.MeetingAtFrom,
		MeetingAtTo:   input.MeetingAtTo,
		FishTypeIds:   input.FishTypeIds,
	}}
	// grpcはnilを受け入れないのでnilチェック
	if input.PrefectureID != nil {
		req.Filter.PrefectureId = *input.PrefectureID
	}
	if input.FishingSpotTypeID != nil {
		req.Filter.FishingSpotTypeId = *input.FishingSpotTypeID
	}
	if input.CanApply != nil {
		req.Filter.CanApply = *input.CanApply
	}
	if input.OrderBy != nil {
		req.Filter.OrderBy = *input.OrderBy
	}
	if input.SortBy != nil {
		req.Filter.SortBy = *input.SortBy
	}
	if input.UserID != nil {
		req.Filter.UserId = *input.UserID
	}
	if first != nil {
		req.PageSize = *first
	}
	if after != nil {
		req.PageToken = *after
	}

	res, err := r.postClient.ListPosts(ctx, req)
	if err != nil {
		return nil, err
	}
	c := &model.PostConnection{
		PageInfo: &model.PageInfo{},
		Nodes:    res.Posts,
	}
	if res.NextPageToken != "" {
		c.PageInfo.HasNextPage = true
		c.PageInfo.EndCursor = &res.NextPageToken
	}
	return c, nil
}

func (r *queryResolver) Post(ctx context.Context, id int64) (*post_grpc.Post, error) {
	return r.postClient.GetPost(ctx, &post_grpc.GetPostReq{Id: id})
}

// ApplyPost returns generated.ApplyPostResolver implementation.
func (r *resolver) ApplyPost() generated.ApplyPostResolver { return &applyPostResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *resolver) Post() generated.PostResolver { return &postResolver{r} }

type applyPostResolver struct{ *resolver }
type postResolver struct{ *resolver }
