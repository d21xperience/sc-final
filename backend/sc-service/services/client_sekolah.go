package services

import (
	"context"
	"fmt"
	"log"
	"os"
	pb "sc-service/generated/sekolah"
	"sc-service/models"
	"sc-service/utils"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SekolahServiceClient struct {
	client      pb.SiswaServiceClient
	infoSekolah pb.SekolahServiceClient
}

func NewSekolahServiceClient() (*SekolahServiceClient, error) {
	grpcSekolah := os.Getenv("GRPC_SEKOLAHURI")
	conn, err := grpc.NewClient(grpcSekolah, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("gagal terhubung ke sekolah_service: %w", err)
	}

	// client := pb.NewSekolahServiceClient(conn)
	client := pb.NewSiswaServiceClient(conn)
	return &SekolahServiceClient{client: client}, nil
}

func (s *SekolahServiceClient) SearchSiswaById(schemaname, pesertaDidikId string) models.Student {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	result, err := s.client.SearchSiswa(ctx, &pb.SearchSiswaRequest{
		Schemaname:     schemaname,
		PesertaDidikId: pesertaDidikId,
	})
	if err != nil {
		log.Printf("gagal mendapatkan siswa: %s", err)
		return models.Student{}
	}
	response := utils.ConvertModelToPB(result, func(item *pb.SearchSiswaResponse) *models.Student {
		tglLahir, err := utils.StringToTime(item.Siswa.TanggalLahir, "2006-01-02")
		if err != nil {
			return nil
		}
		return &models.Student{
			PesertaDidikId: utils.StringToUUID(item.Siswa.PesertaDidikId),
			Nama:           item.Siswa.NmSiswa,
			NIS:            item.Siswa.Nis,
			NISN:           item.Siswa.Nisn,
			TptLahir:       item.Siswa.TempatLahir,
			TglLahir:       tglLahir,
		}
	})
	if response == nil {
		return models.Student{}
	}
	return *response
}

func (s *SekolahServiceClient) SearchSekolah(schemaname string) models.Sekolah {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	infoSekolah, err := s.infoSekolah.GetSekolah(ctx, &pb.GetSekolahRequest{
		Schemaname: schemaname,
	})
	if err != nil {
		log.Printf("gagal mengambil data sekolah di sekolah_service: %s", err)
		return models.Sekolah{}
	}
	response := utils.ConvertModelToPB(infoSekolah, func(item *pb.GetSekolahResponse) *models.Sekolah {
		return &models.Sekolah{
			SekolahID: item.Sekolah.SekolahId,
			Nama:      item.Sekolah.Nama,
			KabKota:   item.Sekolah.KabKota,
			Propinsi:  item.Sekolah.Propinsi,
			Npsn:      item.Sekolah.Npsn,
		}
	})

	if response == nil {
		return models.Sekolah{}
	}
	return *response
}
