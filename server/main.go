package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func test(c echo.Context) error {
	return c.String(http.StatusOK, "Test")
}

func getStock(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, name)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", test)
	e.GET("/stock/:name", getStock)

	//start server
	e.Logger.Fatal(e.Start(":8080"))
}
