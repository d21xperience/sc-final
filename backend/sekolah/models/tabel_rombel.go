package models

type RombonganBelajar struct {
	RombonganBelajarId  string  `json:"rombongan_belajar_id"`
	SekolahId           string  `json:"sekolah_id"`
	SemesterId          string  `json:"semester_id"`
	JurusanId           string  `json:"jurusan_id"`
	PtkId               string  `json:"ptk_id"`
	NmKelas             string  `json:"nm_kelas"`
	TingkatPendidikanId int32   `json:"tingkat_pendidikan_id"`
	JenisRombel         int32   `json:"jenis_rombel"`
	NamaJurusanSp       string  `json:"nama_jurusan_sp"`
	JurusanSpId         *string `gorm:"column:jurusan_sp_id" json:"jurusan_sp_id"` // Nullable field
	KurikulumId         int32   `json:"kurikulum_id"`
}

type RombelAnggota struct {
	AnggotaRombelId    string `json:"anggota_rombel_id"` // UUID
	PesertaDidikId     string `json:"peserta_didik_id"`  // UUID
	RombonganBelajarId string `json:"rombongan_belajar_id"`
	SemesterId         string `json:"semester_id"`
}
