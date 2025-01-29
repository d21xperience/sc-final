CREATE TABLE sekolah_indonesia (
    sekolah_id_enkrip VARCHAR(50) NOT NULL PRIMARY KEY, -- Identifier unik
    nama_sekolah VARCHAR(255) NOT NULL,
    npsn VARCHAR(255) NOT NULL,
    alamat_jalan TEXT,
    kecamatan VARCHAR(100),
    kabupaten VARCHAR(100),
    propinsi VARCHAR(100),
    kode_kecamatan CHAR(6),
    kode_kab CHAR(6),
    kode_prop CHAR(6),
    status VARCHAR(50) CHECK (status IN ('Negeri', 'Swasta')) -- Validasi nilai
);

-- Tambahkan UNIQUE INDEX pada kolom npsn
CREATE UNIQUE INDEX uni_sekolah_indonesia_npsn ON sekolah_indonesia (npsn);

-- Tambahkan hash index untuk kolom sekolah_id_enkrip
CREATE INDEX idx_sekolah_id_enkrip_hash ON sekolah_indonesia (MD5(sekolah_id_enkrip));

-- Tambahkan hash index untuk kolom npsn
CREATE INDEX idx_npsn_hash ON sekolah_indonesia (MD5(npsn));

