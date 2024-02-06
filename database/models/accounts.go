package models

import "gorm.io/gorm"

type Accounts struct {
	gorm.Model
	telegramId int64  `gorm:"unique"`
	username   string `gorm:"unique"`
	apiKey     string `gorm:"unique"`
}
