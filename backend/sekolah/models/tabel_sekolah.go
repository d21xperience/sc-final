package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sekolah struct {
	gorm.Model
	SekolahID                uuid.UUID `gorm:"primaryKey"`
	Nama                     string    `gorm:"type:varchar(255)"`
	NamaNomenklatur          string    `gorm:"type:varchar(255)"`
	NSS                      string    `gorm:"type:varchar(255)"`
	NPSN                     string    `gorm:"type:varchar(255)"`
	BentukPendidikanID       int
	AlamatJalan              string `gorm:"type:varchar(255)"`
	RT                       string `gorm:"type:varchar(255)"`
	RW                       string `gorm:"type:varchar(255)"`
	NamaDusun                string `gorm:"type:varchar(255)"`
	DesaKelurahan            string `gorm:"type:varchar(255)"`
	KodeWilayah              string `gorm:"type:varchar(255)"`
	KodePos                  string `gorm:"type:varchar(255)"`
	Lintang                  int
	Bujur                    int
	NomorTelepon             string `gorm:"type:varchar(255)"`
	NomorFax                 string `gorm:"type:varchar(255)"`
	Email                    string `gorm:"type:varchar(255)"`
	Website                  string `gorm:"type:varchar(255)"`
	KebutuhanKhususID        int
	StatusSekolah            int
	SkPendirianSekolah       string `gorm:"type:varchar(255)"`
	TanggalSkPendirian       string `gorm:"type:varchar(255)"`
	StatusKepemilikanID      string `gorm:"type:varchar(255)"`
	YayasanId                string `gorm:"type:varchar(255)"`
	SkIzinOperasional        string `gorm:"type:varchar(255)"`
	TanggalSkIzinOperasional string `gorm:"type:varchar(255)"`
	NoRekening               string `gorm:"type:varchar(255)"`
	NamaBank                 string `gorm:"type:varchar(255)"`
	CabangKcpUnit            string `gorm:"type:varchar(255)"`
	RekeningAtasNama         string `gorm:"type:varchar(255)"`
	MBS                      string `gorm:"type:varchar(255)"`
	LuasTanahMilik           string `gorm:"type:varchar(255)"`
	LuasTanahBukanMilik      string `gorm:"type:varchar(255)"`
	KodeRegistrasi           string `gorm:"type:varchar(255)"`
	Npwp                     string `gorm:"type:varchar(255)"`
	NmWp                     string `gorm:"type:varchar(255)"`
	Keaktifan                string `gorm:"type:varchar(255)"`
	Flag                     string `gorm:"type:varchar(255)"`
	// CreateDate               time.Time
	// LastUpdate               time.Time
	// SoftDelete               time.Time
	LastSync  time.Time
	UpdaterId string `gorm:"type:varchar(255)"`
	Results   string `gorm:"type:varchar(255)"`
	Id        string `gorm:"type:varchar(255)"`
	Start     string `gorm:"type:varchar(255)"`
	Limit     string `gorm:"type:varchar(255)"`
}
