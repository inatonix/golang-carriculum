package main

import (
	"2/controller"
	"2/repository"
	"2/usecase"
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/mattn/go-sqlite3"
)

func initDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./test.db")
	return db, err
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY, title TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	repository := repository.NewTaskRepository(db)
	usecase := usecase.NewTaskUsecase(repository)
	taskController := controller.NewTaskController(usecase)

	e.GET("/tasks/:id", taskController.Get)
	e.POST("/tasks", taskController.Create)
	e.Start(":8080")
}
