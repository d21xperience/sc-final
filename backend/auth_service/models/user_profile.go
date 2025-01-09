package models

import "time"

type UserProfile struct {
	ID             uint   `gorm:"primaryKey;autoIncrement" json:"biodata_id"`
	UserID         uint   `gorm:"foreignKey:ID"`
	Nama           string `gorm:"size:100" json:"nama"`
	JK             string `gorm:"size:100" json:"jk"`
	Phone          string `gorm:"size:100" json:"phone"`
	TptLahir       string `gorm:"size:100" json:"tpt_lahir"`
	TglLahir       time.Time
	AlamatJalan    string    `gorm:"size:100" json:"alamat_jalan"`
	KotaKab        string    `gorm:"size:100" json:"kota_kab"`
	Prov           string    `gorm:"size:100" json:"prov"`
	KodePos        string    `gorm:"size:100" json:"kode_pos"`
	NamaAyah       string    `gorm:"size:100" json:"nama_ayah"`
	NamaIbu        string    `gorm:"size:100" json:"nama_ibu"`
	ProfilePicture string    `gorm:"type:varchar(255)" json:"photo_url"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
