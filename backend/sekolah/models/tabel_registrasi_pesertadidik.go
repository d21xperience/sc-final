package models

import "time"

type RegistrasiPesertaDidik struct {
	RegistrasiID        string    `gorm:"type:uuid;primaryKey"`
	JurusanSpID         string    `gorm:"type:uuid;default:null"`
	PesertaDidikID      string    `gorm:"type:uuid;not null"`
	SekolahID           string    `gorm:"type:uuid;not null"`
	JenisPendaftaranID  int       `gorm:"type:numeric(1,0);not null"`
	NIPD                string    `gorm:"type:varchar(18);default:null"`
	TanggalMasukSekolah time.Time `gorm:"type:date;not null"`
	JenisKeluarID       string    `gorm:"type:char(1);default:null"`
	TanggalKeluar       time.Time `gorm:"type:date;default:null"`
	Keterangan          string    `gorm:"type:varchar(128);default:null"`
	NoSKHUN             string    `gorm:"type:char(22);default:null"`
	NoPesertaUjian      string    `gorm:"type:char(22);default:null"`
	NoSeriIjazah        string    `gorm:"type:varchar(80);default:null"`
	APernahPAUD         int       `gorm:"type:numeric(1,0);not null;default:0"`
	APernahTK           int       `gorm:"type:numeric(1,0);not null;default:0"`
	SekolahAsal         string    `gorm:"type:varchar(100);default:null"`
	IDHobby             int       `gorm:"type:numeric(5,0);not null"`
	IDCita              int       `gorm:"type:numeric(5,0);not null"`
	CreateDate          time.Time `gorm:"not null;default:now()"`
	LastUpdate          time.Time `gorm:"not null;default:now()"`
	SoftDelete          int       `gorm:"not null"`
	LastSync            time.Time `gorm:"not null;default:'1901-01-01'"`
	UpdaterID           string    `gorm:"type:uuid;not null"`
}
