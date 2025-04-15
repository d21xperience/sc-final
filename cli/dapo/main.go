package main

import (
	"dapo/config"
	"dapo/usecase"
	"fmt"
	"time"
)

func main() {
	// Load configuration from .env
	cfg := config.LoadConfig()
	// SemesterID := "20212"
	// // Inisialisasi API Client
	// client := services.NewAPIClient(cfg.BaseURL, cfg.Token)

	// // Ambil data sekolah (bisa ganti endpoint/data lainnya)
	// sekolahData, err := services.GetOrFetch[models.SekolahResponse](
	// 	client,
	// 	cfg.NPSN,
	// 	SemesterID,
	// 	cfg.OutputDir,
	// 	"/WebService/getSekolah",
	// 	map[string]string{
	// 		"npsn":        cfg.NPSN,
	// 		"semester_id": SemesterID,
	// 	},
	// )
	// if err != nil {
	// 	fmt.Printf("Error fetching sekolah data: %v\n", err)
	// 	return
	// }

	// // Inisialisasi gRPC Client
	// gRPCClient, err := services.NewGRPCSekolahClient(cfg.GRPCHost, cfg.GRPCPort, cfg.GRPCTimeout)
	// if err != nil {
	// 	fmt.Printf("gRPC Connection Error: %v\n", err)
	// 	return
	// }
	// defer gRPCClient.Close()

	// // Kirim data ke gRPC
	// modSekolah := models.Sekolah{
	// 	SekolahID: sekolahData.Rows.SekolahID,
	// 	NSS:       sekolahData.Rows.NSS,
	// 	// NPSN:     sekolahData.Rows.NPSN,
	// 	NomorFax: sekolahData.Rows.NomorFax,
	// 	Website:  sekolahData.Rows.Website,
	// 	Email:    sekolahData.Rows.Email,
	// }
	// resp, err := gRPCClient.SendSekolahData(&modSekolah)
	// if err != nil {
	// 	fmt.Printf("gRPC Error: %v\n", err)
	// 	return
	// }
	// fmt.Printf("gRPC Response: %v\n", resp)
	// for year := 2020; year <= 2020; year++ {
	// 	for semester := 1; semester <= 2; semester++ {
	// 		semesterID := fmt.Sprintf("%d%d", year, semester)

	// 		err := usecase.ProcessSekolah(cfg, semesterID)
	// 		if err != nil {
	// 			fmt.Printf("Gagal memproses semester %s: %v\n", semesterID, err)
	// 			continue
	// 		}
	// 		fmt.Printf("Berhasil memproses semester %s\n", semesterID)
	// 	}
	// 	// Tambahkan delay antar semester (misalnya 2 detik)
	// 	time.Sleep(2 * time.Second)
	// }
	// for year := 2022; year <= 2022; year++ {
	// 	// for semester := 1; semester <= 2; semester++ {
	// 	semesterID := fmt.Sprintf("%d", year)

	// 	err := usecase.ProcessPtk(cfg, semesterID)
	// 	if err != nil {
	// 		fmt.Printf("Gagal memproses semester %s: %v\n", semesterID, err)
	// 		continue
	// 	}
	// 	fmt.Printf("Berhasil memproses semester %s\n", semesterID)
	// 	// }
	// 	// Tambahkan delay antar semester (misalnya 2 detik)
	// 	time.Sleep(2 * time.Second)
	// }
	for year := 2022; year <= 2022; year++ {
		for semester := 1; semester <= 2; semester++ {
			semesterID := fmt.Sprintf("%d", year)

			err := usecase.ProcessSiswa(cfg, semesterID)
			if err != nil {
				fmt.Printf("Gagal memproses semester %s: %v\n", semesterID, err)
				continue
			}
			fmt.Printf("Berhasil memproses semester %s\n", semesterID)
		}
		// Tambahkan delay antar semester (misalnya 2 detik)
		time.Sleep(2 * time.Second)
	}

}
