package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresqlInfo struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func InitGorm(config PostgresqlInfo) *gorm.DB {

	dsn := "host=" + config.Host + " user=" + config.Username + " password=" + config.Password + " dbname=" +
		config.Database + " port=" + config.Port + " sslmode=require"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic("failed to connect postgres database")
	}

	return db
}
