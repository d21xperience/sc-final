package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sekolah/models"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
)

type MultipartHandler struct {
	service             PesertaDidikService
	serviceKelas        RombonganBelajarService
	serviceAnggotaKelas RombelAnggotaService
}

// NewMultipartHandler membuat instance baru MultipartHandler.
func NewMultipartHandler(service PesertaDidikService, serviceKelas RombonganBelajarService, serviceAnggotaKelas RombelAnggotaService) *MultipartHandler {
	return &MultipartHandler{
		service:             service,
		serviceKelas:        serviceKelas,
		serviceAnggotaKelas: serviceAnggotaKelas,
	}
}

// HandleUpload adalah HTTP handler untuk upload file.
func (h *MultipartHandler) HandleBinaryFileUpload(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
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
	ctx := context.Background()
	var dataReponse any
	var messageResponse string
	switch uploadType {
	case "siswa":
		// Proses file (gunakan service jika diperlukan)
		data, err := uploadData[*models.PesertaDidik](tempFile.Name(), uploadType)
		if err != nil {
			http.Error(w, fmt.Sprintf("Gagal memproses file: %v", err), http.StatusInternalServerError)
			return
		}
		err = h.service.SaveMany(ctx, schemaName, data)
		if err != nil {
			http.Error(w, fmt.Sprintf("Gagal menyimpan data: %v", err), http.StatusInternalServerError)
			return
		}
		dataReponse = data
		messageResponse = "Data siswa berhasil diupload"
	case "kelas":
		data, err := uploadData[*models.RombonganBelajar](tempFile.Name(), uploadType)
		if err != nil {
			http.Error(w, fmt.Sprintf("Gagal memproses file: %v", err), http.StatusInternalServerError)
			return
		}
		err = h.serviceKelas.SaveMany(ctx, schemaName, data)
		if err != nil {
			http.Error(w, fmt.Sprintf("Gagal menyimpan data: %v", err), http.StatusInternalServerError)
			return
		}
		dataReponse = data
		messageResponse = "Data kelas berhasil diupload"
	case "anggota_kelas":
		data, err := uploadData[*models.RombelAnggota](tempFile.Name(), uploadType)
		if err != nil {
			http.Error(w, fmt.Sprintf("Gagal memproses file: %v", err), http.StatusInternalServerError)
			return
		}
		err = h.serviceAnggotaKelas.SaveMany(ctx, schemaName, data)
		if err != nil {
			http.Error(w, fmt.Sprintf("Gagal menyimpan data: %v", err), http.StatusInternalServerError)
			return
		}
		dataReponse = data
		messageResponse = "Data anggota kelas berhasil diupload"
	case "mapel":
	case "nilai_akhir":
	}
	// Berikan respon
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": messageResponse,
		"data":    dataReponse,
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
			PesertaDidikID:  row[0],
			Nis:             row[1],
			Nisn:            row[2],
			NmSiswa:         row[3],
			TempatLahir:     row[4],
			TanggalLahir:    row[5],
			JenisKelamin:    row[6],
			Agama:           row[7],
			AlamatSiswa:     parseNullable(row, 8),
			TeleponSiswa:    row[9],
			DiterimaTanggal: row[10],
			NmAyah:          row[11],
			NmIbu:           row[12],
			PekerjaanAyah:   row[13],
			PekerjaanIbu:    row[14],
			NmWali:          parseNullable(row, 15),
			PekerjaanWali:   parseNullable(row, 16),
		}
		siswaList = append(siswaList, siswa)
	}
	return siswaList
}

// Fungsi parsing untuk guru
func parseGuru(rows [][]string) []*models.TabelPTK {
	var guruList []*models.TabelPTK
	for _, row := range rows[1:] {
		if len(row) < 3 {
			continue
		}
		guru := &models.TabelPTK{
			PTKID:             row[1],
			Nama:              row[2],
			NIP:               parseNullable(row, 3),
			JenisPTKID:        row[4],
			JenisKelamin:      row[5],
			TempatLahir:       row[6],
			TanggalLahir:      row[7],
			NUPTK:             parseNullable(row, 8),
			AlamatJalan:       row[9],
			StatusKeaktifanID: row[10],
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
			RombonganBelajarId:  row[1],
			SekolahId:           row[2],
			SemesterId:          row[3],
			JurusanId:           row[4],
			PtkId:               row[5],
			NmKelas:             row[6],
			TingkatPendidikanId: int32(parseInt(row[7])),
			JenisRombel:         int32(parseInt(row[8])),
			NamaJurusanSp:       row[9],
			JurusanSpId:         parseNullable(row, 10),
			KurikulumId:         int32(parseInt(row[11])),
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
			IDNilaiAkhir:    parseUuid(&row[1]),
			AnggotaRombelID: parseUuidPointer(&row[2]),
			MataPelajaranID: parseNullableInt32(row[3]),
			SemesterID:      row[4],
			NilaiPeng:       parseNullableInt32(row[5]),
			PredikatPeng:    row[6],
			NilaiKet:        parseNullableInt32(row[7]),
			PredikatKet:     row[8],
			NilaiSik:        parseNullableInt32(row[8]),
			PredikatSik:     row[10],
			NilaiSikSos:     parseNullableInt32(row[11]),
			PredikatSikSos:  row[12],
			PesertaDidikID:  parseUuidPointer(&row[13]),
			IDMinat:         row[14],
			Semester:        parseNullableInt32(row[15]),
		}
		kelasList = append(kelasList, kelas)
	}
	return kelasList
}

// Helper untuk kolom opsional
func parseNullable(row []string, index int) *string {
	if len(row) > index && strings.TrimSpace(row[index]) != "" {
		value := row[index]
		return &value
	}
	return nil
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

// Fungsi helper untuk mengubah string ke int
func parseInt(value string) int {
	i, _ := strconv.Atoi(value)
	return i
}

func parseNullableInt32(value string) *int32 {
	if value == "" {
		return nil // Nilai kosong, kembalikan nil
	}
	parsed, err := strconv.Atoi(value) // Konversi ke int
	if err != nil {
		return nil // Jika parsing gagal, kembalikan nil
	}
	int32Value := int32(parsed) // Konversi ke int32
	return &int32Value          // Kembalikan pointer
}
func parseUuidPointer(value *string) *uuid.UUID {
	if value == nil || *value == "" {
		return nil // Kembalikan nil jika string kosong atau nil
	}
	parsed, err := uuid.Parse(*value)
	if err != nil {
		return nil // Kembalikan nil jika parsing gagal
	}
	return &parsed // Kembalikan pointer ke uuid.UUID
}

// Fungsi helper untuk mengubah string ke uuid
func parseUuid(value *string) uuid.UUID {
	i, _ := uuid.Parse(*value)
	return i
}

// HandleDownloadTemplate adalah handler untuk mengunduh file template .xlsx.
func (h *MultipartHandler) HandleBinaryFileDownload(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	// Lokasi direktori template
	baseDir := "./templates"

	// Ambil nama file dari query parameter
	filename := r.URL.Query().Get("filename")
	if filename == "" {
		http.Error(w, "Filename is required", http.StatusBadRequest)
		return
	}

	// Tambahkan ekstensi .xlsx jika belum ada
	if !strings.HasSuffix(filename, ".xlsx") {
		filename += ".xlsx"
	}

	// Bersihkan dan gabungkan path
	templatePath := filepath.Join(baseDir, filepath.Clean(filename))

	// Validasi agar file tetap berada di dalam baseDir
	if !strings.HasPrefix(templatePath, filepath.Clean(baseDir)+string(os.PathSeparator)) {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

	// Buka file template
	file, err := os.Open(templatePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Gagal membuka file template: %v", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Mendapatkan informasi file untuk header
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, fmt.Sprintf("Gagal mendapatkan informasi file: %v", err), http.StatusInternalServerError)
		return
	}

	// Set header response untuk file download
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileInfo.Name()))
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

	// Kirim file ke response
	if _, err := io.Copy(w, file); err != nil {
		http.Error(w, fmt.Sprintf("Gagal mengirim file: %v", err), http.StatusInternalServerError)
		return
	}
}
