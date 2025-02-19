package models

type RombonganBelajar struct {
	RombonganBelajarId  string  `gorm:"column:rombongan_belajar_id;primaryKey"`
	SekolahId           string  `gorm:"column:sekolah_id"`
	SemesterId          string  `gorm:"column:semester_id"`
	JurusanId           string  `gorm:"column:jurusan_id"`
	PtkID               string  `gorm:"column:ptk_id"`
	NmKelas             string  `gorm:"column:nm_kelas"`
	TingkatPendidikanId int32   `gorm:"column:tingkat_pendidikan_id"`
	JenisRombel         int32   `gorm:"column:jenis_rombel"`
	NamaJurusanSp       string  `gorm:"column:nama_jurusan_sp"`
	JurusanSpId         *string `gorm:"column:jurusan_sp_id"` // Nullable field
	KurikulumId         int32   `gorm:"column:kurikulum_id"`

	PTK TabelPTK `gorm:"foreignKey:PtkID;references:PtkID"`
	// ERROR: relation "jurusan" does not exist (SQLSTATE 42P01)
	// Jurusan Jurusan  `gorm:"foreignKey:JurusanId;references:JurusanId"`
}

type RombelAnggota struct {
	AnggotaRombelId    string `gorm:"column:anggota_rombel_id"` // UUID
	PesertaDidikId     string `gorm:"column:peserta_didik_id"`  // UUID
	RombonganBelajarId string `gorm:"column:rombongan_belajar_id"`
	SemesterId         string `gorm:"column:semester_id"`

	PesertaDidik     PesertaDidik     `gorm:"foreignKey:PesertaDidikId;references:PesertaDidikId"`
	RombonganBelajar RombonganBelajar `gorm:"foreignKey:RombonganBelajarId;references:RombonganBelajarId"`
}

func (RombonganBelajar) TableName() string {
	return "tabel_kelas"
}
func (RombelAnggota) TableName() string {
	return "tabel_anggotakelas"
}
