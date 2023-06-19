package controller

import (
	"net/http"
	"sluck/usecase"

	"github.com/labstack/echo"
)

type UserController interface {
	Get(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type userController struct {
	u usecase.UserUsecase
}

func NewUserController(u usecase.UserUsecase) UserController {
	return &userController{u}
}

func (c *userController) Get(ctx echo.Context) error {
	id := ctx.Param("id")
	user, err := c.u.GetByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, user)
}

func (c *userController) Create(ctx echo.Context) error {
	var req PostRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	u := post2Model(req)
	c.u.Create(ctx.Request().Context(), u)

	return nil
}

func (c *userController) Update(ctx echo.Context) error {
	var req PutRquest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	u := put2Model(req)
	c.u.Update(ctx.Request().Context(), u)

	return nil
}

func (c *userController) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	c.u.Delete(ctx.Request().Context(), id)

	return nil
}
