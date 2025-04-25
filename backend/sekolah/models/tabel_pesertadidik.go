package models

import (
	"time"

	"github.com/google/uuid"
)

type PesertaDidik struct {
	PesertaDidikId  string     `gorm:"column:peserta_didik_id;primaryKey"` // STRING
	Nis             string     `gorm:"column:nis"`                         // String
	Nisn            string     `gorm:"column:nisn"`                        // String
	NmSiswa         string     `gorm:"column:nm_siswa"`                    // String
	TempatLahir     string     `gorm:"column:tempat_lahir"`                // String
	TanggalLahir    *time.Time `gorm:"column:tanggal_lahir"`               // String (format tanggal, bisa diubah ke time.Time jika perlu)
	JenisKelamin    string     `gorm:"column:jenis_kelamin"`               // String
	Agama           string     `gorm:"column:agama"`                       // String
	AlamatSiswa     *string    `gorm:"column:alamat_siswa"`                // Nullable string
	TeleponSiswa    string     `gorm:"column:telepon_siswa"`               // String
	DiterimaTanggal *time.Time `gorm:"column:diterima_tanggal"`            // String (format tanggal, bisa diubah ke time.Time jika perlu)
	NmAyah          string     `gorm:"column:nm_ayah"`                     // String
	NmIbu           string     `gorm:"column:nm_ibu"`                      // String
	PekerjaanAyah   string     `gorm:"column:pekerjaan_ayah"`              // String
	PekerjaanIbu    string     `gorm:"column:pekerjaan_ibu"`               // String
	NmWali          *string    `gorm:"column:nm_wali"`                     // Nullable string
	PekerjaanWali   *string    `gorm:"column:pekerjaan_wali"`              // Nullable string
	Nik             *string    `gorm:"column:nik"`
}

type PesertaDidikPelengkap struct {
	PelengkapSiswaId string       `gorm:"column:pelengkap_siswa_id;primaryKey"` // UUID
	PesertaDidikId   *string      `gorm:"column:peserta_didik_id"`              // UUID
	StatusDalamKel   *string      `gorm:"column:status_dalam_kel"`              // Nullable string
	AnakKe           *string      `gorm:"column:anak_ke"`                       // Integer
	SekolahAsal      string       `gorm:"column:sekolah_asal"`                  // Non-nullable string
	DiterimaKelas    *string      `gorm:"column:diterima_kelas"`                // Nullable string
	AlamatOrtu       *string      `gorm:"column:alamat_ortu"`                   // Nullable string
	TeleponOrtu      *string      `gorm:"column:telepon_ortu"`                  // Nullable string
	AlamatWali       *string      `gorm:"column:alamat_wali"`                   // Nullable string
	TeleponWali      *string      `gorm:"column:telepon_wali"`                  // Nullable string
	FotoSiswa        *string      `gorm:"column:foto_siswa"`                    // Nullable string
	PesertaDidik     PesertaDidik `gorm:"foreignKey:PesertaDidikId;references:PesertaDidikId"`
}

type Kenaikan struct {
	KdKenaikan      uuid.UUID `gorm:"type:uuid;primaryKey;column:kd_kenaikan"`
	SemesterId      string    `gorm:"type:char(5);not null;column:semester_id"`
	AnggotaRombelId uuid.UUID `gorm:"type:uuid;not null;column:anggota_rombel_id"`
	PesertaDidikId  uuid.UUID `gorm:"type:uuid;column:peserta_didik_id"`
	Kenaikan        int32     `gorm:"type:numeric(3,0);column:kenaikan"`
	Tingkat         int32     `gorm:"type:numeric(3,0);column:tingkat"`

	// Relasi
	AnggotaRombel RombelAnggota `gorm:"foreignKey:AnggotaRombelId;references:AnggotaRombelId"`
	// Semester      Semester      `gorm:"foreignKey:SemesterId;references:SemesterId"`
	// PesertaDidik  PesertaDidik  `gorm:"foreignKey:PesertaDidikId;references:PesertaDidikId"`
}

func (PesertaDidik) TableName() string {
	return "tabel_siswa"
}
func (PesertaDidikPelengkap) TableName() string {
	return "tabel_siswa_pelengkap"
}
func (Kenaikan) TableName() string {
	return "tabel_kenaikan"
}
