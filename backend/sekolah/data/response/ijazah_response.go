package response

type IjazahResponse struct {
	NomorIjazah string  `json:"nomor_ijazah"`
	AsalSekolah string  `json:"asal_sekolah"`
	Nama        string  `json:"nama"`
	TglLahir    string  `json:"tgl_lahir"`
	TptLahir    string  `json:"tpt_lahir"`
	Nis         string  `json:"nis"`
	Nisn        string  `json:"nisn"`
	NamaWali    string  `json:"nama_wali"`
	TahunLulus  string  `json:"tahun_lulus"`
	RerataNilai float32 `json:"rerata_nilai"`
	// valid       bool    `gorm:"type:bool"`
}
