package models

import (
	"gorm.io/gorm"
)

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

type SekolahTenant struct {
	gorm.Model
	NamaSekolah     string `gorm:"column:nama_sekolah"`
	SekolahTenantId uint32 `gorm:"column:sekolah_tenant_id"`
	SchemaName      string `gorm:"column:schema_name"`
	// BentukPendidikanID uint16           `gorm:"not null;default:0;column:bentuk_pendidikan_id"`
	// BentukPendidikan   BentukPendidikan `gorm:"foreignKey:BentukPendidikanID;references:BentukPendidikanID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

type SchemaLog struct {
	gorm.Model
	SchemaName string
}

// func (SekolahTenant) TableName() string {
// 	return "sekolah_tenant"
// }
