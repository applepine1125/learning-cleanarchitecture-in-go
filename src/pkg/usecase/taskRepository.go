package usecase

import "github.com/applepine1125/learning-cleanarchitecture-in-go/src/pkg/domain"

type TaskRepository interface {
	Store(domain.Task) error
	FindById(int) (domain.Task, error)
	FindAll() (domain.Tasks, error)
}
