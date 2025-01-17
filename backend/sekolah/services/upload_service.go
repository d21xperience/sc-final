package services

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"sekolah/models"

	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

// UploadData membaca file Excel dan memproses unggahan berdasarkan jenis data
func UploadData(ctx context.Context, filePath, uploadType, schemaName string) (interface{}, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca file Excel: %w", err)
	}
	defer f.Close()

	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data dari sheet: %w", err)
	}

	if len(rows) < 2 {
		return nil, fmt.Errorf("file Excel kosong atau tidak memiliki data yang valid")
	}

	// Pilih proses parsing berdasarkan uploadType
	switch uploadType {
	case "siswa":
		return parseSiswa(rows)
	case "guru":
		return parseGuru(rows)
	case "kelas":
		return parseKelas(rows)
	default:
		return nil, fmt.Errorf("jenis unggahan tidak dikenali: %s", uploadType)
	}
}

// BatchSave menyimpan data model apa pun secara batch
func BatchSave(ctx context.Context, db *gorm.DB, data interface{}, batchSize int) error {
	// Pastikan data tidak kosong
	if data == nil {
		return nil
	}

	// Gunakan `CreateInBatches` untuk batch insert
	err := db.WithContext(ctx).CreateInBatches(data, batchSize).Error
	if err != nil {
		log.Printf("Gagal menyimpan batch data: %v", err)
		return err
	}

	log.Printf("Berhasil menyimpan batch data")
	return nil
}

// Fungsi parsing siswa dari Excel
func parseSiswa(rows [][]string) ([]*models.PesertaDidik, error) {
	var siswaList []*models.PesertaDidik

	for _, row := range rows[1:] { // Lewati header
		if len(row) < 7 {
			log.Println("Skipping row due to insufficient data:", row)
			continue
		}

		nis, err := strconv.Atoi(row[0])
		if err != nil {
			log.Println("Format NIS tidak valid:", row[0])
			continue
		}

		nisn, err := strconv.Atoi(row[1])
		if err != nil {
			log.Println("Format NISN tidak valid:", row[1])
			continue
		}

		siswa := &models.PesertaDidik{
			NIS:          strconv.Itoa(nis),
			NISN:         strconv.Itoa(nisn),
			NamaSiswa:    row[2],
			TempatLahir:  row[3],
			TanggalLahir: row[4],
			JenisKelamin: row[5],
			Agama:        row[6],
		}
		siswaList = append(siswaList, siswa)
	}

	return siswaList, nil
}

// Fungsi parsing siswa dari Excel
func parseGuru(rows [][]string) ([]*models.PesertaDidik, error) {
	var siswaList []*models.PesertaDidik

	for _, row := range rows[1:] { // Lewati header
		if len(row) < 7 {
			log.Println("Skipping row due to insufficient data:", row)
			continue
		}

		nis, err := strconv.Atoi(row[0])
		if err != nil {
			log.Println("Format NIS tidak valid:", row[0])
			continue
		}

		nisn, err := strconv.Atoi(row[1])
		if err != nil {
			log.Println("Format NISN tidak valid:", row[1])
			continue
		}

		siswa := &models.PesertaDidik{
			NIS:          strconv.Itoa(nis),
			NISN:         strconv.Itoa(nisn),
			NamaSiswa:    row[2],
			TempatLahir:  row[3],
			TanggalLahir: row[4],
			JenisKelamin: row[5],
			Agama:        row[6],
		}
		siswaList = append(siswaList, siswa)
	}

	return siswaList, nil
}

// Fungsi parsing siswa dari Excel
func parseKelas(rows [][]string) ([]*models.PesertaDidik, error) {
	var siswaList []*models.PesertaDidik

	for _, row := range rows[1:] { // Lewati header
		if len(row) < 7 {
			log.Println("Skipping row due to insufficient data:", row)
			continue
		}

		nis, err := strconv.Atoi(row[0])
		if err != nil {
			log.Println("Format NIS tidak valid:", row[0])
			continue
		}

		nisn, err := strconv.Atoi(row[1])
		if err != nil {
			log.Println("Format NISN tidak valid:", row[1])
			continue
		}

		siswa := &models.PesertaDidik{
			NIS:          strconv.Itoa(nis),
			NISN:         strconv.Itoa(nisn),
			NamaSiswa:    row[2],
			TempatLahir:  row[3],
			TanggalLahir: row[4],
			JenisKelamin: row[5],
			Agama:        row[6],
		}
		siswaList = append(siswaList, siswa)
	}

	return siswaList, nil
}
