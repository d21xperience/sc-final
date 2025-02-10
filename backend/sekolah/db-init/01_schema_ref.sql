CREATE SCHEMA IF NOT EXISTS ref;
CREATE TABLE IF NOT EXISTS ref.jurusan (
	jurusan_id VARCHAR(25) NOT NULL,
	nama_jurusan VARCHAR(100) NOT NULL,
	untuk_sma NUMERIC(1,0) NOT NULL,
	untuk_smk NUMERIC(1,0) NOT NULL,
	untuk_pt NUMERIC(1,0) NOT NULL,
	untuk_slb NUMERIC(1,0) NOT NULL,
	untuk_smklb NUMERIC(1,0) NOT NULL,
	jenjang_pendidikan_id NUMERIC(2,0) NULL DEFAULT NULL,
	jurusan_induk VARCHAR(25) NULL DEFAULT NULL,
	level_bidang_id VARCHAR(5) NOT NULL,
	create_date TIMESTAMP NOT NULL,
	last_update TIMESTAMP NOT NULL,
	expired_date TIMESTAMP NULL DEFAULT NULL,
	last_sync TIMESTAMP NOT NULL
);


CREATE TABLE IF NOT EXISTS ref.kurikulum (
	kurikulum_id SMALLINT NOT NULL,
	nama_kurikulum VARCHAR(120) NOT NULL,
	mulai_berlaku DATE NOT NULL,
	sistem_sks NUMERIC(1,0) NOT NULL DEFAULT '0',
	total_sks NUMERIC(3,0) NOT NULL DEFAULT '0',
	jenjang_pendidikan_id NUMERIC(2,0) NOT NULL,
	jurusan_id VARCHAR(25) NULL DEFAULT NULL,
	create_date TIMESTAMP NOT NULL DEFAULT '2019-09-10 14:29:56.948018',
	last_update TIMESTAMP NOT NULL DEFAULT '2019-09-10 14:29:56.948018',
	expired_date TIMESTAMP NULL DEFAULT NULL,
	last_sync TIMESTAMP NOT NULL DEFAULT '1901-01-01 00:00:00',
	PRIMARY KEY (kurikulum_id)
);
CREATE TABLE IF NOT EXISTS ref.tahun_ajaran (
	tahun_ajaran_id NUMERIC(4,0) NOT NULL,
	nama VARCHAR(10) NOT NULL,
	periode_aktif NUMERIC(1,0) NOT NULL,
	tanggal_mulai DATE NOT NULL,
	tanggal_selesai DATE NOT NULL,
	create_date TIMESTAMP NOT NULL DEFAULT '2025-01-19 14:29:59.628052',
	last_update TIMESTAMP NOT NULL DEFAULT '2025-01-19 14:29:59.628052',
	expired_date TIMESTAMP NULL DEFAULT NULL,
	last_sync TIMESTAMP NOT NULL DEFAULT '1901-01-01 00:00:00',
	PRIMARY KEY (tahun_ajaran_id)
);

CREATE TABLE IF NOT EXISTS ref.semester (
	semester_id CHAR(5) NOT NULL,
	tahun_ajaran_id NUMERIC(4,0) NOT NULL,
	nama VARCHAR(20) NOT NULL,
	semester NUMERIC(1,0) NOT NULL,
	periode_aktif NUMERIC(1,0) NOT NULL,
	tanggal_mulai DATE NOT NULL,
	tanggal_selesai DATE NOT NULL,
	create_date TIMESTAMP NOT NULL DEFAULT '2019-09-10 14:29:59.238151',
	last_update TIMESTAMP NOT NULL DEFAULT '2019-09-10 14:29:59.238151',
	expired_date TIMESTAMP NULL DEFAULT NULL,
	last_sync TIMESTAMP NOT NULL DEFAULT '1901-01-01 00:00:00',
	PRIMARY KEY (semester_id)
);

CREATE TABLE IF NOT EXISTS ref.mata_pelajaran (
	mata_pelajaran_id INTEGER NOT NULL,
	nama VARCHAR(80) NOT NULL,
	pilihan_sekolah NUMERIC(1,0) NOT NULL,
	pilihan_buku NUMERIC(1,0) NOT NULL,
	pilihan_kepengawasan NUMERIC(1,0) NOT NULL,
	pilihan_evaluasi NUMERIC(1,0) NOT NULL DEFAULT '0',
	jurusan_id VARCHAR(25) NULL DEFAULT NULL,
	create_date TIMESTAMP NOT NULL DEFAULT '2019-09-10 14:29:57.296154',
	last_update TIMESTAMP NOT NULL DEFAULT '2019-09-10 14:29:57.296154',
	expired_date TIMESTAMP NULL DEFAULT NULL,
	last_sync TIMESTAMP NOT NULL DEFAULT '1901-01-01 00:00:00',
	PRIMARY KEY (mata_pelajaran_id)
);


