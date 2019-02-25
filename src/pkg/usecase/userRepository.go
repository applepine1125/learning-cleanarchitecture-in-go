package usescase

import "github.com/applepine1125/learning-cleanarchitecture-in-go/src/pkg/domain"

type UserRepository interface {
	Store(domain.User) error
	FindById(int) (domain.User, error)
	FindAll() (domain.Users, error)
}
