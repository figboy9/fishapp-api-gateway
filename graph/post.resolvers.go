package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/ezio1119/fishapp-api-gateway/conf"
	"github.com/ezio1119/fishapp-api-gateway/graph/dataloader"
	"github.com/ezio1119/fishapp-api-gateway/graph/generated"
	"github.com/ezio1119/fishapp-api-gateway/graph/gqlerr"
	"github.com/ezio1119/fishapp-api-gateway/graph/model"
	"github.com/ezio1119/fishapp-api-gateway/pb"
	stan "github.com/nats-io/stan.go"
	"google.golang.org/protobuf/encoding/protojson"
)

func (r *applyPostResolver) User(ctx context.Context, obj *pb.ApplyPost) (*pb.User, error) {
	return r.userClient.GetUser(ctx, &pb.GetUserReq{Id: obj.UserId})
}

func (r *applyPostResolver) Post(ctx context.Context, obj *pb.ApplyPost) (*pb.Post, error) {
	return r.postClient.GetPost(ctx, &pb.GetPostReq{Id: obj.PostId})
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

	req := &pb.CreatePostReq{
		Data: &pb.CreatePostReq_Info{
			Info: &pb.CreatePostReqInfo{
				Title:             input.Title,
				Content:           input.Content,
				FishingSpotTypeId: input.FishingSpotTypeID,
				FishTypeIds:       input.FishTypeIds,
				PrefectureId:      input.PrefectureID,
				MeetingPlaceId:    input.MeetingPlaceID,
				MeetingAt:         &input.MeetingAt,
				MaxApply:          input.MaxApply,
				UserId:            uID,
			}},
	}

	stream, err := r.postClient.CreatePost(ctx)
	if err != nil {
		return nil, err
	}

	if err := stream.Send(req); err != nil {
		return nil, err
	}

	for n, image := range input.Images {
		for {
			buf := make([]byte, conf.C.Sv.ChunkDataSize)
			n, err := image.File.Read(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, gqlerr.InternalServerError("cannot read chunk to buffe: %s", err)
			}

			req := &pb.CreatePostReq{
				Data: &pb.CreatePostReq_ImageChunk{
					ImageChunk: buf[:n],
				},
			}

			if stream.Send(req); err != nil {
				return nil, fmt.Errorf("cannot send chunk to server: %w", err)
			}
		}

		if len(input.Images) != n+1 {

			req = &pb.CreatePostReq{
				Data: &pb.CreatePostReq_NextImageSignal{NextImageSignal: true},
			}

			if err := stream.Send(req); err != nil {
				log.Fatal("cannot send chunk to server: ", err)
			}
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return nil, err
	}

	return &model.CreatePostPayload{Post: res.Post, SagaID: res.SagaId}, nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, input model.UpdatePostInput) (*model.UpdatePostPayload, error) {
	req := &pb.UpdatePostReq{
		Data: &pb.UpdatePostReq_Info{
			Info: &pb.UpdatePostReqInfo{
				Id:                input.ID,
				Title:             input.Title,
				Content:           input.Content,
				FishingSpotTypeId: input.FishingSpotTypeID,
				FishTypeIds:       input.FishTypeIds,
				PrefectureId:      input.PrefectureID,
				MeetingPlaceId:    input.MeetingPlaceID,
				MeetingAt:         &input.MeetingAt,
				MaxApply:          input.MaxApply,
				ImageIdsToDelete:  input.ImageIdsToDelete,
			}},
	}

	stream, err := r.postClient.UpdatePost(ctx)
	if err != nil {
		return nil, err
	}

	if err := stream.Send(req); err != nil {
		return nil, err
	}

	for n, image := range input.Images {
		for {
			buf := make([]byte, conf.C.Sv.ChunkDataSize)
			n, err := image.File.Read(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, gqlerr.InternalServerError("cannot read chunk to buffe: %s", err)
			}

			req = &pb.UpdatePostReq{
				Data: &pb.UpdatePostReq_ImageChunk{
					ImageChunk: buf[:n],
				},
			}

			if stream.Send(req); err != nil {
				return nil, fmt.Errorf("cannot send chunk to server: %w", err)
			}
		}

		if len(input.Images) != n+1 { // まだイメージがあるときにNextImageSignalを送る
			req = &pb.UpdatePostReq{
				Data: &pb.UpdatePostReq_NextImageSignal{NextImageSignal: true},
			}

			if err := stream.Send(req); err != nil {
				log.Fatal("cannot send chunk to server: ", err)
			}
		}
	}

	p, err := stream.CloseAndRecv()
	if err != nil {
		return nil, err
	}

	return &model.UpdatePostPayload{Post: p}, nil
}

func (r *mutationResolver) DeletePost(ctx context.Context, input model.DeletePostInput) (*model.DeletePostPayload, error) {
	if _, err := r.postClient.DeletePost(ctx, &pb.DeletePostReq{Id: input.ID}); err != nil {
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

	res, err := r.postClient.GetPost(ctx, &pb.GetPostReq{Id: input.PostID})
	if err != nil {
		return nil, err
	}

	if res.UserId == uID {
		return nil, gqlerr.ForbiddenError("cannot apply your own post")
	}

	a, err := r.postClient.CreateApplyPost(ctx, &pb.CreateApplyPostReq{
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
	res, err := r.postClient.GetApplyPost(ctx, &pb.GetApplyPostReq{Id: input.ID})
	if err != nil {
		return nil, err
	}
	if res.UserId != uID {
		return nil, gqlerr.ForbiddenError("user_id=%d does not have permission to delete apply_post_id=%d: %s", uID, input.ID, err.Error())
	}
	if _, err := r.postClient.DeleteApplyPost(ctx, &pb.DeleteApplyPostReq{Id: input.ID}); err != nil {
		return nil, err
	}
	return &model.DeleteApplyPostPayload{Success: true}, nil
}

func (r *postResolver) ApplyPosts(ctx context.Context, obj *pb.Post) ([]*pb.ApplyPost, error) {
	return dataloader.For(ctx).ApplyPostsByPostIDs.Load(obj.Id)
}

func (r *postResolver) Images(ctx context.Context, obj *pb.Post) ([]*pb.Image, error) {
	res, err := r.imageClient.ListImagesByOwnerID(ctx, &pb.ListImagesByOwnerIDReq{
		OwnerId:   obj.Id,
		OwnerType: pb.OwnerType_POST,
	})
	if err != nil {
		return nil, err
	}

	return res.Images, nil
}

func (r *postResolver) User(ctx context.Context, obj *pb.Post) (*pb.User, error) {
	return r.userClient.GetUser(ctx, &pb.GetUserReq{Id: obj.UserId})
}

func (r *queryResolver) Posts(ctx context.Context, first *int64, after *string, input model.PostsInput) (*model.PostConnection, error) {
	req := &pb.ListPostsReq{Filter: &pb.ListPostsReq_Filter{
		MeetingAtFrom: input.MeetingAtFrom,
		MeetingAtTo:   input.MeetingAtTo,
		FishTypeIds:   input.FishTypeIds,
	}}
	// protobufはnilを受け入れないのでnilチェック
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

func (r *queryResolver) Post(ctx context.Context, id int64) (*pb.Post, error) {
	return r.postClient.GetPost(ctx, &pb.GetPostReq{Id: id})
}

func (r *subscriptionResolver) CreatePostResult(ctx context.Context, input model.CreatePostResultInput) (<-chan *model.CreatePostResultPayload, error) {
	resChan := make(chan *model.CreatePostResultPayload)
	go func() {
		<-ctx.Done()
		close(resChan)
	}()

	_, err := r.natsConn.QueueSubscribe("create.post.result", conf.C.Nats.QueueGroup, func(m *stan.Msg) {

		select {
		case <-ctx.Done():
			if err := m.Sub.Close(); err != nil {
				log.Println(err)
			}
			return
		default:
		}

		e := &pb.Event{}
		if err := protojson.Unmarshal(m.MsgProto.Data, e); err != nil {
			err := err.Error()
			resChan <- &model.CreatePostResultPayload{Error: &err}
			return
		}

		switch e.EventType {
		case "post.approved":

			data := &pb.PostApproved{}

			if err := protojson.Unmarshal(e.EventData, data); err != nil {
				err := err.Error()
				resChan <- &model.CreatePostResultPayload{Error: &err}
				return
			}

			log.Printf("recieved event post.approved input sagaid: %s\nrecieved sagaid: %s\n", input.SagaID, data.SagaId)

			if data.SagaId != input.SagaID {
				log.Println("wrong saga id")
				return
			}

			resChan <- &model.CreatePostResultPayload{Post: data.Post}

		case "post.rejected":

			data := &pb.PostRejected{}

			if err := protojson.Unmarshal(e.EventData, data); err != nil {
				err := err.Error()
				resChan <- &model.CreatePostResultPayload{Error: &err}
				return
			}

			log.Printf("recieved event post.rejected input sagaid: %s\nrecieved sagaid: %s\n", input.SagaID, data.SagaId)

			if data.SagaId != input.SagaID {
				log.Println("wrong saga id")
				return
			}

			resChan <- &model.CreatePostResultPayload{Error: &data.ErrorMessage}
		}

		if err := m.Ack(); err != nil {
			err := err.Error()
			resChan <- &model.CreatePostResultPayload{Error: &err}
		}

	}, stan.SetManualAckMode(), stan.DurableName(conf.C.Nats.QueueGroup))

	if err != nil {
		return nil, err
	}

	return resChan, nil
}

// ApplyPost returns generated.ApplyPostResolver implementation.
func (r *resolver) ApplyPost() generated.ApplyPostResolver { return &applyPostResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *resolver) Post() generated.PostResolver { return &postResolver{r} }

type applyPostResolver struct{ *resolver }
type postResolver struct{ *resolver }
