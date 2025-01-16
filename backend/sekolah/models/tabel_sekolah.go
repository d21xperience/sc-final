package models

import (
	"gorm.io/gorm"
)

type Sekolah struct {
	SekolahID           string `json:"sekolah_id"`
	Nama                string `json:"nama"`
	Npsn                string `json:"npsn"`
	Nss                 string `json:"nss"`
	Alamat              string `json:"alamat"`
	KdPos               string `json:"kd_pos"`
	Telepon             string `json:"telepon"`
	Fax                 string `json:"fax"`
	Kelurahan           string `json:"kelurahan"`
	Kecamatan           string `json:"kecamatan"`
	KabKota             string `json:"kab_kota"`
	Propinsi            string `json:"propinsi"`
	Website             string `json:"website"`
	Email               string `json:"email"`
	NmKepsek            string `json:"nm_kepsek"`
	NipKepsek           string `json:"nip_kepsek"`
	NiyKepsek           string `json:"niy_kepsek"`
	StatusKepemilikanId string `json:"status_kepemilikan_id"`
	KodeAktivasi        string `json:"kode_aktivasi"`
	Jenjang             string `json:"jenjang"`
	BentukPendidikanId  string `json:"bentuk_pendidikan_id"`
}

type SekolahTabelTenant struct {
	gorm.Model
	Nama       string `gorm:"unique;not null"`
	SekolahID  int    `gorm:"unique;not null"`
	NamaSchema string `gorm:"unique;not null"`
}
