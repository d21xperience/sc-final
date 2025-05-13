package models

import "time"

type BentukPendidikan struct {
	BentukPendidikanID   uint16  `gorm:"primaryKey;column:bentuk_pendidikan_id"`
	Nama                 string  `gorm:"size:50;not null"`
	JenjangPaud          uint8   `gorm:"not null"`
	JenjangTk            uint8   `gorm:"not null"`
	JenjangSd            uint8   `gorm:"not null"`
	JenjangSmp           uint8   `gorm:"not null"`
	JenjangSma           uint8   `gorm:"not null"`
	JenjangTinggi        uint8   `gorm:"not null"`
	DirektoratPembinaan  *string `gorm:"size:40;default:NULL"`
	Aktif                uint8   `gorm:"not null"`
	FormalitasPendidikan string  `gorm:"type:char(1);not null"`
	// CreateDate          time.Time `gorm:"not null;default:'2020-04-16 09:40:03.422677'"`
	// LastUpdate          time.Time `gorm:"not null;default:'2020-04-16 09:40:03.422677'"`
	// ExpiredDate         *time.Time `gorm:"default:NULL"`
	// LastSync            time.Time `gorm:"not null;default:'1901-01-01 00:00:00'"`
}

func (BentukPendidikan) TableName() string {
	return "ref.bentuk_pendidikan"
}

type JenjangPendidikan struct {
	JenjangPendidikanID uint8  `gorm:"primaryKey;column:jenjang_pendidikan_id"`
	Nama                string `gorm:"size:25;not null"`
	JenjangLembaga      uint8  `gorm:"not null"`
	JenjangOrang        uint8  `gorm:"not null"`
	// CreateDate         time.Time `gorm:"not null;default:'2019-09-10 14:29:56.540627'"`
	// LastUpdate         time.Time `gorm:"not null;default:'2019-09-10 14:29:56.540627'"`
	// ExpiredDate        *time.Time `gorm:"default:NULL"`
	// LastSync           time.Time `gorm:"not null;default:'1901-01-01 00:00:00'"`
}

func (JenjangPendidikan) TableName() string {
	return "ref.jenjang_pendidikan"
}

type TingkatPendidikan struct {
	TingkatPendidikanID uint8  `gorm:"primaryKey;column:tingkat_pendidikan_id"`
	Kode                string `gorm:"size:5;not null"`
	Nama                string `gorm:"size:20;not null"`
	JenjangPendidikanID uint8  `gorm:"not null"`
	// CreateDate          time.Time  `gorm:"not null;default:'2019-09-10 14:29:59.88044'"`
	// LastUpdate          time.Time  `gorm:"not null;default:'2019-09-10 14:29:59.88044'"`
	// ExpiredDate         *time.Time `gorm:"default:NULL"`
	// LastSync            time.Time  `gorm:"not null;default:'1901-01-01 00:00:00'"`
}

func (TingkatPendidikan) TableName() string {
	return "ref.tingkat_pendidikan"
}

// type TahunAjaran struct {
// 	TahunAjaranID  uint16     `gorm:"primaryKey;column:tahun_ajaran_id"`
// 	Nama           string     `gorm:"size:10;not null"`
// 	PeriodeAktif   uint8      `gorm:"not null"`
// 	TanggalMulai   time.Time  `gorm:"not null"`
// 	TanggalSelesai time.Time  `gorm:"not null"`
// 	CreateDate     time.Time  `gorm:"not null;default:'2025-01-19 14:29:59.628052'"`
// 	LastUpdate     time.Time  `gorm:"not null;default:'2025-01-19 14:29:59.628052'"`
// 	ExpiredDate    *time.Time `gorm:"default:NULL"`
// 	LastSync       time.Time  `gorm:"not null;default:'1901-01-01 00:00:00'"`
// }

// func (TahunAjaran) TableName() string {
// 	return "tahun_ajaran"
// }

type StatusKepemilikan struct {
	StatusKepemilikanID uint8  `gorm:"primaryKey;column:status_kepemilikan_id"`
	Nama                string `gorm:"size:20;not null"`
	// CreateDate          time.Time  `gorm:"not null;default:'2019-09-10 14:29:59.426803'"`
	// LastUpdate          time.Time  `gorm:"not null;default:'2019-09-10 14:29:59.426803'"`
	// ExpiredDate         *time.Time `gorm:"default:NULL"`
	// LastSync            time.Time  `gorm:"not null;default:'1901-01-01 00:00:00'"`
}

func (StatusKepemilikan) TableName() string {
	return "ref.status_kepemilikan"
}

type Jurusan struct {
	JurusanID           string  `gorm:"column:jurusan_id;primaryKey"`
	NamaJurusan         string  `gorm:"column:nama_jurusan;size:100;not null"`
	UntukSMA            uint8   `gorm:"column:untuk_sma;not null"`
	UntukSMK            uint8   `gorm:"column:untuk_smk;not null"`
	UntukPT             uint8   `gorm:"column:untuk_pt;not null"`
	UntukSLB            uint8   `gorm:"column:untuk_slb;not null"`
	UntukSMKLB          uint8   `gorm:"column:untuk_smklb;not null"`
	JenjangPendidikanID *uint16 `gorm:"column:jenjang_pendidikan_id"`
	JurusanInduk        *string `gorm:"column:jurusan_induk;size:25"`
	LevelBidangID       string  `gorm:"column:level_bidang_id;size:5;not null"`
	// CreateDate         time.Time  `gorm:"column:create_date;not null"`
	// LastUpdate         time.Time  `gorm:"column:last_update;not null"`
	// ExpiredDate        *time.Time `gorm:"column:expired_date"`
	// LastSync           time.Time  `gorm:"column:last_sync;not null"`

	// Relasi ke Kurikulum
	Kurikulum []Kurikulum `gorm:"foreignKey:JurusanID;references:JurusanID"`
}

func (Jurusan) TableName() string {
	return "ref.jurusan"
}

type Kurikulum struct {
	KurikulumID         uint16    `gorm:"column:kurikulum_id;primaryKey"`
	NamaKurikulum       string    `gorm:"column:nama_kurikulum;size:120;not null"`
	MulaiBerlaku        time.Time `gorm:"column:mulai_berlaku;not null"`
	SistemSKS           uint8     `gorm:"column:sistem_sks;not null;default:0"`
	TotalSKS            uint16    `gorm:"column:total_sks;not null;default:0"`
	JenjangPendidikanID uint16    `gorm:"column:jenjang_pendidikan_id;not null"`
	JurusanID           *string   `gorm:"column:jurusan_id;size:25"`

	// CreateDate  time.Time  `gorm:"column:create_date;not null;default:'2019-09-10 14:29:56.948018'"`
	// LastUpdate  time.Time  `gorm:"column:last_update;not null;default:'2019-09-10 14:29:56.948018'"`
	// ExpiredDate *time.Time `gorm:"column:expired_date"`
	// LastSync    time.Time  `gorm:"column:last_sync;not null;default:'1901-01-01 00:00:00'"`

	// Relasi ke Jurusan
	Jurusan Jurusan `gorm:"foreignKey:JurusanID;references:JurusanID"`
}

func (Kurikulum) TableName() string {
	return "ref.kurikulum"
}

type MataPelajaran struct {
	MataPelajaranID     int32   `gorm:"column:mata_pelajaran_id;primaryKey"` // Primary key
	Nama                string  `gorm:"column:nama"`                         // Nama Mata Pelajaran
	PilihanSekolah      int32   `gorm:"column:pilihan_sekolah"`              // Kelompok
	PilihanBuku         int32   `gorm:"column:pilihan_buku"`                 // Semester
	PilihanKepengawasan int32   `gorm:"column:pilihan_kepengawasan"`         // Urutan di Rapor
	PilihanEvaluasi     int32   `gorm:"column:pilihan_evaluasi"`             // Nama Lokal
	JurusanID           *string `gorm:"column:jurusan_id"`                   // Foreign key ke tabel jurusan
	// Jurusan             Jurusan `gorm:"foreignKey:JurusanID;references:JurusanID"` // Relasi ke tabel jurusan
}

func (MataPelajaran) TableName() string {
	return "ref.mata_pelajaran"
}

type GelarAkademik struct {
	GelarAkademikID int32  `gorm:"column:gelar_akademik_id;primaryKey"` // Primary key
	Kode            string `gorm:"column:kode"`                         // Kelompok
	Nama            string `gorm:"column:nama"`                         // Nama Mata Pelajaran
	PosisiGelar     int32  `gorm:"column:posisi_gelar"`                 // Semester
}

func (GelarAkademik) TableName() string {
	return "ref.gelar_akademik"
}
