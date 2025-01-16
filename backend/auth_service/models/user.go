package models

type User struct {
	// gorm.Model
	ID        int64  `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null" json:"email" binding:"required,email"`
	Password  string `gorm:"not null"`
	Role      string
	SekolahID int32
	Sekolah   Sekolah `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
