package main

import (
	"net/http"
	"sluck/controller"
	"sluck/infra"
	"sluck/repository"
	"sluck/usecase"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type CustomerValidor struct {
	validator *validator.Validate
}

func (cv *CustomerValidor) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()
	e.Validator = &CustomerValidor{validator: validator.New()}

	db := infra.Connect()
	ur := repository.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur)
	uc := controller.NewUserController(uu)

	e.POST("/users", uc.Create)
	e.Start(":8080")
}
