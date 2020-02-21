package presenter

import (
	"github.com/ezio1119/fishapp-api-gateway/domain/chat_grpc"
	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
)

type ChatPresenter interface {
	TransformChatRoomGraphQL(r *chat_grpc.Room) (*graphql.ChatRoom, error)
	TransformMessageGraphQL(m *chat_grpc.Message) (*graphql.Message, error)
	TransformMemberGraphQL(m *chat_grpc.Member) (*graphql.Member, error)
}
