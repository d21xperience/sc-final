package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Pembelajaran struct {
	PembelajaranId     uuid.UUID  `gorm:"type:uuid;primaryKey"`
	RombonganBelajarId uuid.UUID  `gorm:"type:uuid;not null"`
	MataPelajaranId    int        `gorm:"type:int;not null"`
	SemesterId         string     `gorm:"type:varchar(5);not null"`
	PtkTerdaftarId     *uuid.UUID `gorm:"type:uuid;default:NULL"`
	StatusDiKurikulum  *int       `gorm:"type:numeric(2,0);default:NULL"`
	NamaMataPelajaran  *string    `gorm:"type:varchar(50);default:NULL"`
	IndukPembelajaran  *uuid.UUID `gorm:"type:uuid;default:NULL"`
	IsDapo             *int       `gorm:"type:numeric(1,0);default:1"`

	// Relasi
	// RombonganBelajar RombonganBelajar `gorm:"foreignKey:RombonganBelajarId;references:RombonganBelajarId"`
	// MataPelajaran    MataPelajaran    `gorm:"foreignKey:MataPelajaranId;references:MataPelajaranID"`
	PTKTerdaftar PTKTerdaftar `gorm:"foreignKey:PtkTerdaftarId;references:PtkTerdaftarId"`
}

func (p *Pembelajaran) BeforeCreate(tx *gorm.DB) (err error) {
	if p.PembelajaranId == uuid.Nil {
		p.PembelajaranId = uuid.New()
	}
	return
}

func (Pembelajaran) TableName() string {
	return "tabel_pembelajaran"
}
