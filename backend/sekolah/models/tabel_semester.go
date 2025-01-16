package models

type DataSemester struct {
	SemesterID     string      `json:"semester_id"`
	TahunAjaranID  TahunAjaran `json:"tahun_ajaran_id"`
	NamaSemester   string      `json:"nama_semester"`
	Semester       string      `json:"semester"`
	PeriodeAktif   string      `json:"periode_aktif"`
	TanggalMulai   string      `json:"tanggal_mulai"`
	TanggalSelesai string      `json:"tanggal_selesai"`
}

type TahunAjaran struct {
	TahunAjaranID string `json:"tahun_ajaran_id"`
	Nama          string `json:"nama"`
	PeriodeAktif  string `json:"periode_aktif"`
}
