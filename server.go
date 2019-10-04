package main

import (
	"net/http"
	"github.com/labstack/echo"
)

func main()	{
	e := echo.New()

    e.GET("/users/:id", getUser)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

