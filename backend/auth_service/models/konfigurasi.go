package models

import "gorm.io/gorm"

type Konfigurasi struct {
	gorm.Model
	SekolahID uint32
	CreateDB bool
	
}