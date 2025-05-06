package models

import (
	"time"

	"gorm.io/gorm"
)

type UserProfile struct {
	ID          uint64         `gorm:"primaryKey;autoIncrement;column:id"`
	UserID      uint64         `gorm:"column:user_id"`
	Nama        string         `gorm:"column:nama;size:100"`
	JK          string         `gorm:"column:jk;size:100"`
	Phone       *string        `gorm:"column:phone;size:100"`
	TptLahir    *string        `gorm:"column:tpt_lahir;size:100"`
	TglLahir    *time.Time     `gorm:"column:tgl_lahir"`
	AlamatJalan *string        `gorm:"column:alamat_jalan;size:100"`
	KotaKab     *string        `gorm:"column:kota_kab;size:100"`
	Prov        *string        `gorm:"column:prov;size:100"`
	KodePos     *string        `gorm:"column:kode_pos;size:100"`
	NamaAyah    *string        `gorm:"column:nama_ayah;size:100"`
	NamaIbu     *string        `gorm:"column:nama_ibu;size:100"`
	PhotoURL    *string        `gorm:"column:photo_url;size:255"`
	CreatedAt   *time.Time     `gorm:"column:created_at"`
	UpdatedAt   *time.Time     `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index:idx_user_profiles_deleted_at"`

	// Relasi ke User
	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// TableName overrides the default table name
func (UserProfile) TableName() string {
	return "user_profiles"
}
