package presenter

import (
	"github.com/ezio1119/fishapp-api-gateway/usecase/presenter"
)

type profilePresenter struct{}

func NewProfilePresenter() presenter.ProfilePresenter {
	return &profilePresenter{}
}
