package resolver

import (
	"context"
	"log"
	"strconv"

	"github.com/ezio1119/fishapp-api-gateway/domain/chat_grpc"
	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	gen "github.com/ezio1119/fishapp-api-gateway/interfaces/resolver/graphql"
)

// chatInteractorで発生したエラーをクライアントに返すことができない
func (r *subscriptionResolver) MessageAdded(ctx context.Context, roomIds []string) (<-chan *graphql.Message, error) {
	ctx, cancel := context.WithCancel(ctx)
	c, err := getJwtClaimsCtx(ctx)
	if err != nil {
		cancel()
		return nil, err
	}
	rIDs := make([]int64, len(roomIds))
	for i, rID := range roomIds {
		rIDint64, err := strconv.ParseInt(rID, 10, 64)
		if err != nil {
			cancel()
			return nil, err
		}
		rIDs[i] = rIDint64
	}

	msgChan := make(chan *graphql.Message)
	go func() {
		<-ctx.Done()
		close(msgChan)
	}()

	go func() {
		// ブロック処理だが、返り値にチャネルを返さないと行けないため、goroutine使用。ここで発生したエラーをクライアントに返すことができない。
		if err := r.chatInteractor.MessageAdded(ctx, rIDs, c.UserID, msgChan); err != nil {
			log.Fatal(err)
			cancel()
		}
	}()

	return msgChan, nil
}

func (r *mutationResolver) SendMessage(ctx context.Context, in gen.SendMessageInput) (*graphql.Message, error) {
	i, err := strconv.ParseInt(in.RoomID, 10, 64)
	if err != nil {
		return nil, err
	}
	c, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}
	return r.chatInteractor.SendMessage(ctx, &chat_grpc.SendMessageReq{
		Body:   in.Body,
		RoomId: i,
		UserId: c.UserID,
	})
}

func (r *mutationResolver) CreateChatRoom(ctx context.Context, postID string) (*graphql.ChatRoom, error) {
	i, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		return nil, err
	}
	c, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}
	return r.chatInteractor.CreateChatRoom(ctx, &chat_grpc.CreateRoomReq{PostId: i, UserId: c.UserID})
}

func (r *mutationResolver) AddMemberToChatRoom(ctx context.Context, in gen.AddMemberInput) (*graphql.Member, error) {
	rID, err := strconv.ParseInt(in.RoomID, 10, 64)
	if err != nil {
		return nil, err
	}
	c, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}
	return r.chatInteractor.AddMemberToChatRoom(ctx, rID, c.UserID)
}
