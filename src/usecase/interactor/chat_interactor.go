package interactor

import (
	"context"
	"fmt"
	"time"

	"github.com/ezio1119/fishapp-api-gateway/domain/chat_grpc"
	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	"github.com/ezio1119/fishapp-api-gateway/usecase/presenter"
	"github.com/ezio1119/fishapp-api-gateway/usecase/repository"
	"golang.org/x/sync/errgroup"
)

type chatInteractor struct {
	chatRepository repository.ChatRepository
	chatPresenter  presenter.ChatPresenter
	ctxTimeout     time.Duration
}

func NewChatInteractor(cr repository.ChatRepository, cp presenter.ChatPresenter, t time.Duration) ChatInteractor {
	return &chatInteractor{cr, cp, t}
}

type ChatInteractor interface {
	CreateChatRoom(ctx context.Context, req *chat_grpc.CreateRoomReq) (*graphql.ChatRoom, error)
	AddMemberToChatRoom(ctx context.Context, roomID int64, userID int64) (*graphql.Member, error)
	SendMessage(ctx context.Context, req *chat_grpc.SendMessageReq) (*graphql.Message, error)
	MessageAdded(ctx context.Context, roomIDs []int64, userID int64, gqlChan chan *graphql.Message) error
}

func (i *chatInteractor) CreateChatRoom(ctx context.Context, req *chat_grpc.CreateRoomReq) (*graphql.ChatRoom, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	roomProto, err := i.chatRepository.CreateChatRoom(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.chatPresenter.TransformChatRoomGraphQL(roomProto)
}

func (i *chatInteractor) SendMessage(ctx context.Context, req *chat_grpc.SendMessageReq) (*graphql.Message, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	m, err := i.chatRepository.SendMessage(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.chatPresenter.TransformMessageGraphQL(m)
}

func (i *chatInteractor) AddMemberToChatRoom(ctx context.Context, rID int64, uID int64) (*graphql.Member, error) {
	m, err := i.chatRepository.AddMember(ctx, &chat_grpc.AddMemberReq{
		RoomId: rID,
		UserId: uID,
	})
	if err != nil {
		return nil, err
	}
	return i.chatPresenter.TransformMemberGraphQL(m)
}

func (i *chatInteractor) MessageAdded(ctx context.Context, rIDs []int64, uID int64, gqlChan chan *graphql.Message) error {
	grpcChan := make(chan *chat_grpc.Message)
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		<-ctx.Done()
		close(grpcChan)
		return nil
	})

	eg.Go(func() error {
		for mP := range grpcChan {
			fmt.Println("cascavassa", mP)
			m, err := i.chatPresenter.TransformMessageGraphQL(mP)
			if err != nil {
				return err
			}
			gqlChan <- m
		}
		return nil
	})

	eg.Go(func() error {
		if err := i.chatRepository.StreamMessage(ctx, &chat_grpc.StreamMessageReq{
			RoomIds: rIDs,
			UserId:  uID,
		}, grpcChan); err != nil {
			return err
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		return err
	}

	// err := i.chatRepository.StreamMessage(ctx, &chat_grpc.StreamMessageReq{
	// 	RoomIds: rIDs,
	// 	UserId:  uID,
	// }, grpcChan)
	// if err != nil {
	// 	return err
	// }

	// go func() {
	// 	for {
	// 		mP, err := s.Recv()
	// 		fmt.Println(err)
	// 		if err == io.EOF {
	// 			break
	// 		}
	// 		// if err != nil {
	// 		// 	return err
	// 		// }
	// 		fmt.Println(mP)
	// 		m, _ := i.chatPresenter.TransformMessageGraphQL(mP)
	// 		// if err != nil {
	// 		// 	return err
	// 		// }
	// 		gqlChan <- m
	// 	}

	// }()

	return nil
}
