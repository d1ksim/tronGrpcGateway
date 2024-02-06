package models

import "gorm.io/gorm"

type Accounts struct {
	gorm.Model
	TelegramID int64  `gorm:"unique"`
	Username   string `gorm:"unique"`
	ApiKey     string `gorm:"unique"`
}
