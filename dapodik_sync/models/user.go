package models

import (
	"time"
)

type User struct {
	UserID    uint      `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	Role      string    `gorm:"not null"`
	SekolahID uint      `gorm:"not null"`
	LastLogin time.Time `gorm:"default:NULL"`
}
