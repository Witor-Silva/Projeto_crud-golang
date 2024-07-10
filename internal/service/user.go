package service

import (
	"go-guapi/internal/repository"
	"go-guapi/internal/request"
	"time"

	"github.com/google/uuid"
)

type User struct {
	UseRepositoryInterface repository.UseRepositoryInterface
}

type UserServiveInterface interface {
	Create(data request.UserRequestCreate) error
}

func NewUser(UseRepositoryInterface repository.UseRepositoryInterface) *User {
	return &User{UseRepositoryInterface}
}

func (u *User) Create(data request.UserRequestCreate) error {
	user := request.UserDTO{
		ID:        uuid.New(),
		Name:      data.Name,
		Age:       data.Age,
		Status:    true,
		CreatedAt: time.Now(),
	}

	err := u.UseRepositoryInterface.Create(user)
	if err != nil {

	}

	return nil

}
