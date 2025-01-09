package services

import (
	"sekolah/models"
	"sekolah/repositories"
)

type PesertaDidikService interface {
	Save(PesertaDidik *models.PesertaDidik) error
	FindByID(PesertaDidikID string) (*models.PesertaDidik, error)
	Update(PesertaDidik *models.PesertaDidik) error
	Delete(PesertaDidikID string) error
}

type pesertaDidikServiceImpl struct {
	PesertaDidikRepo repositories.PesertaDidikRepository
}

func NewPesertaDidikService(sr repositories.PesertaDidikRepository) PesertaDidikService {
	return &pesertaDidikServiceImpl{PesertaDidikRepo: sr}
}

func (s pesertaDidikServiceImpl) Save(PesertaDidik *models.PesertaDidik) error {
	return s.Save(PesertaDidik)
}
func (s pesertaDidikServiceImpl) FindByID(PesertaDidikID string) (*models.PesertaDidik, error) {
	return s.PesertaDidikRepo.FindByID(PesertaDidikID)
}
func (s pesertaDidikServiceImpl) Update(PesertaDidik *models.PesertaDidik) error {
	return s.PesertaDidikRepo.Update(PesertaDidik)
}
func (s pesertaDidikServiceImpl) Delete(PesertaDidikID string) error {
	// id, err := uuid.Parse(PesertaDidikID)
	// if err != nil {
	// 	return errors.New("invalid UUID format")
	// }
	return s.PesertaDidikRepo.Delete(PesertaDidikID)
}
