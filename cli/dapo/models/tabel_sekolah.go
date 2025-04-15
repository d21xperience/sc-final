package models

type SekolahResponse struct {
	Results int     `json:"results"`
	ID      string  `json:"id"`
	Start   int     `json:"start"`
	Limit   int     `json:"limit"`
	Rows    Sekolah `json:"rows"`
}

type Sekolah struct {
	SekolahID           string `json:"sekolah_id"`
	Nama                string `json:"nama"`
	NSS                 string `json:"nss"`
	NPSN                string `json:"npsn"`
	BentukPendidikanID  int    `json:"bentuk_pendidikan_id"`
	BentukPendidikanStr string `json:"bentuk_pendidikan_id_str"`
	StatusSekolah       string `json:"status_sekolah"`
	StatusSekolahStr    string `json:"status_sekolah_str"`
	AlamatJalan         string `json:"alamat_jalan"`
	RT                  string `json:"rt"`
	RW                  string `json:"rw"`
	KodeWilayah         string `json:"kode_wilayah"`
	KodePos             string `json:"kode_pos"`
	NomorTelepon        string `json:"nomor_telepon"`
	NomorFax            string `json:"nomor_fax"`
	Email               string `json:"email"`
	Website             string `json:"website"`
	IsSKS               bool   `json:"is_sks"`
	Lintang             string `json:"lintang"`
	Bujur               string `json:"bujur"`
	Dusun               string `json:"dusun"`
	DesaKelurahan       string `json:"desa_kelurahan"`
	Kecamatan           string `json:"kecamatan"`
	KabupatenKota       string `json:"kabupaten_kota"`
	Provinsi            string `json:"provinsi"`
}
