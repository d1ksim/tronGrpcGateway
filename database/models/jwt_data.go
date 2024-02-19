package models

import (
	"gorm.io/gorm"
)

type JwtData struct {
	gorm.Model
	TelegramID  int64  `gorm:"unique"`
	AccessToken string `gorm:"unique"`
	ExpiresAt   int64
}
