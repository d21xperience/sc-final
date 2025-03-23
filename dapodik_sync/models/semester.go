package models

// Model untuk tabel semester
type Semester struct {
	SemesterID    string `gorm:"column:semester_id; type:CHAR(5); not null"`
	TahunAjaranID int    `gorm:"column:tahun_ajaran_id; type:NUMERIC(4,0); not null"`
	Nama          string `gorm:"column:nama; type:VARCHAR(20); not null"`
	Semester      int    `gorm:"column:semester; type:NUMERIC(1,0); not null"`
}
