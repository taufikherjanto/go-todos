package main

import (
	"net/http"

	"github.com/labstack/echo"
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

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello human")
	})

	e.Logger.Fatal(e.Start(":1212"))
}
