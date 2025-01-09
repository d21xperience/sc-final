package request

type CreateTagsRequest struct {
	// Name string `validate:"required,min=1,max=200" json:"name"`
	// NomorIjazah string `json:"text;primaryKey"`
	// PesertaDidikID uuid.UUID `json:"uuid"`
	AsalSekolah string  `json:"asal_sekolah"`
	Nama        string  `json:"nama"`
	TptLahir    string  `json:"tpt_lahir"`
	TglLahir    string  `json:"tgl_lahir"`
	NIS         string  `json:"nis"`
	NISN        string  `json:"nisn"`
	NamaWali    string  `json:"nama_wali"`
	TahunLulus  string  `json:"thn_lulus"`
	RerataNilai float32 `json:"rerata_nilai"`
	// Valid       bool    `json:"date"`
	// ImageURL string `json:"date"` //url untuk menyimpan data foto ke IPFS
}
