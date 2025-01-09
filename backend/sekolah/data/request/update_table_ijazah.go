package request

type UpdateTableIjazah struct {
	NomorIjazah string  `gorm:"type:text"`
	AsalSekolah string  `gorm:"type:text"`
	Nama        string  `gorm:"type:string"`
	TglLahir    string  `gorm:"type:date"`
	TptLahir    string  `gorm:"type:string"`
	Nis         string  `gorm:"type:string"`
	Nisn        string  `gorm:"type:string"`
	NamaWali    string  `gorm:"type:string"`
	TahunLulus  string  `gorm:"type:date"`
	RerataNilai float32 `gorm:"type:date"`
	// valid       bool    `gorm:"type:bool"`
}
