package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"github.com/taufikherjanto/go-todos/model"
)

func GetAllTodo(e *echo.Echo, db *sql.DB) {

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
}

func DeleteTodo(e *echo.Echo, db *sql.DB) {

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
}

func PatchTodo(e *echo.Echo, db *sql.DB) {

	e.PATCH("/todos/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")

		var request model.UpdateRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		_, err := db.Exec(
			"UPDATE todos SET title = ?, description = ? WHERE id = ?",
			request.Title,
			request.Description,
			id,
		)

		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")
	})
}

func CheckTodo(e *echo.Echo, db *sql.DB) {

	e.PATCH("/todos/:id/check", func(ctx echo.Context) error {
		id := ctx.Param("id")

		var request model.CheckRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		var doneValue int
		if request.Done {
			doneValue = 1
		}

		_, err := db.Exec(
			"UPDATE todos SET done = ? WHERE id = ?",
			doneValue,
			id,
		)

		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")
	})
}

func PostTodo(e *echo.Echo, db *sql.DB) {

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
}
