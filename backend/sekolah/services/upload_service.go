package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	pb "sekolah/generated"
	"strconv"
	"strings"
)

// UploadService menangani penyimpanan file yang diunggah
type UploadServiceServer struct {
	pb.UnimplementedUploadDataSekolahServiceServer
	uploadDir string
}

func NewUploadServiceServer() *UploadServiceServer {
	return &UploadServiceServer{
		uploadDir: "uploads",
	}
}

// UploadFile menangani upload melalui gRPC
func (s *UploadServiceServer) UploadDataSekolah(ctx context.Context, req *pb.UploadDataSekolahRequest) (*pb.UploadDataSekolahResponse, error) {
	// Tentukan lokasi penyimpanan file
	filePath := filepath.Join(s.uploadDir, req.Filename)

	// Simpan file yang dikirim dalam bytes
	err := os.WriteFile(filePath, req.File, 0644)
	if err != nil {
		return nil, fmt.Errorf("gagal menyimpan file: %w", err)
	}

	// Kembalikan URL file yang diunggah
	return &pb.UploadDataSekolahResponse{
		Message: "File berhasil diunggah",
		FileUrl: "/storage/uploads/" + req.Filename,
	}, nil
}

// UploadFileHTTP menangani upload file melalui REST API dengan multipart/form-data
func (s *UploadServiceServer) UploadFileHTTP(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	err := r.ParseMultipartForm(10 << 20) // Batas ukuran file 10MB
	if err != nil {
		http.Error(w, "Gagal mem-parsing form data", http.StatusBadRequest)
		return
	}

	// Ambil parameter dan file
	uploadType := r.FormValue("upload_type")
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
	// filePath := tempFile.Name()
	// Proses file (gunakan service jika diperlukan)
	if uploadType == "siswa" {
		// data, err := uploadData[models.PesertaDidik](filePath, uploadType)
		// if err != nil {
		// 	http.Error(w, fmt.Sprintf("Gagal memproses file: %v", err), http.StatusInternalServerError)
		// 	return
		// }
		// h.service.
	}

	// Berikan respon
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "File berhasil diproses",
		// "data":    data,
	})
}

// // DownloadFile menangani download file melalui gRPC dan REST API
// func (s *UploadServiceServer) DownloadDataSekolah(ctx context.Context, req *pb.DownloadDataSekolahRequest) (*pb.DownloadDataSekolahResponse, error) {
// 	// Tentukan path file yang diminta
// 	filePath := filepath.Join(s.uploadDir, req.Filename)

// 	// Baca file dari storage
// 	fileBytes, err := os.ReadFile(filePath)
// 	if err != nil {
// 		return nil, status.Errorf(codes.NotFound, "File tidak ditemukan: %s", req.Filename)
// 	}

// 	// Kembalikan file dalam bentuk bytes
// 	return &pb.DownloadDataSekolahResponse{
// 		File:     fileBytes,
// 		Filename: req.Filename,
// 	}, nil
// }

// HandleDownloadTemplate adalah handler untuk mengunduh file template .xlsx.
func (h *UploadServiceServer) HandleDownloadTemplate(w http.ResponseWriter, r *http.Request) {
	// Lokasi file template di backend
	templatePath := "./templates/template.xlsx" // Ganti dengan path sebenarnya di backend

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

// GetTemplate menyediakan template Excel berdasarkan jenis data
func (s *UploadServiceServer) GetTemplate(ctx context.Context, req *pb.GetTemplateRequest) (*pb.GetTemplateResponse, error) {
	templateType := req.GetTemplateType()
	templatePath := fmt.Sprintf("/tmp/template_%s.xlsx", templateType)

	// Buat file template jika belum ada
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		err := GenerateTemplate(templateType, templatePath)
		if err != nil {
			return nil, fmt.Errorf("gagal membuat template %s: %w", templateType, err)
		}
	}

	// Baca file template
	data, err := os.ReadFile(templatePath)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca template %s: %w", templateType, err)
	}

	return &pb.GetTemplateResponse{
		FileName: fmt.Sprintf("template_%s.xlsx", templateType),
		FileData: data,
	}, nil
}
