package usecase

import "github.com/applepine1125/learning-cleanarchitecture-in-go/src/pkg/domain"

type UserRepository interface {
	Store(domain.User) error
	FindById(int) domain.User
	FindAll() (domain.Users, error)
}
