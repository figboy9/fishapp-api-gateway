package presenter

import (
	"strconv"

	"github.com/ezio1119/fishapp-api-gateway/domain/chat_grpc"
	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	"github.com/ezio1119/fishapp-api-gateway/usecase/presenter"
	"github.com/golang/protobuf/ptypes"
)

type chatPresenter struct{}

func NewChatPresenter() presenter.ChatPresenter {
	return &chatPresenter{}
}

func (*chatPresenter) TransformChatRoomGraphQL(m *chat_grpc.Room) (*graphql.ChatRoom, error) {
	id := strconv.FormatInt(m.Id, 10)
	postID := strconv.FormatInt(m.PostId, 10)
	updatedAt, err := ptypes.Timestamp(m.UpdatedAt)
	if err != nil {
		return nil, err
	}
	createdAt, err := ptypes.Timestamp(m.CreatedAt)
	if err != nil {
		return nil, err
	}
	updatedAt = updatedAt.In(location)
	createdAt = createdAt.In(location)
	return &graphql.ChatRoom{
		ID:        id,
		PostID:    postID,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}, nil
}

func (*chatPresenter) TransformMessageGraphQL(m *chat_grpc.Message) (*graphql.Message, error) {
	id := strconv.FormatInt(m.Id, 10)
	uID := strconv.FormatInt(m.UserId, 10)
	rID := strconv.FormatInt(m.RoomId, 10)
	updatedAt, err := ptypes.Timestamp(m.UpdatedAt)
	if err != nil {
		return nil, err
	}
	createdAt, err := ptypes.Timestamp(m.CreatedAt)
	if err != nil {
		return nil, err
	}
	updatedAt = updatedAt.In(location)
	createdAt = createdAt.In(location)
	return &graphql.Message{
		ID:        id,
		Body:      m.Body,
		RoomID:    rID,
		UserID:    uID,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}, nil
}

func (*chatPresenter) TransformMemberGraphQL(m *chat_grpc.Member) (*graphql.Member, error) {
	id := strconv.FormatInt(m.Id, 10)
	rID := strconv.FormatInt(m.RoomId, 10)
	uID := strconv.FormatInt(m.UserId, 10)
	updatedAt, err := ptypes.Timestamp(m.UpdatedAt)
	if err != nil {
		return nil, err
	}
	createdAt, err := ptypes.Timestamp(m.CreatedAt)
	if err != nil {
		return nil, err
	}
	updatedAt = updatedAt.In(location)
	createdAt = createdAt.In(location)
	return &graphql.Member{
		ID:        id,
		RoomID:    rID,
		UserID:    uID,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}, nil
}
