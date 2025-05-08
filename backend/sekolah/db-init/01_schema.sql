CREATE SEQUENCE sekolah_tenant_id_seq;
CREATE TABLE sekolah_tenant (
	id BIGINT NOT NULL DEFAULT nextval('sekolah_tenant_id_seq'::regclass),
	nama_sekolah TEXT NOT NULL,
	sekolah_tenant_id BIGINT NOT NULL,
	schema_name TEXT NOT NULL,
	-- bentuk_pendidikan_id INTEGER NOT NULL DEFAULT '0',
	created_at TIMESTAMPTZ NULL DEFAULT NULL,
	updated_at TIMESTAMPTZ NULL DEFAULT NULL,
	deleted_at TIMESTAMPTZ NULL DEFAULT NULL,
	PRIMARY KEY (id)
);

CREATE UNIQUE INDEX uni_sekolah_tenant_nama_sekolah ON sekolah_tenant (nama_sekolah);
CREATE UNIQUE INDEX uni_sekolah_tenant_sekolah_id ON sekolah_tenant (sekolah_tenant_id);
CREATE UNIQUE INDEX uni_sekolah_tenant_schema_name ON sekolah_tenant (schema_name);
CREATE INDEX idx_sekolah_tenant_deleted_at ON sekolah_tenant (deleted_at);

CREATE SEQUENCE schema_logs_id_seq;
CREATE TABLE schema_logs (
	id BIGINT NOT NULL DEFAULT nextval('schema_logs_id_seq'::regclass),
	schema_name TEXT NULL DEFAULT NULL,
	created_at TIMESTAMPTZ NULL DEFAULT NULL,
	updated_at TIMESTAMPTZ NULL DEFAULT NULL,
	deleted_at TIMESTAMPTZ NULL DEFAULT NULL,
	PRIMARY KEY (id)
	-- CONSTRAINT "FK_schema_logs_sekolah_tenant" FOREIGN KEY (schema_name) REFERENCES sekolah_tenant (schema_name) ON UPDATE CASCADE ON DELETE CASCADE

);
CREATE INDEX idx_schema_logs_deleted_at ON schema_logs (deleted_at)