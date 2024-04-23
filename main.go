package main

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"github.com/taufikherjanto/go-todos/database"
	"github.com/taufikherjanto/go-todos/model"
)

func main() {
	db := database.InitDb() // initiate database
	defer db.Close()

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello human")
	})

	e.GET("/todos", func(ctx echo.Context) error {
		rows, err := db.Query(
			"SELECT * FROM todos",
		)

		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		var response []model.TodoResponse
		for rows.Next() {
			var id int
			var title string
			var description string
			var done int

			err = rows.Scan(&id, &title, &description, &done)
			if err != nil {
				return ctx.String(http.StatusInternalServerError, err.Error())
			}

			var todo model.TodoResponse
			todo.Id = id
			todo.Title = title
			todo.Description = description
			if done == 1 {
				todo.Done = true
			}

			response = append(response, todo)
		}

		return ctx.JSON(http.StatusOK, response)
	})

	e.DELETE("/todos/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")

		_, err := db.Exec(
			"DELETE FROM todos WHERE Id = ?",
			id,
		)

		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")
	})

	e.POST("/todos", func(ctx echo.Context) error {
		var request model.CreateRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		_, err := db.Exec(
			"INSERT INTO todos (title, description, done) VALUES (?, ?, 0)",
			request.Title,
			request.Description,
		)

		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")
	})

	e.Logger.Fatal(e.Start(":1212"))
}
