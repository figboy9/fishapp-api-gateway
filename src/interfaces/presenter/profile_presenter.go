package presenter

import (
	"strconv"

	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	"github.com/ezio1119/fishapp-api-gateway/domain/profile_grpc"
	"github.com/ezio1119/fishapp-api-gateway/usecase/presenter"
	"github.com/golang/protobuf/ptypes"
)

type profilePresenter struct{}

func NewProfilePresenter() presenter.ProfilePresenter {
	return &profilePresenter{}
}

func (*profilePresenter) TransformProfileGraphQL(p *profile_grpc.Profile) (*graphql.Profile, error) {
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
	return &graphql.Profile{
		ID:        id,
		Name:      p.Name,
		UserID:    userID,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}, nil
}
