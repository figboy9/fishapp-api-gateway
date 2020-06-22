package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"io"
	"log"
	"runtime"
	"strconv"

	"github.com/ezio1119/fishapp-api-gateway/conf"
	"github.com/ezio1119/fishapp-api-gateway/graph/generated"
	"github.com/ezio1119/fishapp-api-gateway/graph/gqlerr"
	"github.com/ezio1119/fishapp-api-gateway/graph/model"
	"github.com/ezio1119/fishapp-api-gateway/pb"
)

func (r *memberResolver) User(ctx context.Context, obj *pb.Member) (*pb.User, error) {
	return r.userClient.GetUser(ctx, &pb.GetUserReq{Id: obj.UserId})
}

func (r *messageResolver) User(ctx context.Context, obj *pb.Message) (*pb.User, error) {
	return r.userClient.GetUser(ctx, &pb.GetUserReq{Id: obj.UserId})
}

func (r *messageResolver) Image(ctx context.Context, obj *pb.Message) (*pb.Image, error) {
	res, err := r.imageClient.ListImagesByOwnerID(ctx, &pb.ListImagesByOwnerIDReq{
		OwnerId:   obj.Id,
		OwnerType: pb.OwnerType_MESSAGE,
	})
	if err != nil {
		return nil, err
	}

	if len(res.Images) == 0 {
		return nil, err
	}

	return res.Images[0], nil
}

func (r *mutationResolver) CreateMessage(ctx context.Context, input model.CreateMessageInput) (*model.CreateMessagePayload, error) {
	if input.Image != nil && input.Body != nil || input.Image == nil && input.Body == nil {
		return nil, gqlerr.UserInputError("invalid CreateMessageInput.Image, CreateMessageInput.Body: value must be set either image or body")
	}

	c, err := getClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	uID, err := strconv.ParseInt(c.User.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	var body string
	if input.Body != nil {
		body = *input.Body
	}

	req := &pb.CreateMessageReq{
		Data: &pb.CreateMessageReq_Info{
			Info: &pb.CreateMessageReqInfo{
				Body:   body,
				RoomId: input.RoomID,
				UserId: uID,
			},
		},
	}

	stream, err := r.chatClient.CreateMessage(ctx)
	if err != nil {
		return nil, err
	}

	if err := stream.Send(req); err != nil {
		return nil, err
	}

	if input.Image != nil {
		for {
			buf := make([]byte, conf.C.Sv.ChunkDataSize)
			n, err := input.Image.File.Read(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, gqlerr.InternalServerError("cannot read chunk to buffe: %s", err)
			}

			req = &pb.CreateMessageReq{
				Data: &pb.CreateMessageReq_ImageChunk{
					ImageChunk: buf[:n],
				},
			}

			if stream.Send(req); err != nil {
				return nil, gqlerr.InternalServerError("cannot read chunk to buffe: %s", err)
			}
		}
	}

	m, err := stream.CloseAndRecv()
	if err != nil {
		return nil, err
	}

	return &model.CreateMessagePayload{Message: m}, nil
}

func (r *queryResolver) Room(ctx context.Context, postID int64) (*pb.Room, error) {
	return r.chatClient.GetRoom(ctx, &pb.GetRoomReq{GetRoom: &pb.GetRoomReq_PostId{PostId: postID}})
}

func (r *roomResolver) Post(ctx context.Context, obj *pb.Room) (*pb.Post, error) {
	return r.postClient.GetPost(ctx, &pb.GetPostReq{Id: obj.PostId})
}

func (r *subscriptionResolver) MessageAdded(ctx context.Context, input model.MessageAddedInput) (<-chan *model.MessageAddedPayload, error) {
	fmt.Println(runtime.NumGoroutine())
	gqlChan := make(chan *model.MessageAddedPayload)
	// go func() {
	// 	<-ctx.Done()
	// 	close(gqlChan)
	// }()

	go func() {
		stream, err := r.chatClient.StreamMessage(ctx, &pb.StreamMessageReq{RoomId: input.RoomID})
		if err != nil {
			log.Println(err)
			return
		}

		for {
			m, err := stream.Recv()
			if err == io.EOF {
				return
			}

			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("msg: %#v\n", m)

			gqlChan <- &model.MessageAddedPayload{Message: m}
		}

	}()

	return gqlChan, nil
}

// Member returns generated.MemberResolver implementation.
func (r *resolver) Member() generated.MemberResolver { return &memberResolver{r} }

// Message returns generated.MessageResolver implementation.
func (r *resolver) Message() generated.MessageResolver { return &messageResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Room returns generated.RoomResolver implementation.
func (r *resolver) Room() generated.RoomResolver { return &roomResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type memberResolver struct{ *resolver }
type messageResolver struct{ *resolver }
type mutationResolver struct{ *resolver }
type queryResolver struct{ *resolver }
type roomResolver struct{ *resolver }
type subscriptionResolver struct{ *resolver }
