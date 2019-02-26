package usecase

import "github.com/applepine1125/learning-cleanarchitecture-in-go/src/pkg/domain"

type UserInteractor struct {
	UserRepo UserRepository
}

func (ui *UserInteractor) Add(u domain.User) error {
	return ui.UserRepo.Store(u)
}

func (ui *UserInteractor) FindById(id int) (domain.User, error) {
	return ui.UserRepo.FindById(id)
}

func (ui *UserInteractor) FindAll() (domain.Users, error) {
	return ui.UserRepo.FindAll()
}
