package main

import (
	"net/http"
	"todo-app/src/handler"
	"todo-app/src/server/database"
	"todo-app/src/server/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db := database.InitializeDatabase()
	router.InitializeRouter(e, handler.InitializeHandler(db.DB))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":3030"))
}
