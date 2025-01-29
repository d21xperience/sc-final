package models

type SekolahIndonesia struct {
	SekolahIdEnkrip string `gorm:"column:sekolah_id_enkrip;not null;unique"`
	Kecamatan       string `gorm:"column:kecamatan"`
	Kabupaten       string `gorm:"column:kabupaten"`
	Propinsi        string `gorm:"column:propinsi"`
	KodeKecamatan   string `gorm:"column:kode_kecamatan"`
	KodeKab         string `gorm:"column:kode_kab"`
	KodeProp        string `gorm:"column:kode_prop"`
	NamaSekolah     string `gorm:"column:nama_sekolah"`
	Npsn            string `gorm:"type:varchar(8);not null;unique"`
	AlamatJalan     string `gorm:"column:alamat_jalan"`
	Status          string `gorm:"column:status"`
}

// Menentukan nama tabel kustom
func (SekolahIndonesia) TableName() string {
	return "sekolah_indonesia"
}
