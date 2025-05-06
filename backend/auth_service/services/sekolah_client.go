package services

import (
	pb "auth_service/generated/sekolah"
	"auth_service/models"
	"auth_service/utils"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SekolahServiceClient struct {
	client pb.SekolahServiceClient
}

func NewSekolahServiceClient() (*SekolahServiceClient, error) {
	sekolahClient := os.Getenv("GRPC_SEKOLAHURI")
	conn, err := grpc.NewClient(sekolahClient, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("gagal terhubung ke sekolah_service: %w", err)
	}

	client := pb.NewSekolahServiceClient(conn)
	return &SekolahServiceClient{client: client}, nil
}

func (s *SekolahServiceClient) RegistrasiSekolah(sekolah *models.SekolahTenant) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	_, err := s.client.RegistrasiSekolah(ctx, &pb.TabelSekolahRequest{
		Sekolah: &pb.Sekolah{
			EnkripId:        sekolah.EnkripID,
			SekolahTenantId: sekolah.ID,
			NamaSekolah:     sekolah.NamaSekolah,
			Kecamatan:       utils.SafeString(sekolah.Kecamatan),
			Kabupaten:       utils.SafeString(sekolah.Kabupaten),
			Propinsi:        utils.SafeString(sekolah.Propinsi),
			KodeKecamatan:   utils.SafeString(sekolah.KodeKecamatan),
			KodeKab:         utils.SafeString(sekolah.KodeKab),
			KodeProp:        utils.SafeString(sekolah.KodeProp),
			Npsn:            sekolah.NPSN,
			AlamatJalan:     utils.SafeString(sekolah.AlamatJalan),
		},
	})
	if err != nil {
		return fmt.Errorf("gagal mendaftarkan sekolah di sekolah_service: %w", err)
	}

	log.Printf("Schema sekolah %s berhasil dibuat di sekolah_service", sekolah.EnkripID)
	return nil
}
func (s *SekolahServiceClient) CreateSekolah(sekolah *models.SekolahTenant) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.client.CreateSekolah(ctx, &pb.CreateSekolahRequest{
		Schemaname: fmt.Sprintf("tabel_%s", sekolah.EnkripID),
		Sekolah: &pb.SekolahDapo{
			SekolahId: sekolah.EnkripID,
			Nama:      sekolah.NamaSekolah,
			Npsn:      sekolah.NPSN,
			Kecamatan: utils.SafeString(sekolah.Kecamatan),
			KabKota:   utils.SafeString(sekolah.Kabupaten),
			Propinsi:  utils.SafeString(sekolah.Propinsi),
			Alamat:    utils.SafeString(sekolah.AlamatJalan),
		},
	})
	if err != nil {
		return fmt.Errorf("gagal mendaftarkan sekolah di sekolah_service: %w", err)
	}
	return nil
}
