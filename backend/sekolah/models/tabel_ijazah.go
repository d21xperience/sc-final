package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Ijazah struct {
	ID                          uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	PesertaDidikId              uuid.UUID  `gorm:"column:peserta_didik_id"`
	AnggotaRombelId             uuid.UUID  `gorm:"column:anggota_rombel_id"`
	NomorIjazah                 string     `gorm:"column:nomor_ijazah;unique;not null"`
	TempatIjazah                string     `gorm:"column:tempat_ijazah"`
	TanggalIjazah               *time.Time `gorm:"column:tanggal_ijazah"`
	ProgramKeahlian             string     `gorm:"column:program_keahlian"`
	PaketKeahlian               string     `gorm:"column:paket_keahlian"`
	SekolahID                   uuid.UUID  `gorm:"column:sekolah_id"`
	NamaOrtuWali                string     `gorm:"column:nama_ortu_wali"`
	SekolahPenyelenggaraUjianUS string     `gorm:"column:sekolah_penyelenggara_ujian_us"`
	SekolahPenyelenggaraUjianUN string     `gorm:"column:sekolah_penyelenggara_ujian_un"`
	BlockexplorerUrl            string     `gorm:"column:blockexplorer_url"`
	CidUrl                      string     `gorm:"column:cid_url"`
	TahunAjaranId               string     `gorm:"column:tahun_ajaran_id"`
	Status                      int32      `gorm:"column:status"` // DRAFT, PENDING, ACTIVE, REVOKED
	Nis                         string     `gorm:"type:varchar(20);unique;not null"`
	NISN                        string     `gorm:"type:varchar(20);unique;not null"`
	CreatedAt                   time.Time
	UpdatedAt                   time.Time
	DeletedAt                   gorm.DeletedAt `gorm:"index"`

	// NPSN                        string         `gorm:"type:varchar(15);not null"`
	// AsalSekolah                 string    `gorm:"type:varchar(200);not null"`
	// KabupatenKota               string    `gorm:"type:varchar(100);not null"`
	// Provinsi                    string    `gorm:"type:varchar(100);not null"`
	// Nama                        string    `gorm:"type:varchar(200);not null"`
	// TempatLahir                 string    `gorm:"type:varchar(100);not null"`
	// TanggalLahir                string    `gorm:"type:date;not null"`
	// Relasi
	PesertaDidik PesertaDidik `gorm:"foreignKey:PesertaDidikId;references:PesertaDidikId"`
	// RombonganBelajar RombonganBelajar `gorm:"foreignKey:RombonganBelajarId;references:RombonganBelajarId"`
	// NilaiUjianSekolah NIlaiRerataAkhir
}

func (Ijazah) TableName() string {
	return "ijazah"
}

type TranskripNilai struct {
	gorm.Model
	NamaMapel          string
	PesertaDidikId     uuid.UUID
	RombonganBelajarId uuid.UUID
	NilaRata           uint32
	TahunAjaranId      string `gorm:"column:tahun_ajaran_id"`
}
