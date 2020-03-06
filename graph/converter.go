package graph

import (
	"time"

	"github.com/ezio1119/fishapp-api-gateway/graph/model"
	"github.com/ezio1119/fishapp-api-gateway/grpc/auth_grpc"
	"github.com/ezio1119/fishapp-api-gateway/grpc/chat_grpc"
	"github.com/ezio1119/fishapp-api-gateway/grpc/post_grpc"
	"github.com/ezio1119/fishapp-api-gateway/grpc/profile_grpc"
	"github.com/golang/protobuf/ptypes"
)

var loc *time.Location

func init() {
	loc = time.Now().Location()
}

func convertPostGQL(p *post_grpc.Post) (*model.Post, error) {
	updatedAt, err := ptypes.Timestamp(p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	createdAt, err := ptypes.Timestamp(p.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &model.Post{
		ID:        p.Id,
		Title:     p.Title,
		Content:   p.Content,
		UserID:    p.UserId,
		UpdatedAt: updatedAt.In(loc),
		CreatedAt: createdAt.In(loc),
	}, nil
}

func convertPostsGQL(listProto []*post_grpc.Post) ([]*model.Post, error) {
	list := make([]*model.Post, len(listProto))
	for i, postProto := range listProto {
		post, err := convertPostGQL(postProto)
		if err != nil {
			return nil, err
		}
		list[i] = post
	}
	return list, nil
}

func convertApplyPostGQL(a *post_grpc.ApplyPost) (*model.ApplyPost, error) {
	updatedAt, err := ptypes.Timestamp(a.UpdatedAt)
	if err != nil {
		return nil, err
	}
	createdAt, err := ptypes.Timestamp(a.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &model.ApplyPost{
		ID:        a.Id,
		PostID:    a.PostId,
		UserID:    a.UserId,
		UpdatedAt: updatedAt.In(loc),
		CreatedAt: createdAt.In(loc),
	}, nil
}

func convertUserGQL(u *auth_grpc.User) (*model.User, error) {
	updatedAt, err := ptypes.Timestamp(u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	createdAt, err := ptypes.Timestamp(u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:        u.Id,
		Email:     u.Email,
		UpdatedAt: updatedAt.In(loc),
		CreatedAt: createdAt.In(loc),
	}, nil
}

func convertTokenPairGQL(t *auth_grpc.TokenPair) *model.TokenPair {
	return &model.TokenPair{
		IDToken:      t.IdToken,
		RefreshToken: t.RefreshToken,
	}
}

func convertProfileGQL(p *profile_grpc.Profile) (*model.Profile, error) {
	updatedAt, err := ptypes.Timestamp(p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	createdAt, err := ptypes.Timestamp(p.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &model.Profile{
		ID:        p.Id,
		Name:      p.Name,
		UserID:    p.UserId,
		UpdatedAt: updatedAt.In(loc),
		CreatedAt: createdAt.In(loc),
	}, nil
}

func convertChatRoomGQL(r *chat_grpc.Room) (*model.ChatRoom, error) {
	updatedAt, err := ptypes.Timestamp(r.UpdatedAt)
	if err != nil {
		return nil, err
	}
	createdAt, err := ptypes.Timestamp(r.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &model.ChatRoom{
		ID:        r.Id,
		PostID:    r.PostId,
		UpdatedAt: updatedAt.In(loc),
		CreatedAt: createdAt.In(loc),
	}, nil
}

func convertRoomMemberGQL(m *chat_grpc.Member) (*model.RoomMember, error) {
	updatedAt, err := ptypes.Timestamp(m.UpdatedAt)
	if err != nil {
		return nil, err
	}
	createdAt, err := ptypes.Timestamp(m.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &model.RoomMember{
		ID:        m.Id,
		RoomID:    m.RoomId,
		UserID:    m.UserId,
		UpdatedAt: updatedAt.In(loc),
		CreatedAt: createdAt.In(loc),
	}, nil
}

func convertRoomMessageGQL(m *chat_grpc.Message) (*model.RoomMessage, error) {
	updatedAt, err := ptypes.Timestamp(m.UpdatedAt)
	if err != nil {
		return nil, err
	}
	createdAt, err := ptypes.Timestamp(m.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &model.RoomMessage{
		ID:        m.Id,
		Body:      m.Body,
		RoomID:    m.RoomId,
		UserID:    m.UserId,
		UpdatedAt: updatedAt.In(loc),
		CreatedAt: createdAt.In(loc),
	}, nil
}
