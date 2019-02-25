package usecase

import (
	"fmt"

	"github.com/applepine1125/learning-cleanarchitecture-in-go/src/pkg/domain"
)

type TaskInteractor struct {
	userRepo UserRepository
	taskRepo TaskRepository
}

func (ti *TaskInteractor) Add(t domain.Task) error {
	return ti.taskRepo.Store(t)
}

func (ti *TaskInteractor) FindById(id int) (domain.Task, error) {
	task, err := ti.taskRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	user, err := ti.userRepo.FindById(task.UserID)
	if err != nil {
		return nil, err
	}

	task.UserFullName = user.GetFullName()

	return task, nil
}

func (ti *TaskInteractor) FindAll() (domain.Tasks, error) {
	tasks, err := ti.taskRepo.FindAll()
	if err != nil {
		return nil, err
	}

	users, err := ti.userRepo.FindAll()
	if err != nil {
		return nil, err
	}

	for _, task := range tasks {
		user := users.FindById(task.UserID)
		if user.IsEmpty() {
			return nil, fmt.Errorf("user(ID:%d) not found", task.UserID)
		}

		task.UserFullName = user.getFullName()
	}

	return tasks, nil
}
