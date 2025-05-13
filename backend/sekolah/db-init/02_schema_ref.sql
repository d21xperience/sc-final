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
	last_sync TIMESTAMP NOT NULL,
	PRIMARY KEY (jurusan_id)
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
	PRIMARY KEY (kurikulum_id),
	CONSTRAINT FK_kurikulum_jurusan FOREIGN KEY (jurusan_id) REFERENCES ref.jurusan (jurusan_id) ON UPDATE NO ACTION ON DELETE NO ACTION
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
	PRIMARY KEY (semester_id),
	CONSTRAINT FK_semester_tahun_ajaran FOREIGN KEY (tahun_ajaran_id) REFERENCES ref.tahun_ajaran (tahun_ajaran_id) ON UPDATE CASCADE ON DELETE CASCADE

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

CREATE TABLE ref.bentuk_pendidikan (
	bentuk_pendidikan_id SMALLINT NOT NULL,
	nama VARCHAR(50) NOT NULL,
	jenjang_paud NUMERIC(1,0) NOT NULL,
	jenjang_tk NUMERIC(1,0) NOT NULL,
	jenjang_sd NUMERIC(1,0) NOT NULL,
	jenjang_smp NUMERIC(1,0) NOT NULL,
	jenjang_sma NUMERIC(1,0) NOT NULL,
	jenjang_tinggi NUMERIC(1,0) NOT NULL,
	direktorat_pembinaan VARCHAR(40) NULL DEFAULT NULL,
	aktif NUMERIC(1,0) NOT NULL,
	formalitas_pendidikan CHAR(1) NOT NULL,
	create_date TIMESTAMP NOT NULL DEFAULT '2020-04-16 09:40:03.422677',
	last_update TIMESTAMP NOT NULL DEFAULT '2020-04-16 09:40:03.422677',
	expired_date TIMESTAMP NULL DEFAULT NULL,
	last_sync TIMESTAMP NOT NULL DEFAULT '1901-01-01 00:00:00',
	PRIMARY KEY (bentuk_pendidikan_id)
);

CREATE TABLE ref.jenjang_pendidikan (
	jenjang_pendidikan_id NUMERIC(2,0) NOT NULL,
	nama VARCHAR(25) NOT NULL,
	jenjang_lembaga NUMERIC(1,0) NOT NULL,
	jenjang_orang NUMERIC(1,0) NOT NULL,
	create_date TIMESTAMP NOT NULL DEFAULT '2019-09-10 14:29:56.540627',
	last_update TIMESTAMP NOT NULL DEFAULT '2019-09-10 14:29:56.540627',
	expired_date TIMESTAMP NULL DEFAULT NULL,
	last_sync TIMESTAMP NOT NULL DEFAULT '1901-01-01 00:00:00',
	PRIMARY KEY (jenjang_pendidikan_id)
);

CREATE TABLE ref.status_kepemilikan (
	status_kepemilikan_id NUMERIC(1,0) NOT NULL,
	nama VARCHAR(20) NOT NULL,
	create_date TIMESTAMP NOT NULL DEFAULT '2019-09-10 14:29:59.426803',
	last_update TIMESTAMP NOT NULL DEFAULT '2019-09-10 14:29:59.426803',
	expired_date TIMESTAMP NULL DEFAULT NULL,
	last_sync TIMESTAMP NOT NULL DEFAULT '1901-01-01 00:00:00',
	PRIMARY KEY (status_kepemilikan_id)
);

CREATE TABLE ref.tingkat_pendidikan (
	tingkat_pendidikan_id NUMERIC(2,0) NOT NULL,
	kode VARCHAR(5) NOT NULL,
	nama VARCHAR(20) NOT NULL,
	jenjang_pendidikan_id NUMERIC(2,0) NOT NULL,
	create_date TIMESTAMP NOT NULL DEFAULT '2019-09-10 14:29:59.88044',
	last_update TIMESTAMP NOT NULL DEFAULT '2019-09-10 14:29:59.88044',
	expired_date TIMESTAMP NULL DEFAULT NULL,
	last_sync TIMESTAMP NOT NULL DEFAULT '1901-01-01 00:00:00',
	PRIMARY KEY (tingkat_pendidikan_id),
	CONSTRAINT fk_tingkat__tingkat_j_jenjang_ FOREIGN KEY (jenjang_pendidikan_id) REFERENCES ref.jenjang_pendidikan (jenjang_pendidikan_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE ref.gelar_akademik (
	gelar_akademik_id INTEGER NOT NULL,
	kode VARCHAR(10) NOT NULL,
	nama VARCHAR(40) NOT NULL,
	posisi_gelar NUMERIC(1,0) NOT NULL,
	create_date TIMESTAMP NOT NULL DEFAULT '2019-09-10 14:29:55.639014',
	last_update TIMESTAMP NOT NULL DEFAULT '2019-09-10 14:29:55.639014',
	expired_date TIMESTAMP NULL DEFAULT NULL,
	last_sync TIMESTAMP NOT NULL DEFAULT '1901-01-01 00:00:00',
	PRIMARY KEY (gelar_akademik_id)
);