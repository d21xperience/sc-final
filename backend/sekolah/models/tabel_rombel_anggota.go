package models

import "github.com/google/uuid"

type RombelAnggota struct {
	AnggotaRombelID    uuid.UUID        `json:"anggota_rombel_id"` // UUID
	PesertaDidikID     PesertaDidik     `json:"peserta_didik_id"`  // UUID
	RombonganBelajarID RombonganBelajar `json:"rombongan_belajar_id"`
	SemesterID         Semester         `json:"semester_id"`
}
