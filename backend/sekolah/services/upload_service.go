package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sekolah/models"
	"strings"

	"github.com/xuri/excelize/v2"
)

type UploadHandler struct {
	service PesertaDidikService
}

// NewUploadHandler membuat instance baru UploadHandler.
func NewUploadHandler(service PesertaDidikService) *UploadHandler {
	return &UploadHandler{service: service}
}

// HandleUpload adalah HTTP handler untuk upload file.
func (h *UploadHandler) HandleBinaryFileUpload(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	// Parse form data
	err := r.ParseMultipartForm(10 << 20) // Batas ukuran file 10MB
	if err != nil {
		http.Error(w, "Gagal mem-parsing form data: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Ambil parameter dan file
	uploadType := r.FormValue("upload_type")
	schemaName := r.FormValue("schema_name")
	fileHeader := r.MultipartForm.File["file"]
	if len(fileHeader) == 0 {
		http.Error(w, "File tidak ditemukan dalam request", http.StatusBadRequest)
		return
	}
	file, err := fileHeader[0].Open()
	if err != nil {
		http.Error(w, "Gagal membuka file: "+err.Error(), http.StatusInternalServerError)
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
		http.Error(w, "Gagal membuat file sementara: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	_, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(w, "Gagal menyimpan file sementara: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Proses file (gunakan service jika diperlukan)
	data, err := uploadData[*models.PesertaDidik](tempFile.Name(), uploadType)
	if err != nil {
		http.Error(w, fmt.Sprintf("Gagal memproses file: %v", err), http.StatusInternalServerError)
		return
	}

	ctx := context.Background()
	// Anda harus implementasikan metode pada service
	// ptrData := toPointerSlice(data) // Konversi ke []*models.PesertaDidik

	err = h.service.SaveMany(ctx, schemaName, data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Gagal menyimpan data: %v", err), http.StatusInternalServerError)
		return
	}
	// Berikan respon
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "File berhasil diproses",
		"data":    data,
	})
}

func uploadData[T any](filePath, uploadType string) ([]T, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca file Excel: %w", err)
	}
	defer f.Close()

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return nil, fmt.Errorf("file Excel tidak memiliki sheet")
	}

	rows, err := f.GetRows(sheets[0]) // Ambil data dari sheet pertama
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data dari sheet: %w", err)
	}

	if len(rows) < 2 { // Pastikan setidaknya ada header + 1 baris data
		return nil, fmt.Errorf("file Excel kosong atau tidak memiliki data yang valid")
	}

	var result []T
	switch uploadType {
	case "siswa":
		// return parseSiswa(rows), nil
		result = castToGeneric[T](parseSiswa(rows))
	case "guru":
		result = castToGeneric[T](parseGuru(rows))
	case "kelas":
		result = castToGeneric[T](parseKelas(rows))
	case "nilaiAkhir":
		result = castToGeneric[T](parseNilaiAkhir(rows))
	default:
		return nil, fmt.Errorf("jenis unggahan tidak dikenali: %s", uploadType)
	}

	return result, nil
}

// Helper untuk casting generik
func castToGeneric[T any](input interface{}) []T {
	if v, ok := input.([]T); ok {
		return v
	}
	return nil
}

// Fungsi parsing untuk siswa
func parseSiswa(rows [][]string) []*models.PesertaDidik {
	var siswaList []*models.PesertaDidik
	for i, row := range rows[1:] {
		if len(row) < 7 {
			fmt.Printf("Baris %d memiliki data yang tidak lengkap\n", i+2)
			continue
		}
		// // Konversi tanggal lahir
		// tanggalLahirExcel, err := strconv.Atoi(row[5]) // Baris 5 adalah tanggal lahir
		// if err != nil {
		// 	fmt.Printf("Baris %d memiliki tanggal lahir tidak valid: %v\n", i+2, err)
		// 	continue
		// }
		// // Konversi tanggal diterima
		// tanggalDiterimaExcel, err := strconv.Atoi(row[10]) // Baris 5 adalah tanggal lahir
		// if err != nil {
		// 	fmt.Printf("Baris %d memiliki tanggal lahir tidak valid: %v\n", i+2, err)
		// 	continue
		// }

		// tanggalLahir, err := parseExcelDate(tanggalLahirExcel)
		// if err != nil {
		// 	fmt.Printf("Baris %d gagal mengonversi tanggal lahir: %v\n", i+2, err)
		// 	continue
		// }
		// tanggaDiterima, err := parseExcelDate(tanggalDiterimaExcel)
		// if err != nil {
		// 	fmt.Printf("Baris %d gagal mengonversi tanggal lahir: %v\n", i+2, err)
		// 	continue
		// }
		// FORMAT TANGGAL YYYY-MM-DD
		siswa := &models.PesertaDidik{
			PesertaDidikID: row[0],
			Nis:            row[1],
			Nisn:           row[2],
			NmSiswa:        row[3],
			TempatLahir:    row[4],
			TanggalLahir:   row[5],
			JenisKelamin:   row[6],
			Agama:          row[7],
			AlamatSiswa:    parseNullable(row, 8),
			// TeleponSiswa:    parseNullable(row, 9),
			DiterimaTanggal: row[10],
		}
		siswaList = append(siswaList, siswa)
	}
	return siswaList
}

// Helper untuk kolom opsional
func parseNullable(row []string, index int) *string {
	if len(row) > index && strings.TrimSpace(row[index]) != "" {
		value := row[index]
		return &value
	}
	return nil
}

// Fungsi parsing untuk guru
func parseGuru(rows [][]string) []*models.TabelPTK {
	var guruList []*models.TabelPTK
	for _, row := range rows[1:] {
		if len(row) < 3 {
			continue
		}
		guru := &models.TabelPTK{
			// Nama:   row[0],
			// Mapel:  row[1],
			// Alamat: row[2],
		}
		guruList = append(guruList, guru)
	}
	return guruList
}

// Fungsi parsing untuk kelas
func parseKelas(rows [][]string) []*models.RombonganBelajar {
	var kelasList []*models.RombonganBelajar
	for _, row := range rows[1:] {
		if len(row) < 2 {
			continue
		}
		kelas := &models.RombonganBelajar{
			// NamaKelas:  row[0],
			// JumlahSiswa: parseInt(row[1]),
		}
		kelasList = append(kelasList, kelas)
	}
	return kelasList
}

// Fungsi parsing untuk kelas
func parseNilaiAkhir(rows [][]string) []*models.NilaiAkhir {
	var kelasList []*models.NilaiAkhir
	for _, row := range rows[1:] {
		if len(row) < 2 {
			continue
		}
		kelas := &models.NilaiAkhir{
			// AnggotaRombelID: (uuid.UUID).row[0],
			// NamaKelas:  row[0],
			// JumlahSiswa: parseInt(row[1]),
			// NilaiPeng: int32(parseInt(row[2])),
		}
		kelasList = append(kelasList, kelas)
	}
	return kelasList
}

// // Konversi dari angka Excel ke waktu Go
// func parseExcelDate(serial int) (string, error) {
// 	// Referensi tanggal Excel (1 Januari 1900) minus offset 2 hari
// 	referenceDate := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)
// 	// Tambahkan jumlah hari ke referensi
// 	parsedDate := referenceDate.AddDate(0, 0, serial)
// 	return parsedDate.Format("2006-01-02"), nil
// }
// func toPointerSlice(data []models.PesertaDidik) []*models.PesertaDidik {
// 	ptrSlice := make([]*models.PesertaDidik, len(data))
// 	for i := range data {
// 		ptrSlice[i] = &data[i]
// 	}
// 	return ptrSlice
// }

// // Fungsi helper untuk mengubah string ke int
// func parseInt(value string) int {
// 	i, _ := strconv.Atoi(value)
// 	return i
// }

// // Fungsi helper untuk mengubah string ke int
// func parseUuid(value *string) uuid.UUID {
// 	i, _ := uuid.Parse(*value)
// 	return i
// }
// Helper untuk kolom opsional
// func parseNullable(row []string, index int) *string {
// 	if len(row) > index && strings.TrimSpace(row[index]) != "" {
// 		value := row[index]
// 		return &value
// 	}
// 	return nil
// }
