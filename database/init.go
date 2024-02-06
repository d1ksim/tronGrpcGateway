package database

import (
	"github.com/d1mpi/tronGrpcGateway/database/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DataBase *gorm.DB

func InitPostgresDatabase() {
	db, err := gorm.Open(postgres.Open(viper.GetString("database.dsn")), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	err = db.AutoMigrate(&models.Accounts{})
	if err != nil {
		return
	}

	DataBase = db
}
