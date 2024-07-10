package main

import (
	"go-guapi/cmd"
	"go-guapi/config"
	"go-guapi/infra/database"
	"go-guapi/infra/migration"
	"go-guapi/internal/handler"
	"go-guapi/internal/repository"
	"go-guapi/internal/service"
)

func main() {
	loadedInfo := config.LoadConfig()
	conn := database.InitGorm(database.PostgresqlInfo{
		Host:     loadedInfo.DBHost,
		Port:     loadedInfo.DBPort,
		Username: loadedInfo.DBUser,
		Password: loadedInfo.DBPass,
		Database: loadedInfo.DBName,
	})
	migrationConn := migration.NewDatabaseMakeMigration(conn)
	migrationConn.MakeMigrations()

	instanceRepository := repository.NewUser(conn)
	instanceService := service.NewUser(instanceRepository)
	instanceHandler := handler.NewUser(instanceService)

	cmd.Server(loadedInfo, *instanceHandler)
}
