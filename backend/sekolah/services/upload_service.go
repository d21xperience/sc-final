package services

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
)

type UploadHandler struct {
	service *PesertaDidikService
}

// NewUploadHandler membuat instance baru UploadHandler.
func NewUploadHandler(service *PesertaDidikService) *UploadHandler {
	return &UploadHandler{service: service}
}

// HandleUpload adalah HTTP handler untuk upload file.
// func (h *UploadHandler) HandleUpload(w http.ResponseWriter, r *http.Request) {
func (h *UploadHandler) HandleBinaryFileUpload(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	// Parse form data
	// Parse form data
	err := r.ParseMultipartForm(10 << 20) // Batas ukuran file 10MB
	if err != nil {
		http.Error(w, "Gagal mem-parsing form data", http.StatusBadRequest)
		return
	}

	// Ambil parameter dan file
	// uploadType := r.FormValue("upload_type")
	fileHeader := r.MultipartForm.File["file"]
	if len(fileHeader) == 0 {
		http.Error(w, "File tidak ditemukan", http.StatusBadRequest)
		return
	}
	file, err := fileHeader[0].Open()
	if err != nil {
		http.Error(w, "Gagal membuka file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Validasi file Excel
	fileName := fileHeader[0].Filename
	if !strings.HasSuffix(fileName, ".xlsx") {
		http.Error(w, "Hanya file Excel (.xlsx) yang diizinkan", http.StatusBadRequest)
		return
	}

	// Simpan file sementara
	tempFile, err := os.CreateTemp("", "upload-*.xlsx")
	if err != nil {
		http.Error(w, "Gagal membuat file sementara", http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	_, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(w, "Gagal menyimpan file sementara", http.StatusInternalServerError)
		return
	}

	// Proses file (gunakan service jika diperlukan)
	// data, err := h.sekolahService.ProcessExcel(context.Background(), tempFile.Name(), uploadType)
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Gagal memproses file: %v", err), http.StatusInternalServerError)
	// 	return
	// }

	// Berikan respon
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "File berhasil diproses",
		// "data":    data,
	})
}

// import (
// 	"context"
// 	"fmt"
// 	"io"
// 	"os"
// 	"sekolah/models"

// 	"github.com/xuri/excelize/v2"
// )

// type Repository interface {
// 	UploadDataSekolah(ctx context.Context, tempFilePath, uploadType, schemaName string) error
// }
// type UploadService struct {
// 	repo Repository
// }

// func NewUploadDataSekolah(repo Repository) *UploadService {
// 	return &UploadService{repo: repo}
// }

// // ProcessUpload menangani proses upload dan penyimpanan data.
// func (s *UploadService) ProcessUpload(ctx context.Context, file io.Reader, uploadType string) error {
// 	// Simpan file ke lokasi sementara
// 	tempFile, err := os.CreateTemp("", "upload-*.xlsx")
// 	if err != nil {
// 		return fmt.Errorf("gagal membuat file sementara: %w", err)
// 	}
// 	defer tempFile.Close()
// 	defer os.Remove(tempFile.Name()) // Hapus file setelah selesai

// 	// Salin konten file yang diunggah ke file sementara
// 	if _, err := io.Copy(tempFile, file); err != nil {
// 		return fmt.Errorf("gagal menyimpan file sementara: %w", err)
// 	}

// 	// Proses file Excel
// 	data, err := s.parseExcel(tempFile.Name(), uploadType)
// 	if err != nil {
// 		return fmt.Errorf("gagal memproses file Excel: %w", err)
// 	}

// 	// Simpan data ke database
// 	if err := s.repo.SaveData(ctx, data); err != nil {
// 		return fmt.Errorf("gagal menyimpan data ke database: %w", err)
// 	}

// 	return nil
// }
// // parseExcel membaca dan memproses file Excel berdasarkan uploadType.
// func (s *UploadService) parseExcel(filePath, uploadType string) (interface{}, error) {
// 	f, err := excelize.OpenFile(filePath)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal membuka file Excel: %w", err)
// 	}
// 	defer f.Close()

// 	rows, err := f.GetRows(f.GetSheetName(0))
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal membaca data dari sheet: %w", err)
// 	}

// 	switch uploadType {
// 	case "siswa":
// 		return parseSiswa(rows), nil
// 	case "guru":
// 		return parseGuru(rows), nil
// 	case "kelas":
// 		return parseKelas(rows), nil
// 	case "nilaiAkhir":
// 		return parseNilaiAkhir(rows), nil
// 	default:
// 		return nil, fmt.Errorf("jenis unggahan tidak dikenali: %s", uploadType)
// 	}
// }

// // Helper functions for parsing (contoh sederhana)
// func parseSiswa(rows [][]string) []map[string]string {
// 	var result []map[string]string
// 	for _, row := range rows[1:] { // Skip header
// 		if len(row) < 2 {
// 			continue
// 		}
// 		result = append(result, map[string]string{
// 			"Nama": row[0],
// 			"NIS":  row[1],
// 		})
// 	}
// 	return result
// }

// Fungsi generik untuk membaca file Excel dan memproses data berdasarkan jenis
// func uploadData[T any](ctx context.Context, filePath, uploadType, schemaName string) ([]T, error) {
// func uploadData[T any](filePath, uploadType string) ([]T, error) {
// 	f, err := excelize.OpenFile(filePath)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal membaca file Excel: %w", err)
// 	}
// 	defer f.Close()

// 	rows, err := f.GetRows(f.GetSheetName(0))
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal mengambil data dari sheet: %w", err)
// 	}

// 	if len(rows) < 2 {
// 		return nil, fmt.Errorf("file Excel kosong atau tidak memiliki data yang valid")
// 	}

// 	// Parsing data berdasarkan uploadType
// 	switch uploadType {
// 	case "siswa":
// 		if v, ok := any(parseSiswa(rows)).([]T); ok {
// 			return v, nil
// 		}
// 	case "guru":
// 		if v, ok := any(parseGuru(rows)).([]T); ok {
// 			return v, nil
// 		}
// 	case "kelas":
// 		if v, ok := any(parseKelas(rows)).([]T); ok {
// 			return v, nil
// 		}
// 	case "nilaiAkhir":
// 		if v, ok := any(parseNilaiAkhir(rows)).([]T); ok {
// 			return v, nil
// 		}
// 	default:
// 		return nil, fmt.Errorf("jenis unggahan tidak dikenali: %s", uploadType)
// 	}

// 	return nil, fmt.Errorf("gagal memproses data dengan tipe yang diberikan")
// }

// // Fungsi parsing untuk siswa
// func parseSiswa(rows [][]string) []*models.PesertaDidik {
// 	var siswaList []*models.PesertaDidik
// 	for _, row := range rows[1:] {
// 		if len(row) < 3 {
// 			continue
// 		}
// 		siswa := &models.PesertaDidik{
// 			PesertaDidikID:  row[0],
// 			Nis:             row[1],
// 			Nisn:            row[2],
// 			NmSiswa:         row[3],
// 			TempatLahir:     row[4],
// 			TanggalLahir:    row[5],
// 			JenisKelamin:    row[6],
// 			Agama:           row[7],
// 			AlamatSiswa:     &row[8],
// 			TeleponSiswa:    row[9],
// 			DiterimaTanggal: row[10],
// 			// Umur:   parseInt(row[1]), // Fungsi parseInt bisa digunakan untuk mengubah string ke int
// 			// Alamat: row[2],
// 		}
// 		siswaList = append(siswaList, siswa)
// 	}
// 	return siswaList
// }

// // Fungsi parsing untuk guru
// func parseGuru(rows [][]string) []*models.TabelPTK {
// 	var guruList []*models.TabelPTK
// 	for _, row := range rows[1:] {
// 		if len(row) < 3 {
// 			continue
// 		}
// 		guru := &models.TabelPTK{
// 			// Nama:   row[0],
// 			// Mapel:  row[1],
// 			// Alamat: row[2],
// 		}
// 		guruList = append(guruList, guru)
// 	}
// 	return guruList
// }

// // Fungsi parsing untuk kelas
// func parseKelas(rows [][]string) []*models.RombonganBelajar {
// 	var kelasList []*models.RombonganBelajar
// 	for _, row := range rows[1:] {
// 		if len(row) < 2 {
// 			continue
// 		}
// 		kelas := &models.RombonganBelajar{
// 			// NamaKelas:  row[0],
// 			// JumlahSiswa: parseInt(row[1]),
// 		}
// 		kelasList = append(kelasList, kelas)
// 	}
// 	return kelasList
// }

// // Fungsi parsing untuk kelas
// func parseNilaiAkhir(rows [][]string) []*models.NilaiAkhir {
// var kelasList []*models.NilaiAkhir
// for _, row := range rows[1:] {
// 	if len(row) < 2 {
// 		continue
// 	}
// 	kelas := &models.NilaiAkhir{
// 		// AnggotaRombelID: (uuid.UUID).row[0],
// 		// NamaKelas:  row[0],
// 		// JumlahSiswa: parseInt(row[1]),
// 		// NilaiPeng: int32(parseInt(row[2])),
// 	}
// 	kelasList = append(kelasList, kelas)
// }
// return kelasList
// }

// Fungsi helper untuk mengubah string ke int
// func parseInt(value string) int {
// 	i, _ := strconv.Atoi(value)
// 	return i
// }

// // Fungsi helper untuk mengubah string ke int
// func parseUuid(value *string) uuid.UUID {
// 	i, _ := uuid.Parse(*value)
// 	return i
// }
