package entity

import (
	"go-guapi/internal/request"
	"time"

	"github.com/google/uuid"
)

type UserEntity struct {
	ID        uuid.UUID `gorm:"primary_key"`
	Name      string    `gorm:"column:name"`
	Age       int       `gorm:"column:age"`
	Status    bool      `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (u *UserEntity) ToDomain(data request.UserDTO) {
	u.ID = data.ID
	u.Name = data.Name
	u.Age = data.Age
	u.Status = data.Status
	u.CreatedAt = data.CreatedAt
}

func (u *UserEntity) TableName() string {
	return "users_gabriel"
}
