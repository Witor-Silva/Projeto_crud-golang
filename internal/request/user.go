package request

import (
	"time"

	"github.com/google/uuid"
)

type UserRequestCreate struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserDTO struct {
	ID        uuid.UUID `json:"ID"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
