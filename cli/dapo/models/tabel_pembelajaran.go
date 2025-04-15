package models

type Pembelajaran struct {
	PembelajaranId     string
	RombonganBelajarId string
	MataPelajaranId    int
	SemesterId         string
	PtkTerdaftarId     string
	StatusDiKurikulum  int
	NamaMataPelajaran  string
	IndukPembelajaran  string
	IsDapo             int
}
