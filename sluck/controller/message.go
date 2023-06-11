package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type MessageController interface {
	Create(ctx echo.Context) error
	Get(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type messageController struct{}

func NewMessageController() MessageController {
	return &messageController{}
}

func (c messageController) Create(ctx echo.Context) error {
	// チャンネル作成処理
	return ctx.JSON(http.StatusCreated, "Channel created.")
}

func (c messageController) Get(ctx echo.Context) error {
	// チャンネル情報取得処理
	return ctx.JSON(http.StatusOK, "Channel details.")
}

func (c messageController) Update(ctx echo.Context) error {
	// チャンネル情報更新処理
	return ctx.JSON(http.StatusOK, "Channel updated.")
}

func (c messageController) Delete(ctx echo.Context) error {
	// チャンネル削除処理
	return ctx.JSON(http.StatusOK, "Channel deleted.")
}
