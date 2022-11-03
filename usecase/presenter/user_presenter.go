package presenter

import "github.com/marcell17002/clean-arch-go/domain/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}
