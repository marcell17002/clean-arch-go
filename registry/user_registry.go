package registry

import (
	"github.com/marcell17002/clean-arch-go/interface/controller"
	ip "github.com/marcell17002/clean-arch-go/interface/presenter"
	ir "github.com/marcell17002/clean-arch-go/interface/repository"
	"github.com/marcell17002/clean-arch-go/usecase/interactor"
	up "github.com/marcell17002/clean-arch-go/usecase/presenter"
	ur "github.com/marcell17002/clean-arch-go/usecase/repository"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}
