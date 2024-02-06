package database

import (
	"github.com/d1mpi/tronGrpcGateway/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DataBase *gorm.DB

func InitPostgresDatabase() {
	dsn := "host=localhost user=postgres password=qwerty dbname=gorm port=5432 sslmode=disable TimeZone=Europe/Moscow"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	err = db.AutoMigrate(&models.Accounts{})
	if err != nil {
		return
	}

	DataBase = db
}
