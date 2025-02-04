package models

import "gorm.io/gorm"

type SekolahTenant struct {
	gorm.Model
	NamaSekolah string
	UserId      int32
	SekolahId   int32
	SchemaName  string `gorm:"unique;not null"`
}

type SchemaLog struct {
	gorm.Model
	SchemaName string
}
