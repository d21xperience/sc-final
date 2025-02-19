package models

type PesertaDidik struct {
	PesertaDidikId  string  `gorm:"column:peserta_didik_id;primaryKey"` // STRING
	Nis             string  `gorm:"column:nis"`                         // String
	Nisn            string  `gorm:"column:nisn"`                        // String
	NmSiswa         string  `gorm:"column:nm_siswa"`                    // String
	TempatLahir     string  `gorm:"column:tempat_lahir"`                // String
	TanggalLahir    string  `gorm:"column:tanggal_lahir"`               // String (format tanggal, bisa diubah ke time.Time jika perlu)
	JenisKelamin    string  `gorm:"column:jenis_kelamin"`               // String
	Agama           string  `gorm:"column:agama"`                       // String
	AlamatSiswa     *string `gorm:"column:alamat_siswa"`                // Nullable string
	TeleponSiswa    string  `gorm:"column:telepon_siswa"`               // String
	DiterimaTanggal string  `gorm:"column:diterima_tanggal"`            // String (format tanggal, bisa diubah ke time.Time jika perlu)
	NmAyah          string  `gorm:"column:nm_ayah"`                     // String
	NmIbu           string  `gorm:"column:nm_ibu"`                      // String
	PekerjaanAyah   string  `gorm:"column:pekerjaan_ayah"`              // String
	PekerjaanIbu    string  `gorm:"column:pekerjaan_ibu"`               // String
	NmWali          *string `gorm:"column:nm_wali"`                     // Nullable string
	PekerjaanWali   *string `gorm:"column:pekerjaan_wali"`              // Nullable string
}

type PesertaDidikPelengkap struct {
	PelengkapSiswaId string       `gorm:"column:pelengkap_siswa_id;primaryKey"` // UUID
	PesertaDidikId   *string      `gorm:"column:peserta_didik_id"`              // UUID
	StatusDalamKel   *string      `gorm:"column:status_dalam_kel"`              // Nullable string
	AnakKe           *string      `gorm:"column:anak_ke"`                       // Integer
	SekolahAsal      string       `gorm:"column:sekolah_asal"`                  // Non-nullable string
	DiterimaKelas    *string      `gorm:"column:diterima_kelas"`                // Nullable string
	AlamatOrtu       *string      `gorm:"column:alamat_ortu"`                   // Nullable string
	TeleponOrtu      *string      `gorm:"column:telepon_ortu"`                  // Nullable string
	AlamatWali       *string      `gorm:"column:alamat_wali"`                   // Nullable string
	TeleponWali      *string      `gorm:"column:telepon_wali"`                  // Nullable string
	FotoSiswa        *string      `gorm:"column:foto_siswa"`                    // Nullable string
	PesertaDidik     PesertaDidik `gorm:"foreignKey:PesertaDidikId;references:PesertaDidikId"`
}

func (PesertaDidik) TableName() string {
	return "tabel_siswa"
}
func (PesertaDidikPelengkap) TableName() string {
	return "siswa_pelengkap"
}
