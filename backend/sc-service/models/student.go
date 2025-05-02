package models

import (
	"time"

	"github.com/google/uuid"
)

type Student struct {
	PesertaDidikId uuid.UUID `gorm:"primaryKey" json:"student_id"`
	Nama           string    `gorm:"size:100" json:"nama"`
	NIS            string    `gorm:"unique" json:"nis"`
	NISN           string    `gorm:"unique" json:"nisn"`
	NIK            string    `gorm:"unique" json:"nik"`
	TptLahir       string    `gorm:"column:tpt_lahir"`
	TglLahir       time.Time `gorm:"column:tgl_lahir"`
	AsalSekolah    string    `gorm:"column:asal_sekolah"`
	// CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	// UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	// Certificates []Certificate `gorm:"foreignKey:StudentID" json:"certificates"`
}

type Sekolah struct {
	SekolahID           string `gorm:"column:sekolah_id"`
	Nama                string `gorm:"column:nama"`
	Npsn                string `gorm:"column:npsn"`
	Nss                 string `gorm:"column:nss"`
	Alamat              string `gorm:"column:alamat"`
	KdPos               string `gorm:"column:kd_pos"`
	Telepon             string `gorm:"column:telepon"`
	Fax                 string `gorm:"column:fax"`
	Kelurahan           string `gorm:"column:kelurahan"`
	Kecamatan           string `gorm:"column:kecamatan"`
	KabKota             string `gorm:"column:kab_kota"`
	Propinsi            string `gorm:"column:propinsi"`
	Website             string `gorm:"column:website"`
	Email               string `gorm:"column:email"`
	NmKepsek            string `gorm:"column:nm_kepsek"`
	NipKepsek           string `gorm:"column:nip_kepsek"`
	NiyKepsek           string `gorm:"column:niy_kepsek"`
	StatusKepemilikanId int32  `gorm:"column:status_kepemilikan_id"`
	KodeAktivasi        string `gorm:"column:kode_aktivasi"`
	BentukPendidikanId  int32  `gorm:"column:bentuk_pendidikan_id"`
	JenjangPendidikanId int32  `gorm:"column:jenjang_pendidikan_id"`
}