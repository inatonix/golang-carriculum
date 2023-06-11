package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	Create(ctx echo.Context) error
	Get(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type userController struct{}

func NewUserController() UserController {
	return &userController{}
}

func (c userController) Create(ctx echo.Context) error {
	// チャンネル作成処理
	return ctx.JSON(http.StatusCreated, "Channel created.")
}

func (c userController) Get(ctx echo.Context) error {
	// チャンネル情報取得処理
	return ctx.JSON(http.StatusOK, "Channel details.")
}

func (c userController) Update(ctx echo.Context) error {
	// チャンネル情報更新処理
	return ctx.JSON(http.StatusOK, "Channel updated.")
}

func (c userController) Delete(ctx echo.Context) error {
	// チャンネル削除処理
	return ctx.JSON(http.StatusOK, "Channel deleted.")
}
