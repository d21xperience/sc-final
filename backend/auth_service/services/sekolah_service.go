package services

import (
	"auth_service/models"
	"auth_service/repository"
)

type SekolahService interface {
	CreateSekolah(sekolah *models.Sekolah) (*models.Sekolah, error)
	GetSekolahByNpsn(npsn string) (*models.Sekolah, error)
}

type sekolahServiceImpl struct {
	SekolahRepo repository.SekolahRepository
}

func NewSekolahService(sr repository.SekolahRepository) SekolahService {
	return &sekolahServiceImpl{SekolahRepo: sr}
}

func (ss sekolahServiceImpl) CreateSekolah(s *models.Sekolah) (*models.Sekolah, error) {
	err := ss.SekolahRepo.CreateSekolah(s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// Get Sekolah by NPSN
func (ss *sekolahServiceImpl) GetSekolahByNpsn(npsn string) (*models.Sekolah, error) {
	sekolah, err := ss.SekolahRepo.GetSekolahByNpsn(npsn)
	if err != nil {
		if err == repository.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return sekolah, nil
}

// // Get Sekolah by NPSN
// func (ss *sekolahServiceImpl) CreateSekolahByNpsn(npsn string) (*models.Sekolah, error) {
// 	sekolah, err := ss.SekolahRepo.GetSekolahByNpsn(npsn)
// 	if err != nil {
// 		return nil, errors.New("sekolah tidak ditemukan")
// 	}
// 	return sekolah, nil
// }
