package models

type PesertaDidik struct {
	PesertaDidikID  string  `json:"peserta_didik_id"` // UUID
	NIS             string  `json:"nis"`              // String
	NISN            string  `json:"nisn"`             // String
	NamaSiswa       string  `json:"nm_siswa"`         // String
	TempatLahir     string  `json:"tempat_lahir"`     // String
	TanggalLahir    string  `json:"tanggal_lahir"`    // String (format tanggal, bisa diubah ke time.Time jika perlu)
	JenisKelamin    string  `json:"jenis_kelamin"`    // String
	Agama           string  `json:"agama"`            // String
	AlamatSiswa     *string `json:"alamat_siswa"`     // Nullable string
	TeleponSiswa    string  `json:"telepon_siswa"`    // String
	DiterimaTanggal string  `json:"diterima_tanggal"` // String (format tanggal, bisa diubah ke time.Time jika perlu)
	NamaAyah        string  `json:"nm_ayah"`          // String
	NamaIbu         string  `json:"nm_ibu"`           // String
	PekerjaanAyah   string  `json:"pekerjaan_ayah"`   // String
	PekerjaanIbu    string  `json:"pekerjaan_ibu"`    // String
	NamaWali        *string `json:"nm_wali"`          // Nullable string
	PekerjaanWali   *string `json:"pekerjaan_wali"`   // Nullable string
}

type PesertaDidikPelengkap struct {
	PelengkapSiswaID string       `json:"pelengkap_siswa_id"` // UUID
	PesertaDidikID   PesertaDidik `json:"peserta_didik_id"`   // UUID
	StatusDalamKel   *string      `json:"status_dalam_kel"`   // Nullable string
	AnakKe           *string      `json:"anak_ke"`            // Integer
	SekolahAsal      string       `json:"sekolah_asal"`       // Non-nullable string
	DiterimaKelas    *string      `json:"diterima_kelas"`     // Nullable string
	AlamatOrtu       *string      `json:"alamat_ortu"`        // Nullable string
	TeleponOrtu      *string      `json:"telepon_ortu"`       // Nullable string
	AlamatWali       *string      `json:"alamat_wali"`        // Nullable string
	TeleponWali      *string      `json:"telepon_wali"`       // Nullable string
	FotoSiswa        *string      `json:"foto_siswa"`         // Nullable string
}
