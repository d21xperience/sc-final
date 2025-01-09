package repositories

// import (
// 	"errors"
// 	"sekolah/data/request"
// 	"sekolah/helper"
// 	"sekolah/models"

// 	"gorm.io/gorm"
// )

// type IjazahRepositoryImpl struct {
// 	Db *gorm.DB
// }

// // Buat fungsi dengan parameter yang berisi tipe data sesuai dengan field yang ada pada struct IjazahRepositoryImpl dan mengembalikan nilai berupa interface
// func NewIjazahRepositoryImpl(Db *gorm.DB) IjazahRepository {
// 	return &IjazahRepositoryImpl{Db: Db}
// }

// // implenentasikan semua fungsi yang telah dibuat pada file sebelumnya
// // Create function
// func (ir IjazahRepositoryImpl) Save(tblIjazah models.TabelIjazah) {
// 	result := ir.Db.Create(&tblIjazah)
// 	helper.ErrorPanic(result.Error)
// }

// // Update
// func (ir IjazahRepositoryImpl) Update(tblIjazah models.TabelIjazah) {
// 	var updateTableIjazah = request.UpdateTableIjazah{
// 		NomorIjazah: tblIjazah.NomorIjazah,
// 		AsalSekolah: tblIjazah.AsalSekolah,
// 		Nama:        tblIjazah.Nama,
// 		TglLahir:    tblIjazah.TglLahir,
// 		TptLahir:    tblIjazah.TptLahir,
// 		Nis:         tblIjazah.Nis,
// 		Nisn:        tblIjazah.Nisn,
// 		NamaWali:    tblIjazah.NamaWali,
// 		TahunLulus:  tblIjazah.TahunLulus,
// 		RerataNilai: tblIjazah.RerataNilai,
// 		// valid       : tblIjazah.valid,
// 	}
// 	result := ir.Db.Model(&tblIjazah).Updates(updateTableIjazah)
// 	helper.ErrorPanic(result.Error)
// }

// // Delete
// func (ir IjazahRepositoryImpl) Delete(noIjazah int) {
// 	var iJazah models.TabelIjazah
// 	results := ir.Db.Where("no_ijazah = ?", noIjazah).Delete(&iJazah)
// 	helper.ErrorPanic(results.Error)
// }

// // findById
// func (ir IjazahRepositoryImpl) FindById(noIjazah int) (models.TabelIjazah, error) {
// 	var iJazah models.TabelIjazah
// 	res := ir.Db.Find(&iJazah, noIjazah)
// 	helper.ErrorPanic(res.Error)
// 	if res != nil {
// 		return iJazah, nil
// 	} else {
// 		return iJazah, errors.New("Ijazah tidak ditemukan")
// 	}
// }

// // Find All
// func (ir IjazahRepositoryImpl) FindAll() []models.TabelIjazah {
// 	var Ijazah []models.TabelIjazah
// 	res := ir.Db.Find(&Ijazah)
// 	helper.ErrorPanic(res.Error)
// 	return Ijazah
// }
