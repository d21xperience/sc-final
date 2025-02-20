package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Pembelajaran struct {
	PembelajaranID     uuid.UUID  `gorm:"type:uuid;primaryKey"`
	RombonganBelajarID uuid.UUID  `gorm:"type:uuid;not null"`
	MataPelajaranID    int        `gorm:"type:int;not null"`
	SemesterID         string     `gorm:"type:varchar(5);not null"`
	PtkTerdaftarID     *uuid.UUID `gorm:"type:uuid;default:NULL"`
	StatusDiKurikulum  *int       `gorm:"type:numeric(2,0);default:NULL"`
	NamaMataPelajaran  *string    `gorm:"type:varchar(50);default:NULL"`
	IndukPembelajaran  *uuid.UUID `gorm:"type:uuid;default:NULL"`
	IsDapo             *int       `gorm:"type:numeric(1,0);default:1"`

	// Relasi
	RombonganBelajar RombonganBelajar `gorm:"foreignKey:RombonganBelajarID;references:RombonganBelajarID"`
	PtkTerdaftar     PTKTerdaftar     `gorm:"foreignKey:PtkTerdaftarID;references:PtkTerdaftarID"`
}

func (p *Pembelajaran) BeforeCreate(tx *gorm.DB) (err error) {
	if p.PembelajaranID == uuid.Nil {
		p.PembelajaranID = uuid.New()
	}
	return
}

func (Pembelajaran) TableName() string {
	return "tabel_pembelajaran"
}
