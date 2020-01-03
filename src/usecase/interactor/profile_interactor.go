package interactor

import (
	"time"

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

type ProfileInteractor interface{}
