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

	// sendPtk(cfg, 2020, 2021)
	// sendSiswa(cfg, 2020, 2021)
	sendRombel(cfg, 2020, 2023)
}

func sendSekolah(cfg *config.AppConfig, tahunBerapa, keBerapa uint16) {
	for year := 2022; year <= 2022; year++ {
		for semester := 1; semester <= 1; semester++ {
			semesterID := fmt.Sprintf("%d%d", year, semester)

			err := usecase.ProcessSekolah(cfg, semesterID)
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

func sendPtk(cfg *config.AppConfig, tahunBerapa, keBerapa uint16) {
	for year := 2022; year <= 2022; year++ {
		// for semester := 1; semester <= 2; semester++ {
		semesterID := fmt.Sprintf("%d", year)

		err := usecase.ProcessPtk(cfg, semesterID)
		if err != nil {
			fmt.Printf("Gagal memproses semester %s: %v\n", semesterID, err)
			continue
		}
		fmt.Printf("Berhasil memproses semester %s\n", semesterID)
		// }
		// Tambahkan delay antar semester (misalnya 2 detik)
		time.Sleep(2 * time.Second)
	}
}

func sendRombel(cfg *config.AppConfig, tahunBerapa, keBerapa uint16) {
	for year := tahunBerapa; year <= keBerapa; year++ {
		for semester := 1; semester <= 2; semester++ {
			semesterID := fmt.Sprintf("%d%d", year, semester)

			err := usecase.ProcessRombel(cfg, semesterID)
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

func sendSiswa(cfg *config.AppConfig, tahunBerapa, keBerapa uint16) {
	for year := tahunBerapa; year <= keBerapa; year++ {
		for semester := 1; semester <= 2; semester++ {
			semesterID := fmt.Sprintf("%d%d", year, semester)

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
