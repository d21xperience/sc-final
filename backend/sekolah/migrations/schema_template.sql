CREATE SCHEMA IF NOT EXISTS {{schema_name}};

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_sekolah (
	sekolah_id UUID NOT NULL,
	nama VARCHAR(100) NOT NULL,
	npsn VARCHAR(8) NULL DEFAULT NULL,
	nss VARCHAR(12) NULL DEFAULT NULL,
	alamat TEXT NULL DEFAULT NULL,
	kd_pos VARCHAR(6) NULL DEFAULT NULL,
	telepon VARCHAR(20) NULL DEFAULT NULL,
	fax VARCHAR(20) NULL DEFAULT NULL,
	kelurahan VARCHAR(60) NULL DEFAULT NULL,
	kecamatan VARCHAR(60) NULL DEFAULT NULL,
	kab_kota VARCHAR(60) NULL DEFAULT NULL,
	propinsi VARCHAR(60) NULL DEFAULT NULL,
	website VARCHAR(100) NULL DEFAULT NULL,
	email VARCHAR(50) NULL DEFAULT NULL,
	nm_kepsek VARCHAR(100) NULL DEFAULT NULL,
	nip_kepsek VARCHAR(25) NULL DEFAULT NULL,
	niy_kepsek VARCHAR(30) NULL DEFAULT NULL,
	status_kepemilikan_id NUMERIC(1,0) NOT NULL,
	kode_aktivasi VARCHAR(30) NULL DEFAULT NULL,
	jenjang VARCHAR(20) NULL DEFAULT NULL,
	bentuk_pendidikan_id SMALLINT NULL DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.semester (
	semester_id VARCHAR(5) NOT NULL,
	tahun_ajaran_id VARCHAR(4) NOT NULL,
	nama_semester VARCHAR(20) NOT NULL,
	semester NUMERIC(1,0) NOT NULL,
	periode_aktif VARCHAR(1) NOT NULL,
	tanggal_mulai DATE NOT NULL,
	tanggal_selesai DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_siswa (
	peserta_didik_id UUID NOT NULL,
	nis VARCHAR(20) NOT NULL,
	nisn VARCHAR(13) NULL DEFAULT NULL,
	nm_siswa VARCHAR(100) NOT NULL,
	tempat_lahir VARCHAR(50) NULL DEFAULT NULL,
	tanggal_lahir DATE NULL DEFAULT NULL,
	jenis_kelamin VARCHAR(1) NULL DEFAULT NULL,
	agama VARCHAR(25) NULL DEFAULT NULL,
	alamat_siswa TEXT NULL DEFAULT NULL,
	telepon_siswa VARCHAR(20) NULL DEFAULT NULL,
	diterima_tanggal DATE NULL DEFAULT NULL,
	nm_ayah VARCHAR(100) NULL DEFAULT NULL,
	nm_ibu VARCHAR(100) NULL DEFAULT NULL,
	pekerjaan_ayah VARCHAR(30) NULL DEFAULT NULL,
	pekerjaan_ibu VARCHAR(30) NULL DEFAULT NULL,
	nm_wali VARCHAR(100) NULL DEFAULT NULL,
	pekerjaan_wali VARCHAR(30) NULL DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_siswa_pelengkap (
	pelengkap_siswa_id UUID NOT NULL,
	peserta_didik_id UUID NOT NULL,
	status_dalam_kel VARCHAR(30) NULL DEFAULT NULL,
	anak_ke NUMERIC(3,0) NULL DEFAULT NULL,
	sekolah_asal VARCHAR(100) NULL DEFAULT NULL,
	diterima_kelas VARCHAR(20) NULL DEFAULT NULL,
	alamat_ortu TEXT NULL DEFAULT NULL,
	telepon_ortu VARCHAR(20) NULL DEFAULT NULL,
	alamat_wali TEXT NULL DEFAULT NULL,
	telepon_wali VARCHAR(20) NULL DEFAULT NULL,
	foto_siswa VARCHAR(100) NULL DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_ptk (
	ptk_id UUID NOT NULL,
	nama VARCHAR(100) NOT NULL,
	nip VARCHAR(18) NULL DEFAULT NULL,
	jenis_ptk_id NUMERIC(2,0) NOT NULL,
	jenis_kelamin VARCHAR(1) NOT NULL,
	tempat_lahir VARCHAR(32) NOT NULL,
	tanggal_lahir DATE NOT NULL,
	nuptk VARCHAR(16) NULL DEFAULT NULL,
	alamat_jalan VARCHAR(80) NOT NULL,
	status_keaktifan_id NUMERIC(2,0) NOT NULL,
	soft_delete NUMERIC(1,0) NOT NULL
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_ptk_terdaftar (
	ptk_terdaftar_id UUID NOT NULL,
	ptk_id UUID NOT NULL,
	tahun_ajaran_id VARCHAR(4) NULL DEFAULT NULL,
	jenis_keluar_id CHAR(1) NULL DEFAULT NULL,
	soft_delete NUMERIC(1,0) NOT NULL
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_kelas (
	rombongan_belajar_id UUID NOT NULL,
	sekolah_id UUID NOT NULL,
	semester_id VARCHAR(5) NOT NULL,
	jurusan_id VARCHAR(25) NULL DEFAULT NULL,
	ptk_id UUID NULL DEFAULT NULL,
	nm_kelas VARCHAR(30) NULL DEFAULT NULL,
	tingkat_pendidikan_id NUMERIC(2,0) NULL DEFAULT NULL,
	jenis_rombel NUMERIC(2,0) NULL DEFAULT NULL,
	nama_jurusan_sp VARCHAR(100) NULL DEFAULT NULL,
	jurusan_sp_id UUID NULL DEFAULT NULL,
	kurikulum_id SMALLINT NOT NULL
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_anggotakelas (
	anggota_rombel_id UUID NOT NULL,
	peserta_didik_id UUID NOT NULL,
	rombongan_belajar_id UUID NOT NULL,
	semester_id VARCHAR(5) NOT NULL
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_mapel (
	mata_pelajaran_id INTEGER NOT NULL,
	nm_mapel VARCHAR(100) NULL DEFAULT NULL,
	kelompok VARCHAR(2) NULL DEFAULT NULL,
	semester NUMERIC(1,0) NULL DEFAULT NULL,
	jurusan_id VARCHAR(25) NULL DEFAULT NULL,
	urut_rapor INTEGER NULL DEFAULT NULL,
	nm_lokal VARCHAR(60) NULL DEFAULT NULL,
	nm_ringkas VARCHAR(10) NULL DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_pembelajaran (
	pembelajaran_id UUID NOT NULL,
	rombongan_belajar_id UUID NOT NULL,
	mata_pelajaran_id INTEGER NOT NULL,
	semester_id VARCHAR(5) NOT NULL,
	ptk_terdaftar_id UUID NULL DEFAULT NULL,
	status_di_kurikulum NUMERIC(2,0) NULL DEFAULT NULL,
	nama_mata_pelajaran VARCHAR(50) NULL DEFAULT NULL,
	induk_pembelajaran UUID NULL DEFAULT NULL,
	is_dapo NUMERIC(1,0) NULL DEFAULT '1'
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_nilaiakhir (
	id_nilai_akhir UUID NOT NULL,
	anggota_rombel_id UUID NULL DEFAULT NULL,
	mata_pelajaran_id INTEGER NULL DEFAULT NULL,
	semester_id VARCHAR(5) NULL DEFAULT NULL,
	nilai_peng NUMERIC(5,0) NULL DEFAULT NULL,
	predikat_peng VARCHAR(1) NULL DEFAULT NULL,
	nilai_ket NUMERIC(5,0) NULL DEFAULT NULL,
	predikat_ket VARCHAR(1) NULL DEFAULT NULL,
	nilai_sik NUMERIC(2,0) NULL DEFAULT NULL,
	predikat_sik VARCHAR(15) NULL DEFAULT NULL,
	nilai_siksos NUMERIC(2,0) NULL DEFAULT NULL,
	predikat_siksos VARCHAR(15) NULL DEFAULT NULL,
	peserta_didik_id UUID NULL DEFAULT NULL,
	id_minat VARCHAR(2) NULL DEFAULT NULL,
	semester NUMERIC(1,0) NULL DEFAULT NULL
);

