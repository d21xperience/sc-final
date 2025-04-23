package models

import (
	"github.com/google/uuid"
)

type NilaiAkhir struct {
	IdNilaiAkhir    uuid.UUID `gorm:"column:id_nilai_akhir;primaryKey"` // Primary key
	AnggotaRombelId uuid.UUID `gorm:"column:anggota_rombel_id"`         // Foreign key ke tabel anggota_rombel
	MataPelajaranId *uint32   `gorm:"column:mata_pelajaran_id"`         // Foreign key ke tabel mata_pelajaran
	SemesterId      string    `gorm:"column:semester_id"`               // Semester
	NilaiPeng       *uint32   `gorm:"column:nilai_peng"`                // Nilai Pengetahuan
	PredikatPeng    string    `gorm:"column:predikat_peng"`             // Predikat Pengetahuan
	NilaiKet        *uint32   `gorm:"column:nilai_ket"`                 // Nilai Keterampilan
	PredikatKet     string    `gorm:"column:predikat_ket"`              // Predikat Keterampilan
	NilaiSik        *uint32   `gorm:"column:nilai_sik"`                 // Nilai Sikap
	PredikatSik     string    `gorm:"column:predikat_sik"`              // Predikat Sikap
	NilaiSikSos     *uint32   `gorm:"column:nilai_siksos"`              // Nilai Sikap Sosial
	PredikatSikSos  string    `gorm:"column:predikat_siksos"`           // Predikat Sikap Sosial
	PesertaDidikId  uuid.UUID `gorm:"column:peserta_didik_id"`          // Foreign key ke tabel peserta_didik
	IDMinat         string    `gorm:"column:id_minat"`                  // ID Minat
	Semester        *uint32   `gorm:"column:semester"`                  // Semester
	// Relasi
	PesertaDidik  PesertaDidik  `gorm:"foreignKey:PesertaDidikId;references:PesertaDidikId"`
	MataPelajaran MataPelajaran `gorm:"foreignKey:MataPelajaranId;references:MataPelajaranID"`
	// AnggotaRombel RombelAnggota `gorm:"foreignKey:AnggotaRombelId;references:AnggotaRombelId"`
}

type NilaiSiswa struct {
	PesertaDidikId     uuid.UUID
	NmSiswa            string
	RombonganBelajarId uuid.UUID
	NmKelas            string
	SemesterId         string
	NilaiAkhir         []NilaiAkhir
}

// type NIlaiRerataAkhir struct {
// 	NilaiRerataAkhirID uuid.UUID `gorm:"column:nilai_rerata_akhir_id"`
// 	IDNilaiAkhir       uuid.UUID `gorm:"column:id_nilai_akhir;foreignKey:id_nilai_akhir"` // Foreign key ke tabel nilai_akhir
// 	NilaiAkhir         []NilaiAkhir
// }

func (NilaiAkhir) TableName() string {
	return "tabel_nilaiakhir"
}
