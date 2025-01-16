package services

import (
	"auth_service/models"
	"auth_service/repository"
)

type SekolahService interface {
	CreateSekolah(sekolah *models.Sekolah) (*models.Sekolah, error)
	GetSekolah(query repository.SekolahQuery) (*models.Sekolah, error)
}

type sekolahServiceImpl struct {
	sekolahRepo repository.SekolahRepository
}

func NewSekolahService(sr repository.SekolahRepository) SekolahService {
	return &sekolahServiceImpl{sekolahRepo: sr}
}

func (ss sekolahServiceImpl) CreateSekolah(s *models.Sekolah) (*models.Sekolah, error) {
	err := ss.sekolahRepo.CreateSekolah(s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// Get Sekolah by NPSN
func (ss *sekolahServiceImpl) GetSekolah(query repository.SekolahQuery) (*models.Sekolah, error) {
	sekolah, err := ss.sekolahRepo.GetSekolah(query)
	if err != nil {
		if err == repository.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return sekolah, nil
	// sekolah, err := ss.sekolahRepo.GetSekolahByNpsn(npsn)
	// if err != nil {
	// 	if errors.Is(err, repository.ErrRecordNotFound) { // Use errors.Is for error comparison
	// 		return nil, ErrNotFound
	// 	}
	// 	return nil, err // Return other errors directly
	// }
	// return sekolah, nil

}