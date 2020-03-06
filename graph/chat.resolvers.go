// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package graph

import (
	"context"

	"github.com/ezio1119/fishapp-api-gateway/graph/generated"
	"github.com/ezio1119/fishapp-api-gateway/graph/model"
	"github.com/ezio1119/fishapp-api-gateway/grpc/chat_grpc"
)

func (r *mutationResolver) CreateChatRoom(ctx context.Context, input model.CreateChatRoomInput) (*model.CreateChatRoomPayload, error) {
	c, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}
	res, err := r.chatClient.CreateChatRoom(ctx, &chat_grpc.CreateRoomReq{PostId: input.PostID, UserId: c.UserID})
	if err != nil {
		return nil, err
	}
	room, err := convertChatRoomGQL(res)
	if err != nil {
		return nil, err
	}
	return &model.CreateChatRoomPayload{Room: room}, nil
}

func (r *mutationResolver) AddMemberChatRoom(ctx context.Context, input model.AddMemberChatRoomInput) (*model.AddMemberChatRoomPayload, error) {
	c, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}
	res, err := r.chatClient.AddMember(ctx, &chat_grpc.AddMemberReq{RoomId: input.RoomID, UserId: c.UserID})
	if err != nil {
		return nil, err
	}
	m, err := convertRoomMemberGQL(res)
	if err != nil {
		return nil, err
	}
	return &model.AddMemberChatRoomPayload{Member: m}, nil
}

func (r *mutationResolver) SendMessageChatRoom(ctx context.Context, input model.SendMessageChatRoomInput) (*model.SendMessageChatRoomPayload, error) {
	c, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}
	res, err := r.chatClient.SendMessage(ctx, &chat_grpc.SendMessageReq{
		Body:   input.Body,
		RoomId: input.RoomID,
		UserId: c.UserID,
	})
	if err != nil {
		return nil, err
	}

	m, err := convertRoomMessageGQL(res)
	if err != nil {
		return nil, err
	}
	return &model.SendMessageChatRoomPayload{Message: m}, nil
}

func (r *subscriptionResolver) MessageAdded(ctx context.Context, input model.MessageAddedInput) (<-chan *model.MessageAddedPayload, error) {
	panic("not implemented")

	// 	c, err := getJwtClaimsCtx(ctx)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	gqlChan := make(chan *model.MessageAddedPayload)
	// 	log.Println("1")

	// 	go func(ch chan *model.MessageAddedPayload) {
	// 		fmt.Println(input.RoomIds, c.UserID)
	// 		s, err := r.chatClient.StreamMessage(ctx, &chat_grpc.StreamMessageReq{
	// 			RoomIds: input.RoomIds,
	// 			UserId:  c.UserID,
	// 		})
	// 		fmt.Println("error: ", err)
	// 		if err != nil {
	// 			log.Println(err)
	// 			return
	// 		}

	// 		for {
	// 			select {
	// 			case _ = <-ctx.Done():
	// 				fmt.Println("ctx is done ")
	// 				return
	// 			default:

	// 				log.Println("NumGoroutine: ", runtime.NumGoroutine())
	// 				mP, err := s.Recv()
	// 				if err == io.EOF {
	// 					return
	// 				}
	// 				if err != nil {
	// 					log.Println(err)
	// 					graphql.AddError(ctx, err)
	// 					return
	// 				}
	// 				fmt.Println("grpc stream recieved: ", mP)
	// 				m, err := convertRoomMessageGQL(mP)
	// 				if err != nil {
	// 					log.Println(err)
	// 					graphql.AddError(ctx, err)
	// 					return
	// 				}
	// 				ch <- &model.MessageAddedPayload{Message: m}
	// 			}
	// 		}
	// 	}(gqlChan)

	// 	return gqlChan, nil
}

func (r *resolver) Mutation() generated.MutationResolver         { return &mutationResolver{r} }
func (r *resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *resolver }
type subscriptionResolver struct{ *resolver }
