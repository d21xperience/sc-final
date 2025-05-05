package models

import "time"

type User struct {
	ID                int64      `gorm:"primaryKey;autoIncrement;column:id"`
	Username          string     `gorm:"column:username;not null;uniqueIndex:uni_users_username"`
	Email             string     `gorm:"column:email;not null;uniqueIndex:uni_users_email"`
	Password          string     `gorm:"column:password;not null"`
	SekolahID         *string    `gorm:"column:sekolah_id"`
	Role              *string    `gorm:"column:role"`
	IdSekolahAnggota  *int64     `gorm:"column:id_sekolah_anggota"`
	IsInitialPassword bool       `gorm:"column:is_initial_password;not null;default:true"`
	InitialPassword   *string    `gorm:"column:initial_password"`
	LastLogin         *time.Time `gorm:"column:last_login"`
	CreatedAt         *time.Time `gorm:"column:created_at"`

	// Relasi dengan tabel Sekolah
	SekolahAnggota    *Sekolah   `gorm:"foreignKey:IdSekolahAnggota;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// TableName sets the insert table name for this struct type
func (User) TableName() string {
	return "users"
}