package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sc-service/services"
	"sync"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func StartServer() {
	grpcHost := os.Getenv("GRPC_HOST")
	if grpcHost == "" {
		grpcHost = "localhost"
	}

	gRPCPort := os.Getenv("GRPC_PORT")
	if gRPCPort == "" {
		gRPCPort = "50053"
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8083"
	}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// Menangani signal dari OS (Ctrl+C, SIGTERM)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	// gRPC Gateway
	// =========================================
	// Jalankan server gRPC dan gateway
	// gRPC Listener
	listener, err := net.Listen("tcp", ":"+gRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := RunGRPCServer() // gRPC di port 50052
	// HTTP Gateway
	// =========================================
	// Inisialisasi mux untuk HTTP Gateway
	// Inisialisasi UploadServiceServer sebelum dipakai
	UploadService := services.NewUploadServiceServer()
	mux := runtime.NewServeMux()
	method, pattern := createPattern("POST", "api", "v1", "ss", "upload", "rest")
	mux.Handle(method, pattern, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		UploadService.UploadFileHTTP(w, r)
	})
	method, pattern = createPattern("GET", "api", "v1", "ss", "download", "template")
	mux.Handle(method, pattern, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		UploadService.DownloadTemplateHTTP(w, r)
	})
	// Middleware CORS
	corsHandler := corsMiddleware(mux)
	// HTTP Server dengan Timeout
	httpServer := &http.Server{
		Addr:         ":" + httpPort,
		Handler:      corsHandler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
	grpcServerEndpoint := fmt.Sprintf("%s:%s", grpcHost, gRPCPort)
	// ================================================
	RunHTTPGateway(ctx, mux, grpcServerEndpoint, httpPort) // HTTP gateway di port 8080
	var wg sync.WaitGroup

	// Tambahkan jumlah goroutine sebelum menjalankan
	wg.Add(2) // Karena kita menjalankan dua goroutine

	// Jalankan gRPC Server dalam Goroutine
	go func() {
		defer wg.Done()
		log.Printf("gRPC server berjalan di :%s", gRPCPort)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Jalankan HTTP Gateway dalam Goroutine
	go func() {
		defer wg.Done()
		log.Printf("HTTP gateway berjalan di :%s", httpPort)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to serve HTTP Gateway: %v", err)
		}
	}()

	// Menangani shutdown dengan signal handling
	<-signalChan
	fmt.Println("\nShutting down servers...")

	// Matikan server gRPC
	grpcServer.GracefulStop()

	// Matikan HTTP Gateway
	httpServer.Close()

	// Batalkan context utama
	cancel()

	// Tunggu semua goroutine selesai
	wg.Wait()
	fmt.Println("Server shutdown complete")
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

func createPattern(method string, pathSegments ...string) (string, runtime.Pattern) {
	pattern := runtime.MustPattern(
		runtime.NewPattern(1, generatePatternIndexes(len(pathSegments)), pathSegments, ""),
	)
	return method, pattern
}

// generatePatternIndexes membantu membuat pola angka yang sesuai dengan jumlah segment
func generatePatternIndexes(segmentCount int) []int {
	indexes := []int{}
	for i := 0; i < segmentCount; i++ {
		indexes = append(indexes, 2, i)
	}
	return indexes
}
