package handler

import (
	"go-guapi/internal/request"
	"go-guapi/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	UserServiveInterface service.UserServiveInterface
}

func NewUser(UserServiveInterface service.UserServiveInterface) *User {
	return &User{UserServiveInterface}
}

func (u *User) Create(c echo.Context) error {
	var data request.UserRequestCreate
	err := c.Bind(&data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = u.UserServiveInterface.Create(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "sucess")
}
