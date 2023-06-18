package main

import (
	"sluck/controller"
	"sluck/repository"
	"sluck/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// チャンネルエンドポイント
	cc := controller.NewChannelController()
	e.POST("/channels", cc.CreateChannel)
	e.GET("/channels/:id", cc.GetChannel)
	e.PUT("/channels/:id", cc.UpdateChannel)
	e.DELETE("/channels/:id", cc.DeleteChannel)

	// メッセージエンドポイント
	mc := controller.NewMessageController()
	e.POST("/messages", mc.Create)
	e.GET("/messages/:id", mc.Get)
	e.PUT("/messages/:id", mc.Update)
	e.DELETE("/messages/:id", mc.Update)

	// ユーザーエンドポイント
	ur := repository.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur)
	uc := controller.NewUserController(uu)
	e.POST("/users", uc.Create)
	e.GET("/users/:id", uc.Get)
	e.PUT("/users/:id", uc.Update)
	e.DELETE("/users/:id", uc.Delete)

	e.Start(":8080")
}
