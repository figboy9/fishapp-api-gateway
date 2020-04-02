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

	"github.com/ezio1119/fishapp-api-gateway/graph/dataloader"
	"github.com/ezio1119/fishapp-api-gateway/graph/generated"
	"github.com/ezio1119/fishapp-api-gateway/graph/model"
	"github.com/ezio1119/fishapp-api-gateway/grpc/chat_grpc"
	"github.com/ezio1119/fishapp-api-gateway/grpc/post_grpc"
	"github.com/ezio1119/fishapp-api-gateway/grpc/profile_grpc"
)

func (r *memberResolver) Profile(ctx context.Context, obj *chat_grpc.Member) (*profile_grpc.Profile, error) {
	p, err := dataloader.For(ctx).ProfileByUserID.Load(obj.UserId)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *messageResolver) Profile(ctx context.Context, obj *chat_grpc.Message) (*profile_grpc.Profile, error) {
	p, err := dataloader.For(ctx).ProfileByUserID.Load(obj.UserId)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *mutationResolver) CreateRoom(ctx context.Context, input model.CreateRoomInput) (*model.CreateRoomPayload, error) {
	c, err := getClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	uID, err := strconv.ParseInt(c.User.ID, 10, 64)
	if err != nil {
		return nil, err
	}
	res, err := r.chatClient.CreateRoom(ctx, &chat_grpc.CreateRoomReq{
		PostId: input.PostID,
		UserId: uID,
	})
	if err != nil {
		return nil, err
	}
	return &model.CreateRoomPayload{Room: res}, nil
}

func (r *mutationResolver) CreateMember(ctx context.Context, input model.CreateMemberInput) (*model.CreateMemberPayload, error) {
	c, err := getClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	uID, err := strconv.ParseInt(c.User.ID, 10, 64)
	if err != nil {
		return nil, err
	}
	m, err := r.chatClient.CreateMember(ctx, &chat_grpc.CreateMemberReq{
		RoomId: input.RoomID,
		UserId: uID,
	})
	if err != nil {
		return nil, err
	}
	return &model.CreateMemberPayload{Member: m}, nil
}

func (r *mutationResolver) DeleteMember(ctx context.Context, input model.DeleteMemberInput) (*model.DeleteMemberPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateMessage(ctx context.Context, input model.CreateMessageInput) (*model.CreateMessagePayload, error) {
	c, err := getClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	uID, err := strconv.ParseInt(c.User.ID, 10, 64)
	if err != nil {
		return nil, err
	}
	m, err := r.chatClient.CreateMessage(ctx, &chat_grpc.CreateMessageReq{
		Body:   input.Body,
		RoomId: input.RoomID,
		UserId: uID,
	})
	if err != nil {
		return nil, err
	}
	return &model.CreateMessagePayload{Message: m}, nil
}

func (r *queryResolver) Room(ctx context.Context, postID int64) (*chat_grpc.Room, error) {
	c, err := getClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	uID, err := strconv.ParseInt(c.User.ID, 10, 64)
	if err != nil {
		return nil, err
	}
	return r.chatClient.GetRoom(ctx, &chat_grpc.GetRoomReq{PostId: postID, UserId: uID})
}

func (r *roomResolver) Post(ctx context.Context, obj *chat_grpc.Room) (*post_grpc.Post, error) {
	return r.postClient.GetPost(ctx, &post_grpc.GetPostReq{Id: obj.PostId})
}

func (r *roomResolver) Messages(ctx context.Context, obj *chat_grpc.Room) ([]*chat_grpc.Message, error) {
	c, err := getClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	uID, err := strconv.ParseInt(c.User.ID, 10, 64)
	if err != nil {
		return nil, err
	}
	res, err := r.chatClient.ListMessages(ctx, &chat_grpc.ListMessagesReq{RoomId: obj.Id, UserId: uID})
	if err != nil {
		return nil, err
	}
	return res.Messages, nil
}

func (r *roomResolver) Members(ctx context.Context, obj *chat_grpc.Room) ([]*chat_grpc.Member, error) {
	c, err := getClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	uID, err := strconv.ParseInt(c.User.ID, 10, 64)
	if err != nil {
		return nil, err
	}
	res, err := r.chatClient.ListMembers(ctx, &chat_grpc.ListMembersReq{RoomId: obj.Id, UserId: uID})
	if err != nil {
		return nil, err
	}
	return res.Members, nil
}

func (r *subscriptionResolver) MessageAdded(ctx context.Context, input model.MessageAddedInput) (<-chan *model.MessageAddedPayload, error) {
	fmt.Println(runtime.NumGoroutine())
	c, err := getClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	uID, err := strconv.ParseInt(c.User.ID, 10, 64)
	if err != nil {
		return nil, err
	}
	if _, err := r.chatClient.GetMember(ctx, &chat_grpc.GetMemberReq{RoomId: input.RoomID, UserId: uID}); err != nil {
		return nil, err
	}
	gqlChan := make(chan *model.MessageAddedPayload)
	go func() {
		<-ctx.Done()
		close(gqlChan)
	}()
	go func() {
		stream, err := r.chatClient.StreamMessage(ctx, &chat_grpc.StreamMessageReq{RoomId: input.RoomID, UserId: uID})
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
				log.Println(err)
				return
			}
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
