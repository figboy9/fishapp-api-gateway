package repository

import (
	"context"
	"fmt"
	"io"

	"github.com/ezio1119/fishapp-api-gateway/domain/chat_grpc"
	"github.com/ezio1119/fishapp-api-gateway/usecase/repository"
)

type chatRepository struct {
	client chat_grpc.ChatServiceClient
}

func NewChatRepository(c chat_grpc.ChatServiceClient) repository.ChatRepository {
	return &chatRepository{c}
}

func (r *chatRepository) CreateChatRoom(ctx context.Context, req *chat_grpc.CreateRoomReq) (*chat_grpc.Room, error) {
	return r.client.CreateChatRoom(ctx, req)
}

func (r *chatRepository) SendMessage(ctx context.Context, req *chat_grpc.SendMessageReq) (*chat_grpc.Message, error) {
	return r.client.SendMessage(ctx, req)
}

func (r *chatRepository) StreamMessage(ctx context.Context, req *chat_grpc.StreamMessageReq, grpcChan chan *chat_grpc.Message) error {
	fmt.Printf("%#v", req)
	s, err := r.client.StreamMessage(ctx, req)
	if err != nil {
		return err
	}

	for {
		m, err := s.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Println("grpc stream recieved: ", m)
		grpcChan <- m
	}
}

func (r *chatRepository) AddMember(ctx context.Context, req *chat_grpc.AddMemberReq) (*chat_grpc.Member, error) {
	return r.client.AddMember(ctx, req)
}
