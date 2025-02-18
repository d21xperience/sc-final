package models

type RombonganBelajar struct {
	RombonganBelajarId  string  `json:"rombongan_belajar_id"`
	SekolahId           string  `json:"sekolah_id"`
	SemesterId          string  `json:"semester_id"`
	JurusanId           string  `json:"jurusan_id"`
	PtkId               string  `json:"ptk_id"`
	NmKelas             string  `json:"nm_kelas"`
	TingkatPendidikanId int32   `json:"tingkat_pendidikan_id"`
	JenisRombel         int32   `json:"jenis_rombel"`
	NamaJurusanSp       string  `json:"nama_jurusan_sp"`
	JurusanSpId         *string `gorm:"column:jurusan_sp_id" json:"jurusan_sp_id"` // Nullable field
	KurikulumId         int32   `json:"kurikulum_id"`
}

type RombelAnggota struct {
	AnggotaRombelId    string `json:"anggota_rombel_id"` // UUID
	PesertaDidikId     string `json:"peserta_didik_id"`  // UUID
	RombonganBelajarId string `json:"rombongan_belajar_id"`
	SemesterId         string `json:"semester_id"`
}

// type TabelKelas struct {
// 	RombonganBelajarID  uuid.UUID  `gorm:"type:uuid;primaryKey"`
// 	SekolahID           uuid.UUID  `gorm:"type:uuid;not null"`
// 	SemesterID          string     `gorm:"type:char(5);not null"`
// 	JurusanID           *string    `gorm:"type:varchar(25);default:NULL"`
// 	PtkID               *uuid.UUID `gorm:"type:uuid;default:NULL"`
// 	NmKelas             *string    `gorm:"type:varchar(30);default:NULL"`
// 	TingkatPendidikanID *int       `gorm:"type:numeric(2,0);default:NULL"`
// 	JenisRombel         *int       `gorm:"type:numeric(2,0);default:NULL"`
// 	NamaJurusanSP       *string    `gorm:"type:varchar(100);default:NULL"`
// 	JurusanSpID         *uuid.UUID `gorm:"type:uuid;default:NULL"`
// 	KurikulumID         int16      `gorm:"type:smallint;not null"`

// 	// Foreign Keys
// 	Jurusan           Jurusan           `gorm:"foreignKey:JurusanID;references:JurusanID"`
// 	Kurikulum         Kurikulum         `gorm:"foreignKey:KurikulumID;references:KurikulumID"`
// 	Semester          Semester          `gorm:"foreignKey:SemesterID;references:SemesterID"`
// 	TingkatPendidikan TingkatPendidikan `gorm:"foreignKey:TingkatPendidikanID;references:TingkatPendidikanID"`
// 	Ptk               Ptk               `gorm:"foreignKey:PtkID;references:PtkID"`
// }
