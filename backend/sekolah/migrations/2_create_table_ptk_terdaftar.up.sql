CREATE TABLE "ptk_terdaftar" (
	"ptk_terdaftar_id" UUID NOT NULL,
	"ptk_id" UUID NOT NULL,
	"sekolah_id" UUID NOT NULL,
	"jenis_keluar_id" CHAR(1) NULL DEFAULT NULL,
	"tahun_ajaran_id" NUMERIC(4,0) NOT NULL,
	"jenis_ptk_id" NUMERIC(2,0) NOT NULL,
	"nomor_surat_tugas" VARCHAR(80) NOT NULL,
	"tanggal_surat_tugas" DATE NOT NULL,
	"ptk_induk" NUMERIC(1,0) NOT NULL,
	"tmt_tugas" DATE NOT NULL,
	"aktif_bulan_01" NUMERIC(1,0) NOT NULL DEFAULT '0',
	"aktif_bulan_02" NUMERIC(1,0) NOT NULL DEFAULT '0',
	"aktif_bulan_03" NUMERIC(1,0) NOT NULL DEFAULT '0',
	"aktif_bulan_04" NUMERIC(1,0) NOT NULL DEFAULT '0',
	"aktif_bulan_05" NUMERIC(1,0) NOT NULL DEFAULT '0',
	"aktif_bulan_06" NUMERIC(1,0) NOT NULL DEFAULT '0',
	"aktif_bulan_07" NUMERIC(1,0) NOT NULL DEFAULT '0',
	"aktif_bulan_08" NUMERIC(1,0) NOT NULL DEFAULT '0',
	"aktif_bulan_09" NUMERIC(1,0) NOT NULL DEFAULT '0',
	"aktif_bulan_10" NUMERIC(1,0) NOT NULL DEFAULT '0',
	"aktif_bulan_11" NUMERIC(1,0) NOT NULL DEFAULT '0',
	"aktif_bulan_12" NUMERIC(1,0) NOT NULL DEFAULT '0',
	"tgl_ptk_keluar" DATE NULL DEFAULT NULL,
	"create_date" TIMESTAMP NOT NULL DEFAULT '2022-06-28 18:43:46.244',
	"last_update" TIMESTAMP NOT NULL DEFAULT '2022-06-28 18:43:46.244',
	"soft_delete" NUMERIC(1,0) NOT NULL,
	"last_sync" TIMESTAMP NOT NULL DEFAULT '1901-01-01 00:00:00',
	"updater_id" UUID NOT NULL,
	PRIMARY KEY ("ptk_terdaftar_id"),
	INDEX "ptk_keluar_fk" ("jenis_keluar_id"),
	UNIQUE INDEX "ptk_sekolah_unique" ("ptk_id", "sekolah_id", "tahun_ajaran_id"),
	INDEX "ptk_terdaftar_mulai_fk" ("tahun_ajaran_id"),
	UNIQUE INDEX "ptk_terdaftar_pk" ("ptk_terdaftar_id"),
	INDEX "ptk_terdaftar_sekolah_fk" ("sekolah_id"),
	CONSTRAINT "fk_ptk_terd_ptk_jenis_jenis_pt" FOREIGN KEY ("jenis_ptk_id") REFERENCES "ref"."jenis_ptk" ("jenis_ptk_id") ON UPDATE RESTRICT ON DELETE RESTRICT,
	CONSTRAINT "fk_ptk_terd_ptk_kelua_jenis_ke" FOREIGN KEY ("jenis_keluar_id") REFERENCES "ref"."jenis_keluar" ("jenis_keluar_id") ON UPDATE RESTRICT ON DELETE RESTRICT,
	CONSTRAINT "ptk_terdaftar_fk_1" FOREIGN KEY ("ptk_id") REFERENCES "ptk" ("ptk_id") ON UPDATE RESTRICT ON DELETE RESTRICT,
	CONSTRAINT "ptk_terdaftar_fk_3" FOREIGN KEY ("tahun_ajaran_id") REFERENCES "ref"."tahun_ajaran" ("tahun_ajaran_id") ON UPDATE RESTRICT ON DELETE RESTRICT,
	CONSTRAINT "ptk_terdaftar_fk_4" FOREIGN KEY ("sekolah_id") REFERENCES "sekolah" ("sekolah_id") ON UPDATE RESTRICT ON DELETE RESTRICT
)
;
COMMENT ON COLUMN "ptk_terdaftar"."ptk_terdaftar_id" IS '';
COMMENT ON COLUMN "ptk_terdaftar"."ptk_id" IS '';
COMMENT ON COLUMN "ptk_terdaftar"."sekolah_id" IS '';
COMMENT ON COLUMN "ptk_terdaftar"."jenis_keluar_id" IS '';
COMMENT ON COLUMN "ptk_terdaftar"."tahun_ajaran_id" IS '';
COMMENT ON COLUMN "ptk_terdaftar"."jenis_ptk_id" IS '';
COMMENT ON COLUMN "ptk_terdaftar"."nomor_surat_tugas" IS '';
COMMENT ON COLUMN "ptk_terdaftar"."tanggal_surat_tugas" IS '';
COMMENT ON COLUMN "ptk_terdaftar"."ptk_induk" IS '';
COMMENT ON COLUMN "ptk_terdaftar"."tmt_tugas" IS '';
COMMENT ON COLUMN "ptk_terdaftar"."aktif_bulan_01" IS '';
COMMENT ON COLUMN "ptk_terdaftar"."aktif_bulan_02" IS '';
COMMENT ON COLUMN "ptk_terdaftar"."aktif_bulan_03" IS '';
COMMENT ON COLUMN "ptk_terdaftar"."aktif_bulan_04" IS '';
COMMENT ON COLUMN "ptk_terdaftar"."aktif_bulan_05" IS '';
COMMENT ON COLUMN "ptk_terdaftar"."aktif_bulan_06" IS '';
COMMENT ON COLUMN "ptk_terdaftar"."aktif_bulan_07" IS '';
COMMENT ON COLUMN "ptk_terdaftar"."aktif_bulan_08" IS '';
COMMENT ON COLUMN "ptk_terdaftar"."aktif_bulan_09" IS '';
COMMENT ON COLUMN "ptk_terdaftar"."aktif_bulan_10" IS '';
COMMENT ON COLUMN "ptk_terdaftar"."aktif_bulan_11" IS '';
COMMENT ON COLUMN "ptk_terdaftar"."aktif_bulan_12" IS '';
COMMENT ON COLUMN "ptk_terdaftar"."tgl_ptk_keluar" IS '';
COMMENT ON COLUMN "ptk_terdaftar"."create_date" IS '';
COMMENT ON COLUMN "ptk_terdaftar"."last_update" IS '';
COMMENT ON COLUMN "ptk_terdaftar"."soft_delete" IS '';
COMMENT ON COLUMN "ptk_terdaftar"."last_sync" IS '';
COMMENT ON COLUMN "ptk_terdaftar"."updater_id" IS '';
