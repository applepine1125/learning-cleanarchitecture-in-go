package controller

import (
	"fmt"
	"strconv"

	"github.com/applepine1125/learning-cleanarchitecture-in-go/src/pkg/domain"
	"github.com/applepine1125/learning-cleanarchitecture-in-go/src/pkg/interfaces/database"
	"github.com/applepine1125/learning-cleanarchitecture-in-go/src/pkg/usecase"
)

type TaskController struct {
	Interactor usecase.TaskInteractor
}

func NewTaskController(sqlHandler database.SqlHandler) *TaskController {
	return &TaskController{
		Interactor: usecase.TaskInteractor{
			UserRepo: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
			TaskRepo: &database.TaskRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *TaskController) Create(c Context) {
	u := domain.Task{}
	c.Bind(&u)
	err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, fmt.Errorf(err.Error()))
		return
	}
	c.JSON(201, "")
}

func (controller *TaskController) GetTask(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := controller.Interactor.FindById(id)
	if task.IsEmpty() {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, task)
}

func (controller *TaskController) GetTasks(c Context) {
	tasks, err := controller.Interactor.FindAll()
	if err != nil {
		c.JSON(500, fmt.Errorf(err.Error()))
		return
	}
	c.JSON(200, tasks)
}
