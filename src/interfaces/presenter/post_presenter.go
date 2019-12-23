package presenter

import (
	"strconv"

	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	"github.com/ezio1119/fishapp-api-gateway/domain/post_grpc"
	"github.com/golang/protobuf/ptypes"
)

type PostPresenter struct{}

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

	return &graphql.Post{
		ID:        id,
		Title:     p.Title,
		Content:   p.Content,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
		UserID:    userID,
	}, nil
}
