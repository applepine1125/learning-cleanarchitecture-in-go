package usecase

import (
	"fmt"

	"github.com/applepine1125/learning-cleanarchitecture-in-go/src/pkg/domain"
)

type TaskInteractor struct {
	UserRepo UserRepository
	TaskRepo TaskRepository
}

func (ti *TaskInteractor) Add(dt domain.Task) error {
	task := Task{dt.ID, dt.UserID, dt.Title, dt.Description}
	return ti.TaskRepo.Store(task)
}

func (ti *TaskInteractor) FindById(id int) (domain.Task, error) {
	task := ti.TaskRepo.FindById(id)
	if task.IsEmpty() {
		return domain.Task{}, fmt.Errorf("task not found")
	}

	user := ti.UserRepo.FindById(task.UserID)
	if user.IsEmpty() {
		return domain.Task{}, fmt.Errorf("user not found")
	}

	domainTask := convertToDomainTask(task, user.GetFullName())

	return domainTask, nil
}

func (ti *TaskInteractor) FindAll() (domain.Tasks, error) {
	tasks, err := ti.TaskRepo.FindAll()
	if err != nil {
		return nil, err
	}

	users, err := ti.UserRepo.FindAll()
	if err != nil {
		return nil, err
	}

	domainTasks := domain.Tasks{}
	for _, task := range tasks {
		user := users.FindById(task.UserID)
		if user.IsEmpty() {
			return nil, fmt.Errorf("user(ID:%d) not found", task.UserID)
		}

		domainTask := convertToDomainTask(task, user.GetFullName())
		domainTasks = append(domainTasks, domainTask)
	}

	return domainTasks, nil
}

func convertToDomainTask(task Task, fullName string) domain.Task {
	return domain.Task{
		ID:           task.ID,
		UserID:       task.UserID,
		UserFullName: fullName,
		Title:        task.Title,
		Description:  task.Description,
	}

}
