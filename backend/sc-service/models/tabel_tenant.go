package models

import "gorm.io/gorm"

type SekolahTenant struct {
	gorm.Model
	NamaSekolah     string
	UserId          int32
	SekolahTenantId uint32
	SchemaName      string `gorm:"unique;not null"`
}

type SchemaLog struct {
	gorm.Model
	SchemaName string
}
