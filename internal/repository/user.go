package repository

import (
	"go-guapi/entity"
	"go-guapi/internal/request"

	"gorm.io/gorm"
)

type User struct {
	conn *gorm.DB
}

type UseRepositoryInterface interface {
	Create(data request.UserDTO) error
}

func NewUser(conn *gorm.DB) *User {
	return &User{
		conn: conn,
	}
}

func (u *User) Create(data request.UserDTO) error {
	var UserEntity entity.UserEntity
	UserEntity.ToDomain(data)

	err := u.conn.Create(&UserEntity)
	if err != nil {
		return err.Error
	}
	return nil
}
