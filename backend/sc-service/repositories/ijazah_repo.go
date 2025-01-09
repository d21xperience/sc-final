package repositories

// import (
// 	"sc-service/data/request"
// 	"sc-service/model"

// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// )

// // Create
// type IjazahRepository interface {
// 	Save(ij model.TabelIjazah)
// 	Update(ij model.TabelIjazah)
// 	Delete(ij uuid.UUID)
// 	// FindByID(ijID int) (model.TabelIjazah, error)
// 	FindAll(thnLulus int) []model.TabelIjazah
// }

// type IjazahRepositoryImpl struct {
// 	Db *gorm.DB
// }

// func NewIjazahImpl(dbs *gorm.DB) IjazahRepository {
// 	return &IjazahRepositoryImpl{Db: dbs}
// }

// func (dbs IjazahRepositoryImpl) Save(ij model.TabelIjazah) {
// 	res := dbs.Db.Create(&ij)
// 	if res != nil {
// 		panic(res.Error)
// 	}
// }
// func (dbs IjazahRepositoryImpl) Update(ij model.TabelIjazah) {
// 	var updateIj = request.CreateTagsRequest{
// 		AsalSekolah: ij.AsalSekolah,
// 		Nama:        ij.Nama,
// 		TptLahir:    ij.TptLahir,
// 		TglLahir:    ij.TglLahir,
// 		NIS:         ij.NIS,
// 		NISN:        ij.NISN,
// 		NamaWali:    ij.NamaWali,
// 		TahunLulus:  ij.TahunLulus,
// 		RerataNilai: ij.RerataNilai,
// 	}
// 	res := dbs.Db.Model(&ij).Updates(updateIj)
// 	if res != nil {
// 		panic(res.Error)
// 	}
// }
// func (dbs IjazahRepositoryImpl) Delete(ijID uuid.UUID) {
// 	var ij model.TabelIjazah
// 	res := dbs.Db.Where("pesertadidik_id = ?", ijID).Delete(&ij)
// 	if res != nil {
// 		panic(res.Error)
// 	}
// }

// // func (dbs IjazahRepositoryImpl) FindByID(ijID uui) (model.TabelIjazah, error) {
// // 	var ij
// // }

// func (dbs IjazahRepositoryImpl) FindAll(thnLulus int) []model.TabelIjazah {
// 	var ij []model.TabelIjazah
// 	res := dbs.Db.Find(&ij).Where("thn_lulus =?", thnLulus)
// 	if res != nil {
// 		panic(res.Error)
// 	}
// 	return ij
// }
