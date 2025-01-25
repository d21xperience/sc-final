package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sekolah/services"
	"sync"
	"syscall"
	"time"

	pb "sekolah/generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type GRPCServer struct {
	grpcServer              *grpc.Server
	schemaService           services.SchemaService
	sekolahService          services.SekolahService
	tahunAjaranService      services.TahunAjaranService
	semesterService         services.SemesterService
	pesertaDidikService     services.PesertaDidikService
	rombonganBelajarService services.RombonganBelajarService
	rombelAnggotaService    services.RombelAnggotaService
	nilaiAkhirService       services.NilaiAkhirService
}

// Jalankan gRPC Server
func (s *GRPCServer) run() {
	// Buat context utama dengan cancel untuk shutdown yang aman
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Pastikan resource dibersihkan saat fungsi keluar

	// Menangani signal dari OS (Ctrl+C, SIGTERM)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// gRPC server endpoint
	grpcServerEndpoint := "localhost:50052"
	// gRPC Listener
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	// HTTP Gateway
	// =========================================
	multiparthandle := services.NewMultipartHandler(s.pesertaDidikService, s.rombonganBelajarService, s.rombelAnggotaService)

	// =========================================

	mux := runtime.NewServeMux()
	mux.HandlePath("POST", "/api/v1/ss/upload", multiparthandle.HandleBinaryFileUpload)
	mux.HandlePath("GET", "/api/v1/ss/download", multiparthandle.HandleBinaryFileDownload)

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err = pb.RegisterSekolahServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Gateway: %v", err)
	}

	err = pb.RegisterTahunAjaranServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Tahun Ajaran Gateway: %v", err)
	}

	err = pb.RegisterSemesterServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Semester Gateway: %v", err)
	}
	err = pb.RegisterSiswaServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Siswa Gateway: %v", err)
	}
	err = pb.RegisterKelasServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Kelas Gateway: %v", err)
	}
	err = pb.RegisterAnggotaKelasServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Kelas Gateway: %v", err)
	}
	err = pb.RegisterNilaiAkhirServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Nilai akhir Gateway: %v", err)
	}
	err = pb.RegisterUploadDataSekolahServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Upload data Sekolah Gateway: %v", err)
	}

	// HTTP Listener
	httpListener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Failed to listen on HTTP Gateway: %v", err)
	}

	// Middleware CORS
	corsHandler := corsMiddleware(mux)
	// Mengurutkan middleware sehingga mereka bekerja secara berantai.
	// multipartHandler := multipartMiddleware(corsHandler)
	// loggingHandler := logMiddleware(multipartHandler)

	// Sync WaitGroup
	var wg sync.WaitGroup
	wg.Add(2)

	// Menjalankan gRPC server dalam Goroutine
	go func() {
		defer wg.Done()
		fmt.Println("gRPC Service running on port 50052")
		if err := s.grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Menjalankan HTTP Gateway dalam Goroutine
	go func() {
		defer wg.Done()
		fmt.Println("HTTP Gateway running on port 8081")
		if err := http.Serve(httpListener, corsHandler); err != nil {
			log.Fatalf("Failed to serve HTTP Gateway: %v", err)
		}
	}()

	// Menunggu sinyal shutdown
	go func() {
		<-signalChan
		fmt.Println("\nShutting down servers...")

		// Timeout shutdown dalam 5 detik
		_, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel() // Pastikan timeout berlaku

		// Matikan server gRPC
		s.grpcServer.GracefulStop()

		// Matikan HTTP Gateway
		httpListener.Close()

		// Batalkan context utama agar semua operasi berhenti
		cancel()
	}()

	// Menunggu semua Goroutine selesai
	wg.Wait()
	fmt.Println("Server shutdown complete")
}

func StartGRPCServer(schemaServices services.SchemaService, sekolahService services.SekolahService, tahunAjaranService services.TahunAjaranService, semesterService services.SemesterService, pesertaDidikService services.PesertaDidikService, rombonganBelajarService services.RombonganBelajarService, rombelAnggotaServices services.RombelAnggotaService, nilaiAkhirService services.NilaiAkhirService) {
	// Buat instance server
	server := &GRPCServer{
		grpcServer:              grpc.NewServer(),
		schemaService:           schemaServices,
		sekolahService:          sekolahService,
		pesertaDidikService:     pesertaDidikService,
		tahunAjaranService:      tahunAjaranService,
		semesterService:         semesterService,
		rombonganBelajarService: rombonganBelajarService,
		rombelAnggotaService:    rombelAnggotaServices,
		nilaiAkhirService:       nilaiAkhirService,
	}
	// gRPC Server
	// grpcServer := grpc.NewServer()
	pb.RegisterSekolahServiceServer(server.grpcServer, &SekolahServiceServer{
		schemaService:  server.schemaService,
		sekolahService: server.sekolahService,
	})
	pb.RegisterTahunAjaranServiceServer(server.grpcServer, &TahunAjaranServiceServer{
		TahunAjaranService: server.tahunAjaranService,
	})
	pb.RegisterSemesterServiceServer(server.grpcServer, &SemesterServiceServer{
		SemesterService: server.semesterService,
	})
	// REGISTER SISWA
	pb.RegisterSiswaServiceServer(server.grpcServer, &SiswaServiceServer{
		pesertaDidikService: server.pesertaDidikService,
	})
	// REGISTER KELAS
	pb.RegisterKelasServiceServer(server.grpcServer, &RombelServiceServer{
		rombonganBelajarService: server.rombonganBelajarService,
	})
	// REGISTER ANGGOTA KELAS
	pb.RegisterAnggotaKelasServiceServer(server.grpcServer, &RombelAnggotaServiceServer{
		rombelAnggotaService: server.rombelAnggotaService,
	})
	// REGISTER ANGGOTA KELAS
	pb.RegisterNilaiAkhirServiceServer(server.grpcServer, &NilaiAkhirServiceServer{
		NilaiAkhirService: server.nilaiAkhirService,
	})
	// // REGISTER UPLOAD SERVICE
	// pb.RegisterUploadDataSekolahServiceServer(server.grpcServer, &UploadDataSekolahServiceServer{
	// 	pd: server.pesertaDidikService,
	// })

	// Jalankan server
	server.run()
}

func corsMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}

// func handleBinaryFileUpload(w http.ResponseWriter, r *http.Request, params map[string]string) {
// 	// Parse form data
// 	err := r.ParseMultipartForm(10 << 20) // Limit file size to 10MB
// 	if err != nil {
// 		http.Error(w, "Unable to parse form", http.StatusBadRequest)
// 		return
// 	}

// 	// Extract parameters
// 	// uploadType := r.FormValue("upload_type")
// 	// schemaName := r.FormValue("schemaname")
// 	fileHeader := r.MultipartForm.File["file"]
// 	if len(fileHeader) == 0 {
// 		http.Error(w, "File is required", http.StatusBadRequest)
// 		return
// 	}
// 	file, err := fileHeader[0].Open()
// 	if err != nil {
// 		http.Error(w, "Unable to open file", http.StatusInternalServerError)
// 		return
// 	}
// 	defer file.Close()

// 	fileName := fileHeader[0].Filename
// 	if !strings.HasSuffix(fileName, ".xlsx") {
// 		http.Error(w, "Only Excel files (.xlsx) are allowed", http.StatusBadRequest)
// 		return
// 	}

// 	// Save the uploaded file to a temporary location
// 	tempFile, err := os.CreateTemp("", "upload-*.xlsx")
// 	if err != nil {
// 		http.Error(w, "Unable to create temp file", http.StatusInternalServerError)
// 		return
// 	}
// 	defer tempFile.Close()

// 	// Write uploaded file content to temporary file
// 	_, err = io.Copy(tempFile, file)
// 	if err != nil {
// 		http.Error(w, "Unable to save file", http.StatusInternalServerError)
// 		return
// 	}

// 	// Process the Excel file
// 	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	// defer cancel() // Pastikan untuk membatalkan konteks setelah selesai
// 	// tempFilePath := tempFile.Name()
// 	// if uploadType == "siswa" {
// 	// 	excelData, err := services.UploadData[*models.PesertaDidik](tempFilePath, uploadType, schemaName)
// 	// 	if err != nil {
// 	// 		http.Error(w, fmt.Sprintf("Failed to process Excel file: %v", err), http.StatusInternalServerError)
// 	// 		return
// 	// 	}
// 	// 	// services.
// 	// 	// Return processed data as JSON response
// 	// 	w.Header().Set("Content-Type", "application/json")
// 	// 	json.NewEncoder(w).Encode(map[string]interface{}{
// 	// 		"message":   "File processed successfully",
// 	// 		"file_name": excelData,
// 	// 		// "data":      excelData,
// 	// 	})
// 	// }
// 	// Return processed data as JSON response
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(map[string]interface{}{
// 		"message":   "File processed successfully",
// 		"file_name": fileName,
// 		// "data":      excelData,
// 	})
// }
