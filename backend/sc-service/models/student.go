package models

import "time"

type Student struct {
	StudentID           uint      `gorm:"primaryKey" json:"student_id"`
	Nama                string    `gorm:"size:100" json:"nama"`
	NIS                 string    `gorm:"unique" json:"nis"`
	NISN                string    `gorm:"unique" json:"nisn"`
	NIK                 string    `gorm:"unique" json:"nik"`
	TptLahir            string    `gorm:"column:tpt_lahir"`
	TglLahir            time.Time `gorm:"column:tgl_lahir"`
	AsalSekolah         string    `gorm:"column:asal_sekolah"`
	UrlBCExplorerEther  string    `gorm:"column:url_bc_explorerEther"`
	UrlBCExplorerQuorum string    `gorm:"column:url_bc_explorerQuorum"`
	UrlBCExplorerFabric string    `gorm:"column:url_bc_explorerFabric"`
	CreatedAt           time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt           time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Certificates []Certificate `gorm:"foreignKey:StudentID" json:"certificates"`
}
