package models

type TabelPTK struct {
	PTKID             string  `json:"ptk_id"`              // UUID
	Nama              string  `json:"nama"`                // String
	NIP               *string `json:"nip"`                 // Nullable string
	JenisPTKID        string  `json:"jenis_ptk_id"`        // String
	JenisKelamin      string  `json:"jenis_kelamin"`       // String
	TempatLahir       string  `json:"tempat_lahir"`        // String
	TanggalLahir      string  `json:"tanggal_lahir"`       // String (format tanggal, bisa diubah ke time.Time jika perlu)
	NUPTK             *string `json:"nuptk"`               // Nullable string
	AlamatJalan       string  `json:"alamat_jalan"`        // String
	StatusKeaktifanID string  `json:"status_keaktifan_id"` // String
}

type PTKPelengkap struct {
	PTKPelengkapID string  `json:"ptk_pelengkap_id"` // UUID
	PTKID          string  `json:"ptk_id"`           // UUID
	GelarDepan     *string `json:"gelar_depan"`      // Nullable string
	GelarBelakang  *string `json:"gelar_belakang"`   // Nullable string
}

type PTKTerdaftar struct {
	PtkTerdaftarID string  `json:"ptk_terdaftar_id"`
	PtkID          string  `json:"ptk_id"`
	TahunAjaranID  string  `json:"tahun_ajaran_id"`
	JenisKeluarID  *string `json:"jenis_keluar_id"` // Nullable field
}
