package models

import (
	"gorm.io/gorm"
	"time"
)

type Accounts struct {
	gorm.Model
	TelegramID     int64  `gorm:"unique"`
	Username       string `gorm:"unique"`
	AndroidUnique  string `gorm:"unique"`
	LastPayment    *time.Time
	SubscribePrice int16 `gorm:"default:50"`
	Active         bool  `gorm:"default:true"`
}
