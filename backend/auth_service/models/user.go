package models

import "time"

type User struct {
	ID                uint64     `gorm:"primaryKey;autoIncrement;column:id"`
	SekolahTenantID   uint32     `gorm:"column:sekolah_tenant_id"`
	Username          string     `gorm:"column:username;not null;uniqueIndex:uni_users_username"`
	Email             string     `gorm:"column:email;not null;uniqueIndex:uni_users_email"`
	Password          string     `gorm:"column:password;not null"`
	Role              string     `gorm:"column:role"`
	IsInitialPassword bool       `gorm:"column:is_initial_password;not null;default:true"`
	InitialPassword   *string    `gorm:"column:initial_password"`
	LastLogin         *time.Time `gorm:"column:last_login"`
	CreatedAt         *time.Time `gorm:"column:created_at"`

	// Relasi dengan tabel Sekolah
	SekolahTenant *SekolahTenant `gorm:"foreignKey:SekolahTenantID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// TableName sets the insert table name for this struct type
func (User) TableName() string {
	return "users"
}
