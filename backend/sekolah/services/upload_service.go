package services

// import (
// 	"fmt"
// 	"io"
// 	"mime/multipart"
// 	"os"
// 	"path/filepath"
// )

// // UploadService menangani penyimpanan file yang diunggah
// type UploadService struct {
// 	uploadDir string
// }

// // NewUploadService membuat instance UploadService baru
// func NewUploadService(uploadDir string) *UploadService {
// 	// Pastikan folder upload ada
// 	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
// 		err := os.MkdirAll(uploadDir, os.ModePerm)
// 		if err != nil {
// 			panic(fmt.Sprintf("Gagal membuat direktori upload: %v", err))
// 		}
// 	}
// 	return &UploadService{uploadDir: uploadDir}
// }

// // SaveFile menyimpan file yang diunggah ke penyimpanan lokal
// func (s *UploadService) SaveFile(file multipart.File, header *multipart.FileHeader) (string, error) {
// 	defer file.Close()

// 	// Tentukan path penyimpanan file
// 	filePath := filepath.Join(s.uploadDir, header.Filename)

// 	// Buat file baru di lokasi tujuan
// 	dst, err := os.Create(filePath)
// 	if err != nil {
// 		return "", fmt.Errorf("gagal membuat file: %w", err)
// 	}
// 	defer dst.Close()

// 	// Salin isi file ke file tujuan
// 	if _, err := io.Copy(dst, file); err != nil {
// 		return "", fmt.Errorf("gagal menyimpan file: %w", err)
// 	}

// 	return filePath, nil
// }
