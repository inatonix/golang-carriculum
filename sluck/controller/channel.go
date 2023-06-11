package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ChannelController interface {
	CreateChannel(ctx echo.Context) error
	GetChannel(ctx echo.Context) error
	UpdateChannel(ctx echo.Context) error
	DeleteChannel(ctx echo.Context) error
}

type channelController struct{}

func NewChannelController() ChannelController {
	return &channelController{}
}

func (c channelController) CreateChannel(ctx echo.Context) error {
	// チャンネル作成処理
	return ctx.JSON(http.StatusCreated, "Channel created.")
}

func (c channelController) GetChannel(ctx echo.Context) error {
	// チャンネル情報取得処理
	return ctx.JSON(http.StatusOK, "Channel details.")
}

func (c channelController) UpdateChannel(ctx echo.Context) error {
	// チャンネル情報更新処理
	return ctx.JSON(http.StatusOK, "Channel updated.")
}

func (c channelController) DeleteChannel(ctx echo.Context) error {
	// チャンネル削除処理
	return ctx.JSON(http.StatusOK, "Channel deleted.")
}
