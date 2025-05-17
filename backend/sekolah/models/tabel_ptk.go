package models

import (
	"time"

	"github.com/google/uuid"
)

type TabelPTK struct {
	PtkID             uuid.UUID  `gorm:"column:ptk_id;primaryKey"`   // UUID
	Nama              string     `gorm:"column:nama"`                // String
	NUPTK             *string    `gorm:"column:nuptk"`               // Nullable string
	NIP               *string    `gorm:"column:nip"`                 // Nullable string
	JenisPtkID        int32      `gorm:"column:jenis_ptk_id"`        // String
	JenisKelamin      string     `gorm:"column:jenis_kelamin"`       // String
	TempatLahir       string     `gorm:"column:tempat_lahir"`        // String
	TanggalLahir      *time.Time `gorm:"column:tanggal_lahir"`       // String (format tanggal, bisa diubah ke time.Time jika perlu)
	AlamatJalan       string     `gorm:"column:alamat_jalan"`        // String
	StatusKeaktifanID int32      `gorm:"column:status_keaktifan_id"` // String
	GelarDepan        *string    `gorm:"column:gelar_depan"`         // Nullable string
	GelarBelakang     *string    `gorm:"column:gelar_belakang"`      // Nullable string
	NIP_NIY           *string    `gorm:"column:nip_niy"`
}

type PTKTerdaftar struct {
	PtkTerdaftarId uuid.UUID `gorm:"column:ptk_terdaftar_id;primaryKey"`
	PtkID          uuid.UUID `gorm:"column:ptk_id"`
	TahunAjaranId  string    `gorm:"column:tahun_ajaran_id"`
	JenisKeluarId  *string   `gorm:"column:jenis_keluar_id"`

	// Relasi ke PTK
	PTK TabelPTK `gorm:"foreignKey:PtkID;references:PtkID"`
	// Pembelajaran []Pembelajaran `gorm:"foreignKey:PtkTerdaftarId;references:PtkTerdaftarId"`
	PTKPelengkap PtkPelengkap `gorm:"foreignKey:PtkID;references:PtkId"` // Menyertakan PTKPelengkap
}

// Menentukan nama tabel kustom
func (PTKTerdaftar) TableName() string {
	return "tabel_ptk_terdaftar"
}

// Menentukan nama tabel kustom
func (TabelPTK) TableName() string {
	return "tabel_ptk"
}

type PtkPelengkap struct {
	PtkPelengkapId uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"ptk_pelengkap_id"`
	PtkId          uuid.UUID `gorm:"type:uuid;not null" json:"ptk_id"`
	GelarDepan    string `gorm:"type:varchar(20)" json:"gelar_depan,omitempty"`
	GelarBelakang string `gorm:"type:varchar(20)" json:"gelar_belakang,omitempty"`
	NipNiy        string `gorm:"type:varchar(18)" json:"nip_niy,omitempty"`
	AlamatJalan   string `gorm:"type:varchar" json:"alamat_jalan,omitempty"`
	RT            string `gorm:"type:varchar" json:"rt,omitempty"`
	RW            string `gorm:"type:varchar" json:"rw,omitempty"`
	Desa          string `gorm:"type:varchar" json:"desa,omitempty"`
	Kec           string `gorm:"type:varchar" json:"kec,omitempty"`
	Kab           string `gorm:"type:varchar" json:"kab,omitempty"`
	Prov          string `gorm:"type:varchar" json:"prov,omitempty"`

	// Relasi opsional ke tabel PTK
	Ptk *TabelPTK `gorm:"foreignKey:PtkId;references:PtkID"`
}

func (PtkPelengkap) TableName() string {
	return "tabel_ptk"
}
