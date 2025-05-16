package models

import (
	"time"

	"github.com/google/uuid"
)

type DataNominasiSementara struct {
	ID                          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	PesertaDidikId              uuid.UUID `gorm:"type:uuid"`
	RombonganBelajarId          uuid.UUID `gorm:"type:uuid"`
	ProgramKeahlian             string    `gorm:"type:varchar(100);not null"`
	PaketKeahlian               string    `gorm:"type:varchar(100);not null"`
	SekolahID                   string    `gorm:"type:uuid;not null"`
	NPSN                        string    `gorm:"type:varchar(15);not null"`
	KabupatenKota               string    `gorm:"type:varchar(100);not null"`
	Provinsi                    string    `gorm:"type:varchar(100);not null"`
	Nama                        string    `gorm:"type:varchar(200);not null"`
	TempatLahir                 string    `gorm:"type:varchar(100);not null"`
	TanggalLahir                time.Time `gorm:"type:date;not null"`
	NIS                         string    `gorm:"type:varchar(20);unique;not null"`
	NISN                        string    `gorm:"type:varchar(20);unique;not null"`
	NamaOrtuWali                string    `gorm:"type:varchar(200);not null"`
	SekolahPenyelenggaraUjianUS string    `gorm:"type:varchar(200);not null"`
	SekolahPenyelenggaraUjianUN string    `gorm:"type:varchar(200);not null"`
	AsalSekolah                 string    `gorm:"type:varchar(200);not null"`
	NomorIjazah                 string    `gorm:"type:varchar(50);unique;not null"`
	TempatIjazah                string    `gorm:"type:varchar(100);not null"`
	TanggalIjazah               time.Time `gorm:"type:date;not null"`
	IsComplete                  bool      `gorm:"type:boolean;default:false;not null"`
	TahunAjaranId               string    `gorm:"column:tahun_ajaran_id"`
	// CreatedAt                   time.Time      `gorm:"type:timestamp;default:current_timestamp"`
	// UpdatedAt                   time.Time      `gorm:"type:timestamp;default:current_timestamp"`
	// DeletedAt                   gorm.DeletedAt `gorm:"index"`

	// Relasi opsional
	PesertaDidik     *PesertaDidik     `gorm:"foreignKey:PesertaDidikId;references:PesertaDidikId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	RombonganBelajar *RombonganBelajar `gorm:"foreignKey:RombonganBelajarId;references:RombonganBelajarId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (DataNominasiSementara) TableName() string {
	return "data_nominasi_sementara"
}
