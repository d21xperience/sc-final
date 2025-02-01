package models

import "time"

type Student struct {
	StudentID   uint      `gorm:"primaryKey" json:"student_id"`
	Nama        string    `gorm:"size:100" json:"nama"`
	NIS         string    `gorm:"unique" json:"nis"`
	NISN        string    `gorm:"unique" json:"nisn"`
	NIK         string    `gorm:"unique" json:"nik"`
	TptLahir    time.Time `json:"tpt_lahir"`
	TglLahir    time.Time `json:"tgl_lahir"`
	AsalSekolah string    `json:"asal_sekolah"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Certificates []Certificate `gorm:"foreignKey:StudentID" json:"certificates"`
}
