package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewDnsRepository(db *gorm.DB) *GenericRepository[models.DataNominasiSementara] {
	return NewGenericRepository[models.DataNominasiSementara](db, "data_nominasi_sementara")
}
