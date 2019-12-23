package interactor

import (
	"time"

	"github.com/ezio1119/fishapp-api-gateway/usecase/presenter"
	"github.com/ezio1119/fishapp-api-gateway/usecase/repository"
)

type UserInteractor struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
	ContextTimeout time.Duration
}

type UUserInteractor interface {}