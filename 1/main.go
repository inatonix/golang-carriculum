package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func validateUser(name string, age int) error {
	if name == "" {
		return errors.New("name is required")
	}
	if len(name) > 100 {
		return errors.New("name must be 100 characters or fewer")
	}
	if age <= 0 || age >= 150 {
		return errors.New("age must be between 1 and 149")
	}
	return nil
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	e := echo.New()

	// ミドルウェアを追加
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, echoフレームワーク!")
	})

	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	e.PUT("/users/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		name := c.FormValue("name")
		age, _ := strconv.Atoi(c.FormValue("age"))

		if err := validateUser(name, age); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		result, err := db.Exec("UPDATE users SET name = ?, age = ? WHERE id = ?", name, age, id)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update user")
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}

		return c.JSON(http.StatusOK, &User{ID: id, Name: name, Age: age})
	})

	e.DELETE("/users/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		result, err := db.Exec("DELETE FROM users WHERE id = ?", id)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete user")
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}

		return c.NoContent(http.StatusNoContent)
	})

	e.GET("/users/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		var user User
		err := db.QueryRow("SELECT id, name, age FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			if err == sql.ErrNoRows {
				return echo.NewHTTPError(http.StatusNotFound, "User not found")
			}
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user")
		}

		return c.JSON(http.StatusOK, user)
	})

	e.Start(":8080")
}
