package services

import (
	"auth_service/config"
	pb "auth_service/generated"
	"auth_service/models"
	"auth_service/repositories"
	"auth_service/utils"
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type SekolahIndonesiaServer struct {
	pb.UnimplementedSekolahIndonesiaServiceServer
	repo repositories.GenericRepository[models.SekolahIndonesia]
}

func NewSekolahIndonesiaServer() *SekolahIndonesiaServer {
	// sekolahIndonesia services.SekolahIndonesiaService
	repoSekolahIndonesia := repositories.NewSekolahIndonesiaRepository(config.DB)
	return &SekolahIndonesiaServer{
		repo: *repoSekolahIndonesia,
	}
}

func (s *SekolahIndonesiaServer) GetSekolahIndonesia(ctx context.Context, req *pb.GetSekolahIndonesiaRequest) (*pb.GetSekolahIndonesiaResponse, error) {
	// Validate request
	if req == nil {
		return nil, fmt.Errorf("invalid request")
	}
	npsn := req.GetNpsn()
	namaSekolah := req.GetNamaSekolah()
	var sIList []*pb.SekolahIndonesia
	var sIModel []*models.SekolahIndonesia
	var err error
	if isNumeric(npsn) {
		sIModel, err = s.repo.FindByQuery(ctx, "npsn", npsn)
		if err != nil {
			log.Printf("[ERROR] Gagal menemukan tahun sekolah di schema '%s'", err)
			return nil, fmt.Errorf("gagal menemukan sekolah '%s", err)
		}
		sIList = utils.ConvertModelsToPB(sIModel, func(model *models.SekolahIndonesia) *pb.SekolahIndonesia {
			return &pb.SekolahIndonesia{
				SekolahIdEnkrip: model.SekolahIdEnkrip,
				Kecamatan:       model.Kecamatan,
				Kabupaten:       model.Kabupaten,
				Propinsi:        model.Propinsi,
				KodeKecamatan:   model.KodeKecamatan,
				KodeKab:         model.KodeKab,
				KodeProp:        model.KodeProp,
				NamaSekolah:     model.NamaSekolah,
				Npsn:            model.Npsn,
				AlamatJalan:     model.AlamatJalan,
				Status:          model.Status,
			}
		})
		return &pb.GetSekolahIndonesiaResponse{Pesan: "sukses", SekolahIndonesia: sIList}, nil
	}

	// Jika yang diinputkan adalah nama sekolah
	sIModel, err = s.repo.FindByTextPattern(ctx, "nama_sekolah", namaSekolah)
	if err != nil {
		log.Printf("[ERROR] Gagal menemukan nama sekolah '%s'", err)
		return nil, fmt.Errorf("gagal menemukan nama sekolah '%s", err)
	}
	sIList = utils.ConvertModelsToPB(sIModel, func(model *models.SekolahIndonesia) *pb.SekolahIndonesia {
		return &pb.SekolahIndonesia{
			SekolahIdEnkrip: model.SekolahIdEnkrip,
			Kecamatan:       model.Kecamatan,
			Kabupaten:       model.Kabupaten,
			Propinsi:        model.Propinsi,
			KodeKecamatan:   model.KodeKecamatan,
			KodeKab:         model.KodeKab,
			KodeProp:        model.KodeProp,
			NamaSekolah:     model.NamaSekolah,
			Npsn:            model.Npsn,
			AlamatJalan:     model.AlamatJalan,
			Status:          model.Status,
		}
	})
	return &pb.GetSekolahIndonesiaResponse{Pesan: "sukses", SekolahIndonesia: sIList}, nil

}

func isNumeric(str string) bool {
	// Cek string kosong sebagai invalid
	if str == "" {
		return false
	}

	// Gunakan regex untuk validasi
	re := regexp.MustCompile(`^\d+$`)
	if !re.MatchString(str) {
		return false
	}

	// Gunakan strconv sebagai validasi tambahan
	_, err := strconv.Atoi(str)
	return err == nil
}
