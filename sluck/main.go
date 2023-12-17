package main

import (
	"net/http"
	"sluck/controller"
	"sluck/infra"
	"sluck/repository"
	"sluck/transaction"
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
	transaction := transaction.NewTransaction(db)

	mr := repository.NewMessageRepository(db)
	ur := repository.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur, mr, transaction)
	uc := controller.NewUserController(uu)

	e.POST("/users", uc.Create)
	e.GET("/users/:id", uc.Get)
	e.PUT("/users", uc.Update)
	e.DELETE("/users/:id", uc.Delete)

	e.Start(":8080")
}
