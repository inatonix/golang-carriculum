package controller

import (
	"context"
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

func NewUserController(u usecase.UserUsecase) UserController {
	return &userController{u: u}
}

func (c userController) Create(ctx echo.Context) error {
	var req UserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	u := ToModel(&req)
	// fmt.Println(u, ctx.Request().Context())
	cc := context.Background()
	id, err := c.u.Create(cc, u)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, "created user_id is "+id)
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
