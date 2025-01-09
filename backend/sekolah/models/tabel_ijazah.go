package models

import (
	"gorm.io/gorm"
)

type Ijazah struct {
	gorm.Model
	ID                          string
	PesertaDidikID              string
	ProgramKeahlian             string
	PaketKeahlian               string
	NPSN                        string
	KabupatenKota               string
	Provinsi                    string
	Nama                        string
	TempatLahir                 string
	TanggalLahir                string
	NamaOrtuWali                string
	Nis                         string
	NISN                        string
	SekolahPenyelenggaraUjianUS string
	SekolahPenyelenggaraUjianUN string
	AsalSekolah                 string
	NomorIjazah                 string
	TempatIjazah                string
	TanggalIjazah               string
	NoIjazah                    string
	NilaiUjianSekolah           TabelNilaiAkhir
}
