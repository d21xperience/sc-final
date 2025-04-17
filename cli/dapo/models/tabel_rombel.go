package models

type RombelResponse struct {
	Results int                `json:"results"`
	ID      string             `json:"id"`
	Start   int                `json:"start"`
	Limit   int                `json:"limit"`
	Rows    []RombonganBelajar `json:"rows"`
}

type RombonganBelajar struct {
	RombonganBelajarID   string          `json:"rombongan_belajar_id"`
	Nama                 string          `json:"nama"`
	TingkatPendidikanID  string          `json:"tingkat_pendidikan_id"`
	TingkatPendidikanStr string          `json:"tingkat_pendidikan_id_str"`
	SemesterID           string          `json:"semester_id"`
	JenisRombel          string          `json:"jenis_rombel"`
	JenisRombelStr       string          `json:"jenis_rombel_str"`
	KurikulumID          int             `json:"kurikulum_id"`
	KurikulumIDStr       string          `json:"kurikulum_id_str"`
	IDRuang              string          `json:"id_ruang"`
	IDRuangStr           string          `json:"id_ruang_str"`
	MovingClass          string          `json:"moving_class"`
	PTKID                string          `json:"ptk_id"`
	PTKIDStr             string          `json:"ptk_id_str"`
	JurusanID            string          `json:"jurusan_id"`
	JurusanIDStr         string          `json:"jurusan_id_str"`
	AnggotaRombel        []AnggotaRombel `json:"anggota_rombel"`
	Pembelajaran         []Pembelajaran  `json:"pembelajaran"`
}

type AnggotaRombel struct {
	AnggotaRombelID     string `json:"anggota_rombel_id"`
	PesertaDidikID      string `json:"peserta_didik_id"`
	JenisPendaftaranID  string `json:"jenis_pendaftaran_id"`
	JenisPendaftaranStr string `json:"jenis_pendaftaran_id_str"`
}
type Pembelajaran struct {
	PembelajaranID       string  `gorm:"type:uuid;primaryKey" json:"pembelajaran_id"`
	MataPelajaranID      int64   `gorm:"column:mata_pelajaran_id" json:"mata_pelajaran_id"`
	MataPelajaranIDStr   string  `gorm:"column:mata_pelajaran_id_str" json:"mata_pelajaran_id_str"`
	PTKTerdaftarID       string  `gorm:"type:uuid" json:"ptk_terdaftar_id"`
	PTKID                string  `gorm:"type:uuid" json:"ptk_id"`
	NamaMataPelajaran    string  `gorm:"column:nama_mata_pelajaran" json:"nama_mata_pelajaran"`
	IndukPembelajaranID  *string `gorm:"type:uuid" json:"induk_pembelajaran_id"` // Nullable UUID
	JamMengajarPerMinggu string  `gorm:"column:jam_mengajar_per_minggu" json:"jam_mengajar_per_minggu"`
	StatusDiKurikulum    string  `gorm:"column:status_di_kurikulum" json:"status_di_kurikulum"`
	StatusDiKurikulumStr string  `gorm:"column:status_di_kurikulum_str" json:"status_di_kurikulum_str"`
}
