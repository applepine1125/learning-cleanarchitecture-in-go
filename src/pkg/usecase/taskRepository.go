package usecase

type TaskRepository interface {
	Store(Task) error
	FindById(int) (Task, error)
	FindAll() (Tasks, error)
}
