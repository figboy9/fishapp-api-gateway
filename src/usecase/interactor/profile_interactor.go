package interactor

import (
	"context"
	"time"

	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	"github.com/ezio1119/fishapp-api-gateway/domain/profile_grpc"
	"github.com/ezio1119/fishapp-api-gateway/usecase/presenter"
	"github.com/ezio1119/fishapp-api-gateway/usecase/repository"
)

type profileInteractor struct {
	profileRepository repository.ProfileRepository
	profilePresenter  presenter.ProfilePresenter
	ctxTimeout        time.Duration
}

func NewProfileInteractor(r repository.ProfileRepository, p presenter.ProfilePresenter, t time.Duration) ProfileInteractor {
	return &profileInteractor{r, p, t}
}

type ProfileInteractor interface {
	Profile(ctx context.Context, userID *profile_grpc.ID) (*graphql.Profile, error)
	CreateProfile(ctx context.Context, req *profile_grpc.CreateReq) (*graphql.Profile, error)
	UpdateProfile(ctx context.Context, req *profile_grpc.UpdateReq) (*graphql.Profile, error)
	DeleteProfile(ctx context.Context, userID *profile_grpc.ID) (bool, error)
}

func (i *profileInteractor) Profile(ctx context.Context, userID *profile_grpc.ID) (*graphql.Profile, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	profileProto, err := i.profileRepository.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return i.profilePresenter.TransformProfileGraphQL(profileProto)
}

func (i *profileInteractor) CreateProfile(ctx context.Context, req *profile_grpc.CreateReq) (*graphql.Profile, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	profileProto, err := i.profileRepository.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.profilePresenter.TransformProfileGraphQL(profileProto)
}

func (i *profileInteractor) UpdateProfile(ctx context.Context, req *profile_grpc.UpdateReq) (*graphql.Profile, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	profileProto, err := i.profileRepository.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.profilePresenter.TransformProfileGraphQL(profileProto)
}

func (i *profileInteractor) DeleteProfile(ctx context.Context, userID *profile_grpc.ID) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	success, err := i.profileRepository.Delete(ctx, userID)
	if err != nil {
		return false, err
	}
	return success.Value, nil
}
