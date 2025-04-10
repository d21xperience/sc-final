package models

// type Jurusan struct {
// 	JurusanId           string  `gorm:"column:jurusan_id;primaryKey"` // Primary key
// 	NamaJurusan         string  `gorm:"column:nama_jurusan"`          // Nama jurusan
// 	UntukSMA            int     `gorm:"column:untuk_sma"`             // Indikator untuk SMA
// 	UntukSMK            int     `gorm:"column:untuk_smk"`             // Indikator untuk SMK
// 	UntukPT             int     `gorm:"column:untuk_pt"`              // Indikator untuk PT
// 	UntukSLB            int     `gorm:"column:untuk_slb"`             // Indikator untuk SLB
// 	UntukSMKLB          int     `gorm:"column:untuk_smklb"`           // Indikator untuk SMKLB
// 	JenjangPendidikanID *int    `gorm:"column:jenjang_pendidikan_id"` // ID jenjang pendidikan
// 	JurusanInduk        *string `gorm:"column:jurusan_induk"`         // Foreign key ke jurusan (self-referencing)
// 	LevelBidangID       string  `gorm:"column:level_bidang_id"`       // Level bidang
// 	// gorm.Model
// 	CreateDate  time.Time  `gorm:"column:create_date"`  // Tanggal pembuatan
// 	LastUpdate  time.Time  `gorm:"column:last_update"`  // Tanggal update terakhir
// 	ExpiredDate *time.Time `gorm:"column:expired_date"` // Tanggal kedaluwarsa
// 	LastSync    time.Time  `gorm:"column:last_sync"`    // Tanggal sinkronisasi terakhir
// 	// Relasi
// 	// ParentJurusan *Jurusan  `gorm:"foreignKey:JurusanInduk"` // Relasi self-referencing ke jurusan induk
// 	// ChildJurusan  []Jurusan `gorm:"foreignKey:JurusanInduk"` // Relasi ke anak-anak jurusan
// }

// func (Jurusan) TableName() string {
// 	return "jurusan"
// }

// type TabelKurikulum struct {
// 	KurikulumID         int16      `gorm:"column:kurikulum_id;primaryKey;not null"`                 // SMALLINT, primary key
// 	NamaKurikulum       string     `gorm:"column:nama_kurikulum;type:varchar(120);not null"`        // VARCHAR(120), tidak boleh null
// 	MulaiBerlaku        time.Time  `gorm:"column:mulai_berlaku;not null"`                           // DATE, tidak boleh null
// 	SistemSKS           int8       `gorm:"column:sistem_sks;type:numeric(1,0);not null"`            // NUMERIC(1,0), tidak boleh null
// 	TotalSKS            int16      `gorm:"column:total_sks;type:numeric(3,0);not null"`             // NUMERIC(3,0), tidak boleh null
// 	JenjangPendidikanID int8       `gorm:"column:jenjang_pendidikan_id;type:numeric(2,0);not null"` // NUMERIC(2,0), tidak boleh null
// 	JurusanID           *string    `gorm:"column:jurusan_id;type:varchar(25);default:null"`         // VARCHAR(25), bisa null
// 	CreateDate          time.Time  `gorm:"column:create_date;not null"`                             // TIMESTAMP, tidak boleh null
// 	LastUpdate          time.Time  `gorm:"column:last_update;not null"`                             // TIMESTAMP, tidak boleh null
// 	ExpiredDate         *time.Time `gorm:"column:expired_date;default:null"`                        // TIMESTAMP, bisa null
// 	LastSync            time.Time  `gorm:"column:last_sync;not null"`                               // TIMESTAMP, tidak boleh null
// }



type TahunAjaran struct {
	TahunAjaranID  uint32 `gorm:"column:tahun_ajaran_id"`
	Nama           string `gorm:"column:nama"`
	PeriodeAktif   string `gorm:"column:periode_aktif"`
	TanggalMulai   string `gorm:"column:tanggal_mulai"`
	TanggalSelesai string `gorm:"column:tanggal_selesai"`
}

type Semester struct {
	SemesterID     string `gorm:"column:semester_id"`
	TahunAjaranID  uint32 `gorm:"column:tahun_ajaran_id"`
	Nama           string `gorm:"column:nama"`
	Semester       int32  `gorm:"column:semester"`
	PeriodeAktif   int32  `gorm:"column:periode_aktif"`
	TanggalMulai   string `gorm:"column:tanggal_mulai"`
	TanggalSelesai string `gorm:"column:tanggal_selesai"`
	// Relasi ke TahunAjaran
	TahunAjaran TahunAjaran `gorm:"foreignKey:TahunAjaranID;references:TahunAjaranID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
