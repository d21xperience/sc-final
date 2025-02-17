package models

import "time"

type Sekolah struct {
	ID              uint32 `gorm:"primaryKey"`
	NamaSekolah     string `json:"nama_sekolah"`
	NPSN            string `gorm:"unique" json:"npsn"`
	SekolahIDEnkrip string `gorm:"unique" json:"sekolah_id_enkrip"`
	Kecamatan       string `json:"kecamatan"`
	Kabupaten       string `json:"kabupaten"`
	Propinsi        string `json:"propinsi"`
	KodeKecamatan   string `json:"kode_kecamatan"`
	KodeKab         string `json:"kode_kab"`
	KodeProp        string `json:"kode_prop"`
	AlamatJalan     string `json:"alamat_jalan"`
	Status          string `json:"status"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// type SekolahRef struct {
// 	ID              int    `gorm:"primaryKey"`
// 	SekolahIDEnkrip string `gorm:"unique" json:"sekolah_id_enkrip"`
// 	Kecamatan       string `json:"kecamatan"`
// 	Kabupaten       string `json:"kabupaten"`
// 	Propinsi        string `json:"propinsi"`
// 	KodeKecamatan   string `json:"kode_kecamatan"`
// 	KodeKab         string `json:"kode_kab"`
// 	KodeProp        string `json:"kode_prop"`
// 	NamaSekolah     string `json:"nama_sekolah"`
// 	NPSN            string `gorm:"unique" json:"npsn"`
// 	AlamatJalan     string `json:"alamat_jalan"`
// 	Status          string `json:"status"`
// }
