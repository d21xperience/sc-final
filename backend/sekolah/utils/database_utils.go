package utils

// import (
// 	"fmt"
// 	"sekolah/models"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// func CreateSchoolDatabase(schoolName string) error {
// 	dsn := fmt.Sprintf("host=localhost user=postgres password=secret dbname=%s sslmode=disable", schoolName)
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return err
// 	}

// 	// Initialize schema for the school
// 	if err := db.AutoMigrate(&models.PesertaDidik{}, &models.TabelNilaiAkhir{}); err != nil {
// 		return err
// 	}

// 	return nil
// }
