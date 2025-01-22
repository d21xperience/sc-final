package models

import (
	"github.com/google/uuid"
)

type NilaiAkhir struct {
	IDNilaiAkhir    uuid.UUID  `gorm:"column:id_nilai_akhir;primaryKey"` // Primary key
	AnggotaRombelID *uuid.UUID `gorm:"column:anggota_rombel_id"`         // Foreign key ke tabel anggota_rombel
	MataPelajaranID *int32     `gorm:"column:mata_pelajaran_id"`         // Foreign key ke tabel mata_pelajaran
	SemesterID      string     `gorm:"column:semester_id"`               // Semester
	NilaiPeng       *int32     `gorm:"column:nilai_peng"`                // Nilai Pengetahuan
	PredikatPeng    string     `gorm:"column:predikat_peng"`             // Predikat Pengetahuan
	NilaiKet        *int32     `gorm:"column:nilai_ket"`                 // Nilai Keterampilan
	PredikatKet     string     `gorm:"column:predikat_ket"`              // Predikat Keterampilan
	NilaiSik        *int32     `gorm:"column:nilai_sik"`                 // Nilai Sikap
	PredikatSik     string     `gorm:"column:predikat_sik"`              // Predikat Sikap
	NilaiSikSos     *int32     `gorm:"column:nilai_siksos"`              // Nilai Sikap Sosial
	PredikatSikSos  string     `gorm:"column:predikat_siksos"`           // Predikat Sikap Sosial
	PesertaDidikID  *uuid.UUID `gorm:"column:peserta_didik_id"`          // Foreign key ke tabel peserta_didik
	IDMinat         string     `gorm:"column:id_minat"`                  // ID Minat
	Semester        *int32     `gorm:"column:semester"`                  // Semester
	// Relasi
	// AnggotaRombel RombelAnggota `gorm:"foreignKey:AnggotaRombelID"`
	// MataPelajaran MataPelajaran `gorm:"foreignKey:MataPelajaranID"`
	// PesertaDidik  PesertaDidik  `gorm:"foreignKey:PesertaDidikID"`
}

// type NIlaiRerataAkhir struct {
// 	NilaiRerataAkhirID uuid.UUID `gorm:"column:nilai_rerata_akhir_id"`
// 	IDNilaiAkhir       uuid.UUID `gorm:"column:id_nilai_akhir;foreignKey:id_nilai_akhir"` // Foreign key ke tabel nilai_akhir
// 	NilaiAkhir         []NilaiAkhir
// }
