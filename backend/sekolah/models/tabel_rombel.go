package models

import "time"

type RombonganBelajar struct {
	RombonganBelajarID  string    `gorm:"type:uuid;primaryKey"`
	SemesterID          string    `gorm:"type:char(5);not null"`
	IDRuang             string    `gorm:"type:uuid;not null"`
	SekolahID           string    `gorm:"type:uuid;not null"`
	TingkatPendidikanID int       `gorm:"type:numeric(2,0);not null"`
	JurusanSpID         string    `gorm:"type:uuid;default:null"`
	KurikulumID         int16     `gorm:"not null"`
	Nama                string    `gorm:"type:varchar(30);not null"`
	PtkID               string    `gorm:"type:uuid;default:null"`
	MovingClass         int       `gorm:"type:numeric(1,0);not null"`
	JenisRombel         int       `gorm:"type:numeric(2,0);not null"`
	SKS                 int       `gorm:"type:numeric(2,0);default:0"`
	CreateDate          time.Time `gorm:"not null;default:now()"`
	LastUpdate          time.Time `gorm:"not null;default:now()"`
	SoftDelete          int       `gorm:"not null"`
	LastSync            time.Time `gorm:"not null;default:'1901-01-01'"`
	UpdaterID           string    `gorm:"type:uuid;not null"`
}
