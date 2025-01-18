package models

type RombonganBelajar struct {
	RombonganBelajarID  string       `json:"rombongan_belajar_id"`
	SekolahID           Sekolah      `json:"sekolah_id"`
	SemesterID          Semester `json:"semester_id"`
	JurusanID           string       `json:"jurusan_id"`
	PtkTerdaftarID      PTKTerdaftar `json:"ptk_id"`
	NmKelas             string       `json:"nm_kelas"`
	TingkatPendidikanID string       `json:"tingkat_pendidikan_id"`
	JenisRombel         string       `json:"jenis_rombel"`
	NamaJurusanSP       string       `json:"nama_jurusan_sp"`
	JurusanSpID         *string      `json:"jurusan_sp_id"` // Nullable field
	KurikulumID         string       `json:"kurikulum_id"`
}
