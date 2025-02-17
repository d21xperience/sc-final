package services

import (
	pb "auth_service/generated/sekolah"
	"auth_service/models"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SekolahServiceClient struct {
	client pb.SekolahServiceClient
}

func NewSekolahServiceClient() (*SekolahServiceClient, error) {
	conn, err := grpc.NewClient("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("gagal terhubung ke sekolah_service: %w", err)
	}

	client := pb.NewSekolahServiceClient(conn)
	return &SekolahServiceClient{client: client}, nil
}

func (s *SekolahServiceClient) RegistrasiSekolah(sekolah *models.Sekolah) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.client.RegistrasiSekolah(ctx, &pb.TabelSekolahRequest{
		Sekolah: &pb.Sekolah{
			SekolahIdEnkrip: sekolah.SekolahIDEnkrip,
			SekolahId:       int32(sekolah.ID),
			NamaSekolah:     sekolah.NamaSekolah,
		},
	})
	if err != nil {
		return fmt.Errorf("gagal mendaftarkan sekolah di sekolah_service: %w", err)
	}

	log.Printf("Schema sekolah %s berhasil dibuat di sekolah_service", sekolah.SekolahIDEnkrip)
	return nil
}
func (s *SekolahServiceClient) CreateSekolah(sekolah *models.Sekolah) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.client.CreateSekolah(ctx, &pb.CreateSekolahRequest{
		SchemaName: fmt.Sprintf("tabel_%s", sekolah.SekolahIDEnkrip),
		Sekolah: &pb.SekolahDapo{
			Alamat:    sekolah.AlamatJalan,
			Npsn:      sekolah.NPSN,
			Nama:      sekolah.NamaSekolah,
			Kecamatan: sekolah.Kecamatan,
			KabKota:   sekolah.Kabupaten,
			Propinsi:  sekolah.Propinsi,
			// // StatusKepemilikanId: 0,
			// BentukPendidikanId: 0,
		},
	})
	if err != nil {
		return fmt.Errorf("gagal mendaftarkan sekolah di sekolah_service: %w", err)
	}
	return nil
}
