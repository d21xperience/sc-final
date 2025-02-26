package models

import "time"

type User struct {
	// gorm.Model
	ID                uint64 `gorm:"primaryKey"`
	Username          string `gorm:"unique;not null"`
	Email             string `gorm:"unique" json:"email" binding:"required,email"`
	Password          string `gorm:"not null"`
	Role              string
	SekolahID         uint32
	IsInitialPassword bool `gorm:"default:true;not null"`
	InitialPassword   string
	LastLogin         time.Time `gorm:"column:last_login"`
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	Sekolah           Sekolah   `gorm:"foreignKey:SekolahID;references:ID"`
}
