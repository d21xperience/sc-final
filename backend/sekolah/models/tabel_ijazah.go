package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Ijazah struct {
	ID                          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	PesertaDidikId              string    `gorm:"column:peserta_didik_id;uique;not null"`
	ProgramKeahlian             string    `gorm:"type:varchar(100);not null"`
	PaketKeahlian               string    `gorm:"type:varchar(100);not null"`
	SekolahID                   uuid.UUID `gorm:"type:uuid;not null;index"`
	NPSN                        string    `gorm:"type:varchar(15);not null"`
	KabupatenKota               string    `gorm:"type:varchar(100);not null"`
	Provinsi                    string    `gorm:"type:varchar(100);not null"`
	Nama                        string    `gorm:"type:varchar(200);not null"`
	TempatLahir                 string    `gorm:"type:varchar(100);not null"`
	TanggalLahir                string    `gorm:"type:date;not null"`
	Nis                         string    `gorm:"type:varchar(20);unique;not null"`
	NISN                        string    `gorm:"type:varchar(20);unique;not null"`
	NamaOrtuWali                string    `gorm:"type:varchar(200);not null"`
	SekolahPenyelenggaraUjianUS string    `gorm:"type:varchar(200);not null"`
	SekolahPenyelenggaraUjianUN string    `gorm:"type:varchar(200);not null"`
	AsalSekolah                 string    `gorm:"type:varchar(200);not null"`
	NomorIjazah                 string    `gorm:"type:varchar(50);unique;not null"`
	TempatIjazah                string    `gorm:"type:varchar(100);not null"`
	TanggalIjazah               string    `gorm:"type:date;not null"`
	NoIjazah                    string    `gorm:"type:varchar(50);unique;not null"`
	CreatedAt                   time.Time
	UpdatedAt                   time.Time
	DeletedAt                   gorm.DeletedAt `gorm:"index"`
	Status                      string         `json:"status"` // DRAFT, PENDING, ACTIVE, REVOKED
	// Relasi
	// PesertaDidik PesertaDidik `gorm:"foreignKey:PesertaDidikID"`
	// NilaiUjianSekolah NIlaiRerataAkhir
}

func (Ijazah) TableName() string {
	return "ijazah"
}
