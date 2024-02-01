package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgresDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password= dbname=gorm port=5432 sslmode=disable TimeZone=Europe/Moscow"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}
