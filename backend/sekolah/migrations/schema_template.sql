CREATE EXTENSION IF NOT EXISTS uuid-ossp;
CREATE SCHEMA IF NOT EXISTS {{schema_name}};

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_sekolah (
	sekolah_id UUID NOT NULL,
	nama VARCHAR(100) NOT NULL,
	npsn VARCHAR(8) NULL DEFAULT 'NULL::character varying',
	nss VARCHAR(12) NULL DEFAULT 'NULL::character varying',
	alamat TEXT NULL DEFAULT NULL,
	kd_pos VARCHAR(6) NULL DEFAULT 'NULL::character varying',
	telepon VARCHAR(20) NULL DEFAULT 'NULL::character varying',
	fax VARCHAR(20) NULL DEFAULT 'NULL::character varying',
	kelurahan VARCHAR(60) NULL DEFAULT 'NULL::character varying',
	kecamatan VARCHAR(60) NULL DEFAULT 'NULL::character varying',
	kab_kota VARCHAR(60) NULL DEFAULT 'NULL::character varying',
	propinsi VARCHAR(60) NULL DEFAULT 'NULL::character varying',
	website VARCHAR(100) NULL DEFAULT 'NULL::character varying',
	email VARCHAR(50) NULL DEFAULT 'NULL::character varying',
	nm_kepsek VARCHAR(100) NULL DEFAULT 'NULL::character varying',
	nip_kepsek VARCHAR(25) NULL DEFAULT 'NULL::character varying',
	niy_kepsek VARCHAR(30) NULL DEFAULT 'NULL::character varying',
	status_kepemilikan_id NUMERIC(1,0) NULL DEFAULT NULL,
	kode_aktivasi VARCHAR(30) NULL DEFAULT 'NULL::character varying',
	bentuk_pendidikan_id SMALLINT NULL DEFAULT NULL,
	jenjang_pendidikan_id NUMERIC(2,0) NULL DEFAULT NULL,
	PRIMARY KEY (sekolah_id),
	CONSTRAINT FK_tabel_sekolah_ref.bentuk_pendidikan FOREIGN KEY (bentuk_pendidikan_id) REFERENCES ref.bentuk_pendidikan (bentuk_pendidikan_id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT FK_tabel_sekolah_ref.jenjang_pendidikan FOREIGN KEY (jenjang_pendidikan_id) REFERENCES ref.jenjang_pendidikan (jenjang_pendidikan_id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT FK_tabel_sekolah_ref.status_kepemilikan FOREIGN KEY (status_kepemilikan_id) REFERENCES ref.status_kepemilikan (status_kepemilikan_id) ON UPDATE CASCADE ON DELETE CASCADE
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
	pekerjaan_wali VARCHAR(30) NULL DEFAULT NULL,
	nik VARCHAR(30) NULL DEFAULT NULL,
	PRIMARY KEY (peserta_didik_id)
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_siswa_pelengkap (
	pelengkap_siswa_id UUID NOT NULL DEFAULT public.uuid_generate_v4(),
	peserta_didik_id UUID NULL DEFAULT NULL,
	status_dalam_kel VARCHAR(30) NULL DEFAULT 'NULL::character varying',
	anak_ke NUMERIC(3,0) NULL DEFAULT NULL,
	sekolah_asal VARCHAR(100) NULL DEFAULT 'NULL::character varying',
	diterima_kelas VARCHAR(20) NULL DEFAULT 'NULL::character varying',
	alamat_ortu TEXT NULL DEFAULT NULL,
	telepon_ortu VARCHAR(20) NULL DEFAULT 'NULL::character varying',
	alamat_wali TEXT NULL DEFAULT NULL,
	telepon_wali VARCHAR(20) NULL DEFAULT 'NULL::character varying',
	foto_siswa VARCHAR(100) NULL DEFAULT 'NULL::character varying',
	PRIMARY KEY (pelengkap_siswa_id),
	CONSTRAINT FK_tabel_siswa_pelengkap_tabel_siswa FOREIGN KEY (peserta_didik_id) REFERENCES {{schema_name}}.tabel_siswa (peserta_didik_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_ptk (
	ptk_id UUID NOT NULL,
	jenis_ptk_id NUMERIC(2,0) NOT NULL,
	nama VARCHAR(100) NOT NULL,
	nip VARCHAR(18) NULL DEFAULT NULL,
	nik CHAR(16) NOT NULL,
	no_kk CHAR(16) NULL DEFAULT NULL,
	agama VARCHAR(25) NULL DEFAULT NULL,
	jenis_kelamin VARCHAR(1) NOT NULL,
	tempat_lahir VARCHAR(32) NOT NULL,
	tanggal_lahir DATE NOT NULL,
	nuptk VARCHAR(16) NULL DEFAULT NULL,
	alamat_jalan VARCHAR(200) NOT NULL,
	rt NUMERIC(2,0) NULL DEFAULT NULL,
	rw NUMERIC(2,0) NULL DEFAULT NULL,
	desa_kelurahan VARCHAR(60) NOT NULL,
	kab_kota VARCHAR(60) NULL DEFAULT 'NULL::character varying',
	propinsi VARCHAR(60) NULL DEFAULT 'NULL::character varying',
	kode_pos CHAR(5) NULL DEFAULT NULL,
	no_telepon_rumah VARCHAR(20) NULL DEFAULT NULL,
	no_hp VARCHAR(20) NULL DEFAULT NULL,
	email VARCHAR(60) NULL DEFAULT NULL,
	status_keaktifan_id NUMERIC(2,0) NOT NULL DEFAULT 1,  -- Contoh nilai default
    soft_delete NUMERIC(1,0) NOT NULL DEFAULT 0,  -- 0 = aktif, 1 = terhapus
	PRIMARY KEY (ptk_id)
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_ptk_terdaftar (
	ptk_terdaftar_id UUID NOT NULL DEFAULT public.uuid_generate_v4(),
	ptk_id UUID NOT NULL,
	tahun_ajaran_id VARCHAR(4) NULL DEFAULT NULL,
	jenis_keluar_id CHAR(1) NULL DEFAULT NULL,
	soft_delete NUMERIC(1,0) NOT NULL DEFAULT 0,  -- 0 = aktif, 1 = terhapus
	PRIMARY KEY (ptk_terdaftar_id),
	CONSTRAINT FK_tabel_ptk_terdaftar_tabel_ptk FOREIGN KEY (ptk_id) REFERENCES {{schema_name}}.tabel_ptk (ptk_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_kelas (
	rombongan_belajar_id UUID NOT NULL DEFAULT public.uuid_generate_v4(),
	sekolah_id UUID NOT NULL,
	semester_id CHAR(5) NOT NULL,
	jurusan_id VARCHAR(25) NULL DEFAULT NULL,
	ptk_id UUID NULL DEFAULT NULL,
	nm_kelas VARCHAR(30) NULL DEFAULT NULL,
	tingkat_pendidikan_id NUMERIC(2,0) NULL DEFAULT NULL,
	jenis_rombel NUMERIC(2,0) NULL DEFAULT NULL,
	nama_jurusan_sp VARCHAR(100) NULL DEFAULT NULL,
	jurusan_sp_id UUID NULL DEFAULT NULL,
	kurikulum_id SMALLINT NOT NULL,
	PRIMARY KEY (rombongan_belajar_id),
	CONSTRAINT FK_tabel_kelas_ref.jurusan FOREIGN KEY (jurusan_id) REFERENCES ref.jurusan (jurusan_id) ON UPDATE CASCADE ON DELETE SET NULL,
	CONSTRAINT FK_tabel_kelas_ref.kurikulum FOREIGN KEY (kurikulum_id) REFERENCES ref.kurikulum (kurikulum_id) ON UPDATE CASCADE ON DELETE SET NULL,
	CONSTRAINT FK_tabel_kelas_ref.semester FOREIGN KEY (semester_id) REFERENCES ref.semester (semester_id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT FK_tabel_kelas_ref.tingkat_pendidikan FOREIGN KEY (tingkat_pendidikan_id) REFERENCES ref.tingkat_pendidikan (tingkat_pendidikan_id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT FK_tabel_kelas_tabel_ptk FOREIGN KEY (ptk_id) REFERENCES {{schema_name}}.tabel_ptk (ptk_id) ON UPDATE CASCADE ON DELETE SET NULL

);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_anggotakelas (
	anggota_rombel_id UUID NOT NULL DEFAULT public.uuid_generate_v4(),
	peserta_didik_id UUID NOT NULL,
	rombongan_belajar_id UUID NOT NULL,
	semester_id CHAR(5) NOT NULL,
	status_keaktifan INTEGER NOT NULL DEFAULT 0, -- 0 = aktif(sedang studi); 1= lulus; 2=keluar 
	PRIMARY KEY (anggota_rombel_id),
	CONSTRAINT FK_tabel_anggotakelas_tabel_anggotakelas FOREIGN KEY (rombongan_belajar_id) REFERENCES {{schema_name}}.tabel_kelas (rombongan_belajar_id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT FK_tabel_anggotakelas_tabel_siswa FOREIGN KEY (peserta_didik_id) REFERENCES {{schema_name}}.tabel_siswa (peserta_didik_id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT FK_tabel_anggotakelas_ref.semester FOREIGN KEY (semester_id) REFERENCES ref.semester (semester_id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_pembelajaran (
	pembelajaran_id UUID NOT NULL,
	rombongan_belajar_id UUID NOT NULL,
	mata_pelajaran_id INTEGER NOT NULL,
	semester_id CHAR(5) NOT NULL,
	ptk_terdaftar_id UUID NULL DEFAULT NULL,
	status_di_kurikulum NUMERIC(2,0) NULL DEFAULT NULL,
	nama_mata_pelajaran VARCHAR(50) NULL DEFAULT NULL,
	induk_pembelajaran UUID NULL DEFAULT NULL,
	is_dapo NUMERIC(1,0) NULL DEFAULT '1',
	PRIMARY KEY (pembelajaran_id),
	CONSTRAINT FK_tabel_pembelajaran_ref.mata_pelajaran FOREIGN KEY (mata_pelajaran_id) REFERENCES ref.mata_pelajaran (mata_pelajaran_id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT FK_tabel_pembelajaran_ref.semester FOREIGN KEY (semester_id) REFERENCES ref.semester (semester_id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT FK_tabel_pembelajaran_tabel_kelas FOREIGN KEY (rombongan_belajar_id) REFERENCES {{schema_name}}.tabel_kelas (rombongan_belajar_id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT FK_tabel_pembelajaran_tabel_ptk_terdaftar FOREIGN KEY (ptk_terdaftar_id) REFERENCES {{schema_name}}.tabel_ptk_terdaftar (ptk_terdaftar_id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_nilaiakhir (
	id_nilai_akhir UUID NOT NULL DEFAULT public.uuid_generate_v4(),
	anggota_rombel_id UUID NULL DEFAULT NULL,
	mata_pelajaran_id INTEGER NULL DEFAULT NULL,
	semester_id CHAR(5) NULL DEFAULT NULL,
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
	semester NUMERIC(1,0) NULL DEFAULT NULL,
	PRIMARY KEY (id_nilai_akhir),
	CONSTRAINT FK_tabel_nilaiakhir_tabel_anggotakelas FOREIGN KEY (anggota_rombel_id) REFERENCES {{schema_name}}.tabel_anggotakelas (anggota_rombel_id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT FK_tabel_nilaiakhir_tabel_siswa FOREIGN KEY (peserta_didik_id) REFERENCES {{schema_name}}.tabel_siswa (peserta_didik_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_kenaikan (
	kd_kenaikan UUID NOT NULL DEFAULT public.uuid_generate_v4(),
	semester_id CHAR(5) NOT NULL,
	anggota_rombel_id UUID NOT NULL,
	peserta_didik_id UUID NULL DEFAULT NULL,
	kenaikan NUMERIC(3,0) NULL DEFAULT NULL,
	tingkat NUMERIC(3,0) NULL DEFAULT NULL,
	PRIMARY KEY (kd_kenaikan),
	CONSTRAINT FK_tabel_kenaikan_ref.semester FOREIGN KEY (semester_id) REFERENCES ref.semester (semester_id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT FK_tabel_kenaikan_tabel_anggotakelas FOREIGN KEY (anggota_rombel_id) REFERENCES {{schema_name}}.tabel_anggotakelas (anggota_rombel_id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT FK_tabel_kenaikan_tabel_siswa FOREIGN KEY (peserta_didik_id) REFERENCES {{schema_name}}.tabel_siswa (peserta_didik_id) ON UPDATE CASCADE ON DELETE CASCADE

);

CREATE TABLE IF NOT EXISTS {{schema_name}}.ijazah (
    id UUID NOT NULL DEFAULT public.uuid_generate_v4(),
	PRIMARY KEY (id),
	peserta_didik_id UUID NULL DEFAULT NULL,
    program_keahlian VARCHAR(100) NOT NULL,
    paket_keahlian VARCHAR(100) NOT NULL,
    sekolah_id UUID NOT NULL,
    npsn VARCHAR(15) NOT NULL,
    kabupaten_kota VARCHAR(100) NOT NULL,
    provinsi VARCHAR(100) NOT NULL,
    nama VARCHAR(200) NOT NULL,
    tempat_lahir VARCHAR(100) NOT NULL,
    tanggal_lahir DATE NOT NULL,
    nis VARCHAR(20) UNIQUE NOT NULL,
    nisn VARCHAR(20) UNIQUE NOT NULL,
    nama_ortu_wali VARCHAR(200) NOT NULL,
    sekolah_penyelenggara_ujian_us VARCHAR(200) NOT NULL,
    sekolah_penyelenggara_ujian_un VARCHAR(200) NOT NULL,
    asal_sekolah VARCHAR(200) NOT NULL,
    nomor_ijazah VARCHAR(50) UNIQUE NOT NULL,
    tempat_ijazah VARCHAR(100) NOT NULL,
    tanggal_ijazah DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
   	CONSTRAINT FK_tabel_ijazah FOREIGN KEY (peserta_didik_id) REFERENCES {{schema_name}}.tabel_siswa (peserta_didik_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.tabel_ptk_pelengkap (
	ptk_pelengkap_id UUID NOT NULL DEFAULT public.uuid_generate_v4(),
	ptk_id UUID NOT NULL,
	gelar_depan VARCHAR(20) NULL DEFAULT 'NULL::character varying',
	gelar_belakang VARCHAR(20) NULL DEFAULT 'NULL::character varying',
	nip_niy VARCHAR(18) NULL DEFAULT 'NULL::character varying',
	PRIMARY KEY (ptk_pelengkap_id),
	CONSTRAINT FK_tabel_ptk_pelengkap_tabel_ptk FOREIGN KEY (ptk_id) REFERENCES {{schema_name}}.tabel_ptk (ptk_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS {{schema_name}}.data_nominasi_sementara (
    id UUID NOT NULL DEFAULT public.uuid_generate_v4(),
	PRIMARY KEY (id),
	peserta_didik_id UUID NULL DEFAULT NULL,
	rombongan_belajar_id NULL DEFAULT NULL,
	tahun_ajaran_id VARCHAR(4) NULL DEFAULT NULL,
    program_keahlian VARCHAR(100) NOT NULL,
    paket_keahlian VARCHAR(100) NOT NULL,
    sekolah_id UUID NOT NULL,
    npsn VARCHAR(15) NOT NULL,
    kabupaten_kota VARCHAR(100) NOT NULL,
    provinsi VARCHAR(100) NOT NULL,
    nama VARCHAR(200) NOT NULL,
    tempat_lahir VARCHAR(100) NOT NULL,
    tanggal_lahir DATE NOT NULL,
    nis VARCHAR(20) UNIQUE NOT NULL,
    nisn VARCHAR(20) UNIQUE NOT NULL,
    nama_ortu_wali VARCHAR(200) NOT NULL,
    sekolah_penyelenggara_ujian_us VARCHAR(200) NOT NULL,
    sekolah_penyelenggara_ujian_un VARCHAR(200) NOT NULL,
    asal_sekolah VARCHAR(200) NOT NULL,
    nomor_ijazah VARCHAR(50) UNIQUE NOT NULL,
    tempat_ijazah VARCHAR(100) NOT NULL,
    tanggal_ijazah DATE NOT NULL,
	is_complete BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
   	CONSTRAINT FK_tabel_ijazah FOREIGN KEY (peserta_didik_id) REFERENCES {{schema_name}}.tabel_siswa (peserta_didik_id) ON UPDATE CASCADE ON DELETE CASCADE,
   	CONSTRAINT FK_data_nominasi_sementara_tabel_kelas FOREIGN KEY (rombongan_belajar_id) REFERENCES {{schema_name}}.tabel_kelas (rombongan_belajar_id) ON UPDATE CASCADE ON DELETE CASCADE
);

-- CREATE OR REPLACE FUNCTION {{schema_name}}.check_is_complete()
-- RETURNS TRIGGER AS $$
-- BEGIN
--     -- Ganti pengecekan NULL sesuai kebutuhan kolom
--     IF NEW.peserta_didik_id IS NOT NULL AND
--        NEW.rombongan_belajar_id IS NOT NULL AND
--        NEW.program_keahlian IS NOT NULL AND
--        NEW.paket_keahlian IS NOT NULL AND
--        NEW.sekolah_id IS NOT NULL AND
--        NEW.npsn IS NOT NULL AND
--        NEW.kabupaten_kota IS NOT NULL AND
--        NEW.provinsi IS NOT NULL AND
--        NEW.nama IS NOT NULL AND
--        NEW.tempat_lahir IS NOT NULL AND
--        NEW.tanggal_lahir IS NOT NULL AND
--        NEW.nis IS NOT NULL AND
--        NEW.nisn IS NOT NULL AND
--        NEW.nama_ortu_wali IS NOT NULL AND
--        NEW.sekolah_penyelenggara_ujian_us IS NOT NULL AND
--        NEW.sekolah_penyelenggara_ujian_un IS NOT NULL AND
--        NEW.asal_sekolah IS NOT NULL AND
--        NEW.nomor_ijazah IS NOT NULL AND
--        NEW.tempat_ijazah IS NOT NULL AND
--        NEW.tanggal_ijazah IS NOT NULL THEN
       
--        NEW.is_complete := TRUE;
--     ELSE
--        NEW.is_complete := FALSE;
--     END IF;
--     RETURN NEW;
-- END;
-- $$ LANGUAGE plpgsql;

-- CREATE TRIGGER trigger_check_is_complete
-- BEFORE INSERT OR UPDATE ON {{schema_name}}.data_nominasi_sementara
-- FOR EACH ROW
-- EXECUTE FUNCTION {{schema_name}}.check_is_complete();
