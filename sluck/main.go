package main

import (
	"net/http"
	"sluck/controller"
	"sluck/repository"
	"sluck/usecase"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type CustomerValidator struct {
	validator *validator.Validate
}

func (cv *CustomerValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()
	e.Validator = &CustomerValidator{validator: validator.New()}

	ur := repository.NewUserRepository(nil)
	uu := usecase.NewUserUsecase(ur)
	uc := controller.NewUserController(uu)

	e.POST("/users", uc.Create)
	e.Start(":8080")
}
