package models

import "time"

type SekolahTenant struct {
	ID            uint32     `gorm:"primaryKey;autoIncrement;column:id"`
	SekolahID     *string    `gorm:"column:sekolah_id"`
	NamaSekolah   string     `gorm:"column:nama_sekolah;not null"`
	NPSN          string     `gorm:"column:npsn;size:8;uniqueIndex:uni_sekolahs_npsn"`
	EnkripID      string     `gorm:"column:enkrip_id;uniqueIndex:uni_sekolahs_sekolah_id_enkrip"`
	Kecamatan     *string    `gorm:"column:kecamatan;size:50"`
	Kabupaten     *string    `gorm:"column:kabupaten;size:50"`
	Propinsi      *string    `gorm:"column:propinsi;size:50"`
	KodeKecamatan *string    `gorm:"column:kode_kecamatan;size:50"`
	KodeKab       *string    `gorm:"column:kode_kab;size:50"`
	KodeProp      *string    `gorm:"column:kode_prop;size:50"`
	AlamatJalan   *string    `gorm:"column:alamat_jalan;size:50"`
	CreatedAt     *time.Time `gorm:"column:created_at"`
	UpdatedAt     *time.Time `gorm:"column:updated_at"`
}

// TableName sets the insert table name for this struct type
func (SekolahTenant) TableName() string {
	return "sekolah_tenant"
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
