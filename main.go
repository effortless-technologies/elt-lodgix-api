package main

import (
	"net/http"

	"github.com/effortless-technologies/elt-lodgix-api/server"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	e.GET("/properties", server.GetProperties)

	e.Logger.Fatal(e.Start(":1323"))
}
