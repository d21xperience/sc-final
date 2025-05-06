CREATE SEQUENCE sekolah_id_seq;
CREATE TABLE sekolah_tenant (
	id BIGINT NOT NULL DEFAULT nextval('sekolah_id_seq'::regclass),
	sekolah_id VARCHAR NULL,
	nama_sekolah VARCHAR NOT NULL,
	npsn VARCHAR(8) NULL,
	enkrip_id VARCHAR NULL,
	kecamatan VARCHAR(50) NULL,
	kabupaten VARCHAR(50) NULL,
	propinsi VARCHAR(50) NULL,
	kode_kecamatan VARCHAR(50) NULL,
	kode_kab VARCHAR(50) NULL,
	kode_prop VARCHAR(50) NULL,
	alamat_jalan VARCHAR(50) NULL,
	created_at TIMESTAMPTZ NULL,
	updated_at TIMESTAMPTZ NULL,
	PRIMARY KEY (id)
);
CREATE UNIQUE INDEX uni_sekolah_tenant_sekolah_id_enkrip ON sekolah_tenant (enkrip_id);
CREATE UNIQUE INDEX uni_sekolah_tenant_npsn ON sekolah_tenant (npsn); 

CREATE SEQUENCE users_id_seq;
CREATE TABLE users (
	id BIGINT NOT NULL DEFAULT nextval('users_id_seq'::regclass),
	sekolah_tenant_id BIGINT NULL DEFAULT NULL,
	username VARCHAR NOT NULL,
	email VARCHAR NOT NULL,
	password VARCHAR NOT NULL,
	role VARCHAR NULL DEFAULT NULL,
	is_initial_password BOOLEAN NOT NULL DEFAULT 'true',
	initial_password VARCHAR NULL DEFAULT NULL,
	last_login TIMESTAMPTZ NULL DEFAULT NULL,
	created_at TIMESTAMPTZ NULL DEFAULT NULL,
	PRIMARY KEY (id),
	CONSTRAINT fk_users_sekolah FOREIGN KEY (sekolah_tenant_id) REFERENCES sekolah_tenant (id) ON UPDATE CASCADE ON DELETE CASCADE
);
CREATE UNIQUE INDEX uni_users_username ON users (username);
CREATE UNIQUE INDEX uni_users_email ON users (email);

CREATE SEQUENCE user_profiles_id_seq;
CREATE TABLE user_profiles (
	id BIGINT NOT NULL DEFAULT nextval('user_profiles_id_seq'::regclass),
	user_id BIGINT NULL DEFAULT NULL,
	nama VARCHAR(100) NULL DEFAULT NULL,
	jk VARCHAR(100) NULL DEFAULT NULL,
	phone VARCHAR(100) NULL DEFAULT NULL,
	tpt_lahir VARCHAR(100) NULL DEFAULT NULL,
	tgl_lahir DATE NULL,
	alamat_jalan VARCHAR(100) NULL DEFAULT NULL,
	kota_kab VARCHAR(100) NULL DEFAULT NULL,
	prov VARCHAR(100) NULL DEFAULT NULL,
	kode_pos VARCHAR(100) NULL DEFAULT NULL,
	nama_ayah VARCHAR(100) NULL DEFAULT NULL,
	nama_ibu VARCHAR(100) NULL DEFAULT NULL,
	photo_url VARCHAR(255) NULL DEFAULT NULL,
	created_at TIMESTAMPTZ NULL DEFAULT NULL,
	updated_at TIMESTAMPTZ NULL DEFAULT NULL,
	deleted_at TIMESTAMPTZ NULL DEFAULT NULL,
	PRIMARY KEY (id),
	CONSTRAINT fk_user_profiles_user FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE INDEX idx_user_profiles_deleted_at ON user_profiles (deleted_at);