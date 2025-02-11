package handlers

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"sekolah/registry"
// 	"sekolah/services"
// )

// // MultipartHandler menangani upload file dan berinteraksi dengan registry
// type MultipartHandler struct {
// 	serviceRegistry *registry.ServiceRegistry
// }

// // NewMultipartHandler membuat instance baru MultipartHandler
// func NewMultipartHandler(serviceRegistry *registry.ServiceRegistry) *MultipartHandler {
// 	return &MultipartHandler{serviceRegistry: serviceRegistry}
// }

// // UploadHandler menangani request upload file
// func (h *MultipartHandler) UploadHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Metode harus POST", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// Ambil file dari request
// 	file, header, err := r.FormFile("file")
// 	if err != nil {
// 		http.Error(w, "Gagal membaca file", http.StatusBadRequest)
// 		return
// 	}
// 	defer file.Close()

// 	// Ambil UploadService dari registry
// 	uploadService, exists := h.serviceRegistry.GetService("UploadService")
// 	if !exists {
// 		http.Error(w, "UploadService tidak ditemukan", http.StatusInternalServerError)
// 		return
// 	}

// 	// Konversi service ke tipe UploadService
// 	uploader, ok := uploadService.(*services.UploadService)
// 	if !ok {
// 		http.Error(w, "Gagal mendapatkan UploadService", http.StatusInternalServerError)
// 		return
// 	}

// 	// Simpan file menggunakan UploadService
// 	filePath, err := uploader.SaveFile(file, header)
// 	if err != nil {
// 		http.Error(w, "Gagal menyimpan file", http.StatusInternalServerError)
// 		return
// 	}

// 	// Berikan respons sukses
// 	log.Printf("File berhasil diupload: %s", filePath)
// 	fmt.Fprintf(w, "File berhasil diupload: %s", filePath)
// }
