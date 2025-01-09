package services

import (
	"sekolah/models"
	"sekolah/repositories"
)

type SekolahService interface {
	Save(sekolah *models.Sekolah) error
	FindByID(sekolahID string) (*models.Sekolah, error)
	Update(sekolah *models.Sekolah) error
	Delete(sekolahID string) error
}

type sekolahServiceImpl struct {
	SekolahRepo repositories.SekolahRepository
}

func NewSekolahService(sr repositories.SekolahRepository) SekolahService {
	return &sekolahServiceImpl{SekolahRepo: sr}
}

func (s sekolahServiceImpl) Save(sekolah *models.Sekolah) error {
	return s.Save(sekolah)
}
func (s sekolahServiceImpl) FindByID(sekolahID string) (*models.Sekolah, error) {
	return s.SekolahRepo.FindByID(sekolahID)
}
func (s sekolahServiceImpl) Update(sekolah *models.Sekolah) error {
	return s.SekolahRepo.Update(sekolah)
}
func (s sekolahServiceImpl) Delete(sekolahID string) error {
	// id, err := uuid.Parse(sekolahID)
	// if err != nil {
	// 	return errors.New("invalid UUID format")
	// }
	return s.SekolahRepo.Delete(sekolahID)
}
