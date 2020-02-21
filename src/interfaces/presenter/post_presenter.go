package presenter

import (
	"strconv"
	"time"

	"github.com/ezio1119/fishapp-api-gateway/domain/entry_post_grpc"
	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	"github.com/ezio1119/fishapp-api-gateway/domain/post_grpc"
	"github.com/ezio1119/fishapp-api-gateway/usecase/presenter"
	"github.com/golang/protobuf/ptypes"
)

type postPresenter struct{}

func NewPostPresenter() presenter.PostPresenter {
	return &postPresenter{}
}

var location *time.Location

func (*postPresenter) TransformPostGraphQL(p *post_grpc.Post) (*graphql.Post, error) {
	id := strconv.FormatInt(p.Id, 10)
	userID := strconv.FormatInt(p.UserId, 10)
	updatedAt, err := ptypes.Timestamp(p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	createdAt, err := ptypes.Timestamp(p.CreatedAt)
	if err != nil {
		return nil, err
	}
	updatedAt = updatedAt.In(location)
	createdAt = createdAt.In(location)
	return &graphql.Post{
		ID:        id,
		Title:     p.Title,
		Content:   p.Content,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
		UserID:    userID,
	}, nil
}
func (p *postPresenter) TransformListPostGraphQL(listProto []*post_grpc.Post) ([]*graphql.Post, error) {
	list := make([]*graphql.Post, len(listProto))
	for i, postProto := range listProto {
		post, err := p.TransformPostGraphQL(postProto)
		if err != nil {
			return nil, err
		}
		list[i] = post
	}
	return list, nil
}

func init() {
	location = time.Now().Location()
}

func (p *postPresenter) TransformEntryPostGraphQL(e *entry_post_grpc.Entry) (*graphql.EntryPost, error) {
	id := strconv.FormatInt(e.Id, 10)
	userID := strconv.FormatInt(e.UserId, 10)
	postID := strconv.FormatInt(e.PostId, 10)
	updatedAt, err := ptypes.Timestamp(e.UpdatedAt)
	if err != nil {
		return nil, err
	}
	createdAt, err := ptypes.Timestamp(e.CreatedAt)
	if err != nil {
		return nil, err
	}
	updatedAt = updatedAt.In(location)
	createdAt = createdAt.In(location)
	return &graphql.EntryPost{
		ID:        id,
		UserID:    userID,
		PostID:    postID,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}, nil
}

func (p *postPresenter) TransformEntriesGraphQL(entriesProto []*entry_post_grpc.Entry) ([]*graphql.EntryPost, error) {
	entries := make([]*graphql.EntryPost, len(entriesProto))
	for i, entryProto := range entriesProto {
		entry, err := p.TransformEntryPostGraphQL(entryProto)
		if err != nil {
			return nil, err
		}
		entries[i] = entry
	}
	return entries, nil
}
