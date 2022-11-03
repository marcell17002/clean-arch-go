package presenters

import (
	"github.com/marcell17002/clean-arch-go/domain/model"
	"github.com/marcell17002/clean-arch-go/usecase/presenter"
)

type userPresenter struct{}

func NewUserPresenter() presenter.UserPresenter {
	return &userPresenter{}
}

func (up *userPresenter) ResponseUsers(us []*model.User) []*model.User {
	for _, u := range us {
		u.Name = "Mr." + u.Name
	}
	return us
}
