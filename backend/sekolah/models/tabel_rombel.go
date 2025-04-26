package models

import "github.com/google/uuid"

type RombonganBelajar struct {
	RombonganBelajarId  uuid.UUID  `gorm:"column:rombongan_belajar_id;primaryKey"`
	SekolahId           uuid.UUID  `gorm:"column:sekolah_id"`
	SemesterId          string     `gorm:"column:semester_id"`
	JurusanId           string     `gorm:"column:jurusan_id"`
	PtkID               uuid.UUID  `gorm:"column:ptk_id"`
	NmKelas             string     `gorm:"column:nm_kelas"`
	TingkatPendidikanId int32      `gorm:"column:tingkat_pendidikan_id"`
	JenisRombel         int32      `gorm:"column:jenis_rombel"`
	NamaJurusanSp       string     `gorm:"column:nama_jurusan_sp"`
	JurusanSpId         *uuid.UUID `gorm:"column:jurusan_sp_id"` // Nullable field
	KurikulumId         int32      `gorm:"column:kurikulum_id"`

	PTK TabelPTK `gorm:"foreignKey:PtkID;references:PtkID"`
	// ERROR: relation "jurusan" does not exist (SQLSTATE 42P01)
	// Foreign Key ke Jurusan
	Jurusan Jurusan `gorm:"foreignKey:JurusanId;references:JurusanID"`

	// Foreign Key ke Kurikulum
	Kurikulum Kurikulum `gorm:"foreignKey:KurikulumId;references:KurikulumID"`

	// Foreign Key ke Tingkat Pendidikan
	TingkatPendidikan TingkatPendidikan `gorm:"foreignKey:TingkatPendidikanId;references:TingkatPendidikanID"`
	Pembelajaran      []Pembelajaran
	AnggotaKelas      []RombelAnggota
}

type RombelAnggota struct {
	AnggotaRombelId    uuid.UUID        `gorm:"column:anggota_rombel_id"` // UUID
	PesertaDidikId     uuid.UUID        `gorm:"column:peserta_didik_id"`  // UUID
	RombonganBelajarId uuid.UUID        `gorm:"column:rombongan_belajar_id"`
	SemesterId         string           `gorm:"column:semester_id"`
	StatusKeaktifan    uint32           `gorm:"column:status_keaktifan"`
	PesertaDidik       PesertaDidik     `gorm:"foreignKey:PesertaDidikId;references:PesertaDidikId"`
	RombonganBelajar   RombonganBelajar `gorm:"foreignKey:RombonganBelajarId;references:RombonganBelajarId"`
	NilaiAkhir         []NilaiAkhir     `gorm:"foreignKey:AnggotaRombelId;references:AnggotaRombelId"`
}

func (RombonganBelajar) TableName() string {
	return "tabel_kelas"
}
func (RombelAnggota) TableName() string {
	return "tabel_anggotakelas"
}
