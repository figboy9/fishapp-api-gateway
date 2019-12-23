package presenter

import (
	"strconv"
	"time"

	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	"github.com/ezio1119/fishapp-api-gateway/domain/post_grpc"
	"github.com/golang/protobuf/ptypes"
)

type PostPresenter struct{}

var location *time.Location

func (*PostPresenter) TransformPostGraphQL(p *post_grpc.Post) (*graphql.Post, error) {
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
func (p *PostPresenter) TransformListPostGraphQL(listRPC []*post_grpc.Post) ([]*graphql.Post, error) {
	list := make([]*graphql.Post, len(listRPC))
	for i, postRPC := range listRPC {
		post, err := p.TransformPostGraphQL(postRPC)
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
