package controller

import (
	"fmt"
	"sluck/usecase"

	"github.com/labstack/echo"
)

type UserController interface {
	Create(ctx echo.Context) error
}

type userController struct {
	u usecase.UserUsecase
}

func NewUserController(u usecase.UserUsecase) UserController {
	return &userController{u}
}

func (c *userController) Create(ctx echo.Context) error {
	fmt.Println("creating...")
	return nil
}
