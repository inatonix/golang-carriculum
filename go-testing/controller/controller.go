package controller

import (
	"go-testing/usecase"

	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type TaskController interface {
	Get(c echo.Context) error
}

type taskController struct {
	u usecase.TaskUsecase
}

func NewTaskController(u usecase.TaskUsecase) TaskController {
	return &taskController{u: u}
}

func (t *taskController) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		msg := fmt.Errorf("parse error: %v", err.Error())
		return c.JSON(http.StatusBadRequest, msg.Error())
	}
	task, err := t.u.GetTask(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, task)
}
