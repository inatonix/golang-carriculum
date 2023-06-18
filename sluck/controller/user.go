package controller

import (
	"net/http"

	"sluck/usecase"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	Create(ctx echo.Context) error
	Get(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type userController struct {
	u usecase.UserUsecase
}

func NewUserController() UserController {
	return &userController{}
}

func (c userController) Create(ctx echo.Context) error {
	var req UserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	u := ToModel(&req)
	err := c.u.Create(ctx.Request().Context(), &u)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, "user created.")
}

func (c userController) Get(ctx echo.Context) error {
	// チャンネル情報取得処理
	return ctx.JSON(http.StatusOK, "user details.")
}

func (c userController) Update(ctx echo.Context) error {
	// チャンネル情報更新処理
	return ctx.JSON(http.StatusOK, "user updated.")
}

func (c userController) Delete(ctx echo.Context) error {
	// チャンネル削除処理
	return ctx.JSON(http.StatusOK, "user deleted.")
}
