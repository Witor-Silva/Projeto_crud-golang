package cmd

import (
	"go-guapi/config"
	"go-guapi/internal/handler"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Server(env config.LoadedEnv, user handler.User) {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})
	e.POST("/user/create", user.Create)
	e.Logger.Fatal(e.Start(env.ApiPort))
}
