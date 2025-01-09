package services

import (
	"sekolah/data/request"
)

type IjazahService interface {
	Create(ijazah request.CreateIjazahRequest)
	Update(ijazah request.UpdateIjazahRequest)
	Delete(noIjazah int)
	// FindById(noIjazah int)
	// FindAll() []response.IjazahResponse
}
type IjazahServiceImpl struct {
	// ID          uuid.UUID
	// Nama        string
	// AsalSekolah string
	// RerataNilai string
	Ija *IjazahService
}

func NewIjazahService(ijazahServ IjazahServiceImpl) IjazahService {
	return IjazahServiceImpl{ijazahServ.Ija}
}

func (ij IjazahServiceImpl) Create(ijazah request.CreateIjazahRequest) {

}

func (ij IjazahServiceImpl) Update(ijazah request.UpdateIjazahRequest) {

}

func (ij IjazahServiceImpl) Delete(noIjazah int) {

}
