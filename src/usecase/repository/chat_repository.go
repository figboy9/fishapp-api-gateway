package repository

import (
	"context"

	"github.com/ezio1119/fishapp-api-gateway/domain/chat_grpc"
)

type ChatRepository interface {
	CreateChatRoom(ctx context.Context, req *chat_grpc.CreateRoomReq) (*chat_grpc.Room, error)
	SendMessage(ctx context.Context, req *chat_grpc.SendMessageReq) (*chat_grpc.Message, error)
	StreamMessage(ctx context.Context, req *chat_grpc.StreamMessageReq, grpcChan chan *chat_grpc.Message) error
	AddMember(ctx context.Context, req *chat_grpc.AddMemberReq) (*chat_grpc.Member, error)
}
