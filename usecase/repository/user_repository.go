package repository

import "github.com/marcell17002/clean-arch-go/domain/model"

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
}
