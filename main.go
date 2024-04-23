package main

import (
	"github.com/labstack/echo"
	"github.com/taufikherjanto/go-todos/controller"
	"github.com/taufikherjanto/go-todos/database"
)

func main() {
	db := database.InitDb() // initiate database
	defer db.Close()

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	controller.GetAllTodo(e, db)
	controller.DeleteTodo(e, db)
	controller.PatchTodo(e, db)
	controller.CheckTodo(e, db)
	controller.PostTodo(e, db)

	e.Logger.Fatal(e.Start(":1212"))
}
