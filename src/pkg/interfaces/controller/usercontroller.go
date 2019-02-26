package controller

import (
	"fmt"
	"strconv"

	"github.com/applepine1125/learning-cleanarchitecture-in-go/src/pkg/domain"
	"github.com/applepine1125/learning-cleanarchitecture-in-go/src/pkg/interfaces/database"
	"github.com/applepine1125/learning-cleanarchitecture-in-go/src/pkg/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepo: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c Context) {
	u := domain.User{}
	c.Bind(&u)
	err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, fmt.Errorf(err.Error()))
		return
	}
	c.JSON(201, "")
}

func (controller *UserController) GetUser(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := controller.Interactor.FindById(id)
	if user.IsEmpty() {
		c.JSON(500, fmt.Errorf("user not found"))
		return
	}
	c.JSON(200, user)
}

func (controller *UserController) GetUsers(c Context) {
	users, err := controller.Interactor.FindAll()
	if err != nil {
		c.JSON(500, fmt.Errorf(err.Error()))
		return
	}
	c.JSON(200, users)
}
