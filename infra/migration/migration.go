package migration

import (
	"go-guapi/entity"

	"gorm.io/gorm"
)

type DatabaseMakeMigration struct {
	conn *gorm.DB
}

func NewDatabaseMakeMigration(conn *gorm.DB) *DatabaseMakeMigration {
	return &DatabaseMakeMigration{conn: conn}
}

func (m *DatabaseMakeMigration) MakeMigrations() {
	err := m.conn.Migrator().AutoMigrate(&entity.UserEntity{})
	if err != nil {
		panic(err)
	}
}
