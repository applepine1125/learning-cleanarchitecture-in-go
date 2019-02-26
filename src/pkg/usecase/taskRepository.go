package usecase

type TaskRepository interface {
	Store(Task) error
	FindById(int) Task
	FindAll() (Tasks, error)
}
