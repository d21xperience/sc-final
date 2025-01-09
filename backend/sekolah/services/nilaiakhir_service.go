package services

import (
	"sekolah/models"
	"sekolah/repositories"
)

type NilaiAkhirService interface {
	Save(NilaiAkhir *models.TabelNilaiAkhir) error
	FindByID(NilaiAkhirID string) (*models.TabelNilaiAkhir, error)
	Update(NilaiAkhir *models.TabelNilaiAkhir) error
	Delete(NilaiAkhirID string) error
}

type NilaiAkhirServiceImpl struct {
	NilaiAkhirRepo repositories.NilaiAkhirRepository
}

func NewNilaiAkhirService(sr repositories.NilaiAkhirRepository) NilaiAkhirService {
	return &NilaiAkhirServiceImpl{NilaiAkhirRepo: sr}
}

func (s NilaiAkhirServiceImpl) Save(NilaiAkhir *models.TabelNilaiAkhir) error {
	return s.Save(NilaiAkhir)
}
func (s NilaiAkhirServiceImpl) FindByID(NilaiAkhirID string) (*models.TabelNilaiAkhir, error) {
	return s.NilaiAkhirRepo.FindByID(NilaiAkhirID)
}
func (s NilaiAkhirServiceImpl) Update(NilaiAkhir *models.TabelNilaiAkhir) error {
	return s.NilaiAkhirRepo.Update(NilaiAkhir)
}
func (s NilaiAkhirServiceImpl) Delete(NilaiAkhirID string) error {
	// id, err := uuid.Parse(NilaiAkhirID)
	// if err != nil {
	// 	return errors.New("invalid UUID format")
	// }
	return s.NilaiAkhirRepo.Delete(NilaiAkhirID)
}
