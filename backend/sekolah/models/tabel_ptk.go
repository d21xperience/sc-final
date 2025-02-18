package models

import "github.com/google/uuid"

type TabelPTK struct {
	PtkID             string  `gorm:"column:ptk_id;primaryKey"`   // UUID
	Nama              string  `gorm:"column:nama"`                // String
	NIP               *string `gorm:"column:nip"`                 // Nullable string
	JenisPtkID        int32   `gorm:"column:jenis_ptk_id"`        // String
	JenisKelamin      string  `gorm:"column:jenis_kelamin"`       // String
	TempatLahir       string  `gorm:"column:tempat_lahir"`        // String
	TanggalLahir      string  `gorm:"column:tanggal_lahir"`       // String (format tanggal, bisa diubah ke time.Time jika perlu)
	NUPTK             *string `gorm:"column:nuptk"`               // Nullable string
	AlamatJalan       string  `gorm:"column:alamat_jalan"`        // String
	StatusKeaktifanID int32   `gorm:"column:status_keaktifan_id"` // String
}

type PTKPelengkap struct {
	PTKPelengkapID string  `json:"ptk_pelengkap_id"` // UUID
	PTKID          string  `json:"ptk_id"`           // UUID
	GelarDepan     *string `json:"gelar_depan"`      // Nullable string
	GelarBelakang  *string `json:"gelar_belakang"`   // Nullable string
}

type PTKTerdaftar struct {
	PtkTerdaftarID uuid.UUID `gorm:"column:ptk_terdaftar_id;primaryKey"`
	PtkID          uuid.UUID `gorm:"column:ptk_id"`
	TahunAjaranID  string    `gorm:"column:tahun_ajaran_id"`
	JenisKeluarID  *string   `gorm:"column:jenis_keluar_id"`

	// Relasi ke PTK
	PTK TabelPTK `gorm:"foreignKey:PtkID;references:PtkID"`
}

// Menentukan nama tabel kustom
func (PTKTerdaftar) TableName() string {
	return "tabel_ptk_terdaftar"
}

// Menentukan nama tabel kustom
func (TabelPTK) TableName() string {
	return "tabel_ptk"
}
