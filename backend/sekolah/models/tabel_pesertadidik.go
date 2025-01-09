package models

import (
	"time"
)

type PesertaDidik struct {
	PesertaDidikID        string    `gorm:"type:uuid;primaryKey"`
	Nama                  string    `gorm:"type:varchar(100);not null"`
	JenisKelamin          string    `gorm:"type:char(1);not null"`
	NISN                  string    `gorm:"type:char(10);default:null"`
	NIK                   string    `gorm:"type:char(16);default:null"`
	NoKK                  string    `gorm:"type:char(16);default:null"`
	TempatLahir           string    `gorm:"type:varchar(32);default:null"`
	TanggalLahir          time.Time `gorm:"type:date;not null"`
	AgamaID               int16     `gorm:"not null"`
	KebutuhanKhususID     int       `gorm:"not null"`
	AlamatJalan           string    `gorm:"type:varchar(80);not null"`
	RT                    int       `gorm:"type:numeric(2,0);default:null"`
	RW                    int       `gorm:"type:numeric(2,0);default:null"`
	NamaDusun             string    `gorm:"type:varchar(60);default:null"`
	DesaKelurahan         string    `gorm:"type:varchar(60);not null"`
	KodeWilayah           string    `gorm:"type:char(8);not null"`
	KodePos               string    `gorm:"type:char(5);default:null"`
	Lintang               float64   `gorm:"type:numeric(18,12);default:null"`
	Bujur                 float64   `gorm:"type:numeric(18,12);default:null"`
	JenisTinggalID        int       `gorm:"type:numeric(2,0);default:null"`
	AlatTransportasiID    int       `gorm:"type:numeric(2,0);default:null"`
	NikAyah               string    `gorm:"type:char(16);default:null"`
	NikIbu                string    `gorm:"type:char(16);default:null"`
	AnakKeberapa          int       `gorm:"type:numeric(2,0);default:null"`
	NikWali               string    `gorm:"type:char(16);default:null"`
	NomorTeleponRumah     string    `gorm:"type:varchar(20);default:null"`
	NomorTeleponSeluler   string    `gorm:"type:varchar(20);default:null"`
	Email                 string    `gorm:"type:varchar(60);default:null"`
	PenerimaKPS           int       `gorm:"not null"`
	NoKPS                 string    `gorm:"type:varchar(80);default:null"`
	LayakPIP              int       `gorm:"not null;default:0"`
	PenerimaKIP           int       `gorm:"not null"`
	NoKIP                 string    `gorm:"type:varchar(6);default:null"`
	NmKIP                 string    `gorm:"type:varchar(100);default:null"`
	NoKKS                 string    `gorm:"type:varchar(6);default:null"`
	RegAktaLahir          string    `gorm:"type:varchar(80);default:null"`
	IDLayakPIP            int       `gorm:"type:numeric(2,0);default:null"`
	IDBank                string    `gorm:"type:char(3);default:null"`
	RekeningBank          string    `gorm:"type:varchar(20);default:null"`
	NamaKCP               string    `gorm:"type:varchar(100);default:null"`
	RekeningAtasNama      string    `gorm:"type:varchar(100);default:null"`
	StatusData            int       `gorm:"default:null"`
	NamaAyah              string    `gorm:"type:varchar(100);default:null"`
	TahunLahirAyah        int       `gorm:"type:numeric(4,0);default:null"`
	JenjangPendidikanAyah int       `gorm:"type:numeric(2,0);default:null"`
	PekerjaanIDAyah       int       `gorm:"default:null"`
	PenghasilanIDAyah     int       `gorm:"default:null"`
	KebutuhanKhususIDAyah int       `gorm:"not null"`
	NamaIbuKandung        string    `gorm:"type:varchar(100);not null"`
	TahunLahirIbu         int       `gorm:"type:numeric(4,0);default:null"`
	JenjangPendidikanIbu  int       `gorm:"type:numeric(2,0);default:null"`
	PenghasilanIDIbu      int       `gorm:"default:null"`
	PekerjaanIDIbu        int       `gorm:"default:null"`
	KebutuhanKhususIDIbu  int       `gorm:"not null"`
	NamaWali              string    `gorm:"type:varchar(30);default:null"`
	TahunLahirWali        int       `gorm:"type:numeric(4,0);default:null"`
	JenjangPendidikanWali int       `gorm:"type:numeric(2,0);default:null"`
	PekerjaanIDWali       int       `gorm:"default:null"`
	PenghasilanIDWali     int       `gorm:"default:null"`
	Kewarganegaraan       string    `gorm:"type:char(2);not null"`
	PekerjaanID           int       `gorm:"default:null"`
	CreateDate            time.Time `gorm:"not null;default:now()"`
	LastUpdate            time.Time `gorm:"not null;default:now()"`
	SoftDelete            int       `gorm:"not null"`
	LastSync              time.Time `gorm:"not null;default:'1901-01-01'"`
	UpdaterID             string    `gorm:"type:uuid;not null"`
}

func (PesertaDidik) TableName() string {
	return "peserta_didik"
}
