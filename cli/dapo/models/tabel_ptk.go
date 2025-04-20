package models

type PTKTerdaftarResponse struct {
	Results int            `json:"results"`
	ID      string         `json:"id"`
	Start   int            `json:"start"`
	Limit   int            `json:"limit"`
	Rows    []PTKTerdaftar `json:"rows"`
}

type PTKTerdaftar struct {
	TahunAjaranID           string              `json:"tahun_ajaran_id"`
	PTKTerdaftarID          string              `json:"ptk_terdaftar_id"`
	PTKID                   string              `json:"ptk_id"`
	PTKInduk                string              `json:"ptk_induk"`
	TanggalSuratTugas       string              `json:"tanggal_surat_tugas"`
	Nama                    string              `json:"nama"`
	JenisKelamin            string              `json:"jenis_kelamin"`
	TempatLahir             string              `json:"tempat_lahir"`
	TanggalLahir            string              `json:"tanggal_lahir"`
	AgamaID                 int                 `json:"agama_id"`
	AgamaIDStr              string              `json:"agama_id_str"`
	NUPTK                   string              `json:"nuptk"`
	NIK                     string              `json:"nik"`
	JenisPTKID              int                 `json:"jenis_ptk_id"`
	JenisPTKIDStr           string              `json:"jenis_ptk_id_str"`
	StatusKepegawaianID     int                 `json:"status_kepegawaian_id"`
	StatusKepegawaianIDStr  string              `json:"status_kepegawaian_id_str"`
	NIP                     string              `json:"nip"`
	PendidikanTerakhir      string              `json:"pendidikan_terakhir"`
	BidangStudiTerakhir     string              `json:"bidang_studi_terakhir"`
	PangkatGolonganTerakhir string              `json:"pangkat_golongan_terakhir"`
	RiwayatPendidikanFormal []RiwayatPendidikan `json:"rwy_pend_formal"`
	RiwayatKepangkatan      []interface{}       `json:"rwy_kepangkatan"` // Kosong di data, bisa diubah sesuai struktur nanti
}

type RiwayatPendidikan struct {
	RiwayatPendidikanFormalID string  `json:"riwayat_pendidikan_formal_id"`
	SatuanPendidikanFormal    string  `json:"satuan_pendidikan_formal"`
	Fakultas                  string  `json:"fakultas"`
	Kependidikan              string  `json:"kependidikan"`
	TahunMasuk                string  `json:"tahun_masuk"`
	TahunLulus                string  `json:"tahun_lulus"`
	NIM                       string  `json:"nim"`
	StatusKuliah              string  `json:"status_kuliah"`
	Semester                  string  `json:"semester"`
	IPK                       string  `json:"ipk"`
	Prodi                     *string `json:"prodi"`     // nullable
	IDRegPD                   *string `json:"id_reg_pd"` // nullable
	BidangStudiIDStr          string  `json:"bidang_studi_id_str"`
	JenjangPendidikanIDStr    string  `json:"jenjang_pendidikan_id_str"`
	GelarAkademikIDStr        string  `json:"gelar_akademik_id_str"`
}
