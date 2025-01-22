package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Ijazah struct {
	gorm.Model
	PesertaDidikID              *uuid.UUID `gorm:"column:peserta_didik_id"`
	ProgramKeahlian             string
	PaketKeahlian               string
	SekolahID                   *uuid.UUID `gorm:"column:sekolah_id"`
	NPSN                        string
	KabupatenKota               string
	Provinsi                    string
	Nama                        string
	TempatLahir                 string
	TanggalLahir                string
	Nis                         string
	NISN                        string
	NamaOrtuWali                string
	SekolahPenyelenggaraUjianUS string
	SekolahPenyelenggaraUjianUN string
	AsalSekolah                 string
	NomorIjazah                 string
	TempatIjazah                string
	TanggalIjazah               string
	NoIjazah                    string

	// Relasi
	PesertaDidik PesertaDidik `gorm:"foreignKey:PesertaDidikID"`
	// NilaiUjianSekolah NIlaiRerataAkhir
}
