package main

import "gorm.io/gorm"

type Accounts struct {
	gorm.Model
	telegramId int64  `gorm:"unique"`
	apiKey     string `gorm:"unique"`
}
