CREATE TABLE "sekolah" (
	"sekolah_id" UUID NOT NULL,
	"nama" VARCHAR(100) NOT NULL,
	"nama_nomenklatur" VARCHAR(100) NULL DEFAULT NULL,
	"nss" CHAR(12) NULL DEFAULT NULL,
	"npsn" CHAR(8) NULL DEFAULT NULL,
	"bentuk_pendidikan_id" SMALLINT NOT NULL,
	"alamat_jalan" VARCHAR(80) NOT NULL,
	"rt" NUMERIC(2,0) NULL DEFAULT NULL,
	"rw" NUMERIC(2,0) NULL DEFAULT NULL,
	"nama_dusun" VARCHAR(60) NULL DEFAULT NULL,
	"desa_kelurahan" VARCHAR(60) NOT NULL,
	"kode_wilayah" CHAR(8) NOT NULL,
	"kode_pos" CHAR(5) NULL DEFAULT NULL,
	"lintang" NUMERIC(18,12) NULL DEFAULT NULL,
	"bujur" NUMERIC(18,12) NULL DEFAULT NULL,
	"nomor_telepon" VARCHAR(20) NULL DEFAULT NULL,
	"nomor_fax" VARCHAR(20) NULL DEFAULT NULL,
	"email" VARCHAR(60) NULL DEFAULT NULL,
	"website" VARCHAR(100) NULL DEFAULT NULL,
	"kebutuhan_khusus_id" INTEGER NOT NULL,
	"status_sekolah" NUMERIC(1,0) NOT NULL,
	"sk_pendirian_sekolah" VARCHAR(80) NULL DEFAULT NULL,
	"tanggal_sk_pendirian" DATE NULL DEFAULT NULL,
	"status_kepemilikan_id" NUMERIC(1,0) NOT NULL,
	"yayasan_id" UUID NULL DEFAULT NULL,
	"sk_izin_operasional" VARCHAR(80) NULL DEFAULT NULL,
	"tanggal_sk_izin_operasional" DATE NULL DEFAULT NULL,
	"no_rekening" VARCHAR(20) NULL DEFAULT NULL,
	"nama_bank" VARCHAR(20) NULL DEFAULT NULL,
	"cabang_kcp_unit" VARCHAR(60) NULL DEFAULT NULL,
	"rekening_atas_nama" VARCHAR(50) NULL DEFAULT NULL,
	"mbs" NUMERIC(1,0) NOT NULL,
	"luas_tanah_milik" NUMERIC(7,0) NOT NULL,
	"luas_tanah_bukan_milik" NUMERIC(7,0) NOT NULL,
	"kode_registrasi" BIGINT NULL DEFAULT NULL,
	"npwp" VARCHAR(25) NULL DEFAULT NULL,
	"nm_wp" VARCHAR(100) NULL DEFAULT NULL,
	"keaktifan" NUMERIC(1,0) NOT NULL DEFAULT '1',
	"flag" CHAR(3) NULL DEFAULT NULL,
	"create_date" TIMESTAMP NOT NULL DEFAULT '2020-04-16 09:40:03.422677',
	"last_update" TIMESTAMP NOT NULL DEFAULT '2020-04-16 09:40:03.422677',
	"soft_delete" NUMERIC(1,0) NOT NULL,
	"last_sync" TIMESTAMP NOT NULL DEFAULT '1901-01-01 00:00:00',
	"updater_id" UUID NOT NULL,
	"results" BIGINT NULL DEFAULT NULL,
	"id" TEXT NULL DEFAULT NULL,
	"start" BIGINT NULL DEFAULT NULL,
	"limit" BIGINT NULL DEFAULT NULL,
	PRIMARY KEY ("sekolah_id"),
	INDEX "alamat_kecamatan_fk" ("kode_wilayah"),
	INDEX "sekolah_bentuk_fk" ("bentuk_pendidikan_id"),
	INDEX "sekolah_layanan_kk_fk" ("kebutuhan_khusus_id"),
	UNIQUE INDEX "sekolah_pk" ("sekolah_id"),
	INDEX "sekolah_st_kepemilikan_fk" ("status_kepemilikan_id"),
	INDEX "sp_ls" ("last_sync"),
	INDEX "sp_lu" ("last_update"),
	INDEX "sp_yayasan_fk" ("yayasan_id"),
	CONSTRAINT "fk_sekolah_alamat_ke_mst_wila" FOREIGN KEY ("kode_wilayah") REFERENCES "ref"."mst_wilayah" ("kode_wilayah") ON UPDATE RESTRICT ON DELETE RESTRICT,
	CONSTRAINT "fk_sekolah_sekolah_l_kebutuha" FOREIGN KEY ("kebutuhan_khusus_id") REFERENCES "ref"."kebutuhan_khusus" ("kebutuhan_khusus_id") ON UPDATE RESTRICT ON DELETE RESTRICT,
	CONSTRAINT "sekolah_fk_1" FOREIGN KEY ("bentuk_pendidikan_id") REFERENCES "ref"."bentuk_pendidikan" ("bentuk_pendidikan_id") ON UPDATE RESTRICT ON DELETE RESTRICT,
	CONSTRAINT "sekolah_fk_17" FOREIGN KEY ("yayasan_id") REFERENCES "yayasan" ("yayasan_id") ON UPDATE RESTRICT ON DELETE RESTRICT,
	CONSTRAINT "sekolah_fk_7" FOREIGN KEY ("status_kepemilikan_id") REFERENCES "ref"."status_kepemilikan" ("status_kepemilikan_id") ON UPDATE RESTRICT ON DELETE RESTRICT
)
;
COMMENT ON COLUMN "sekolah"."sekolah_id" IS '';
COMMENT ON COLUMN "sekolah"."nama" IS '';
COMMENT ON COLUMN "sekolah"."nama_nomenklatur" IS '';
COMMENT ON COLUMN "sekolah"."nss" IS '';
COMMENT ON COLUMN "sekolah"."npsn" IS '';
COMMENT ON COLUMN "sekolah"."bentuk_pendidikan_id" IS '';
COMMENT ON COLUMN "sekolah"."alamat_jalan" IS '';
COMMENT ON COLUMN "sekolah"."rt" IS '';
COMMENT ON COLUMN "sekolah"."rw" IS '';
COMMENT ON COLUMN "sekolah"."nama_dusun" IS '';
COMMENT ON COLUMN "sekolah"."desa_kelurahan" IS '';
COMMENT ON COLUMN "sekolah"."kode_wilayah" IS '';
COMMENT ON COLUMN "sekolah"."kode_pos" IS '';
COMMENT ON COLUMN "sekolah"."lintang" IS '';
COMMENT ON COLUMN "sekolah"."bujur" IS '';
COMMENT ON COLUMN "sekolah"."nomor_telepon" IS '';
COMMENT ON COLUMN "sekolah"."nomor_fax" IS '';
COMMENT ON COLUMN "sekolah"."email" IS '';
COMMENT ON COLUMN "sekolah"."website" IS '';
COMMENT ON COLUMN "sekolah"."kebutuhan_khusus_id" IS '';
COMMENT ON COLUMN "sekolah"."status_sekolah" IS '';
COMMENT ON COLUMN "sekolah"."sk_pendirian_sekolah" IS '';
COMMENT ON COLUMN "sekolah"."tanggal_sk_pendirian" IS '';
COMMENT ON COLUMN "sekolah"."status_kepemilikan_id" IS '';
COMMENT ON COLUMN "sekolah"."yayasan_id" IS '';
COMMENT ON COLUMN "sekolah"."sk_izin_operasional" IS '';
COMMENT ON COLUMN "sekolah"."tanggal_sk_izin_operasional" IS '';
COMMENT ON COLUMN "sekolah"."no_rekening" IS '';
COMMENT ON COLUMN "sekolah"."nama_bank" IS '';
COMMENT ON COLUMN "sekolah"."cabang_kcp_unit" IS '';
COMMENT ON COLUMN "sekolah"."rekening_atas_nama" IS '';
COMMENT ON COLUMN "sekolah"."mbs" IS '';
COMMENT ON COLUMN "sekolah"."luas_tanah_milik" IS '';
COMMENT ON COLUMN "sekolah"."luas_tanah_bukan_milik" IS '';
COMMENT ON COLUMN "sekolah"."kode_registrasi" IS '';
COMMENT ON COLUMN "sekolah"."npwp" IS '';
COMMENT ON COLUMN "sekolah"."nm_wp" IS '';
COMMENT ON COLUMN "sekolah"."keaktifan" IS '';
COMMENT ON COLUMN "sekolah"."flag" IS '';
COMMENT ON COLUMN "sekolah"."create_date" IS '';
COMMENT ON COLUMN "sekolah"."last_update" IS '';
COMMENT ON COLUMN "sekolah"."soft_delete" IS '';
COMMENT ON COLUMN "sekolah"."last_sync" IS '';
COMMENT ON COLUMN "sekolah"."updater_id" IS '';
COMMENT ON COLUMN "sekolah"."results" IS '';
COMMENT ON COLUMN "sekolah"."id" IS '';
COMMENT ON COLUMN "sekolah"."start" IS '';
COMMENT ON COLUMN "sekolah"."limit" IS '';
