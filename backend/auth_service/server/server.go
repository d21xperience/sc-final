package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func StartGRPCServer() {
	// Menggunakan environment variable untuk fleksibilitas
	grpcHost := os.Getenv("GRPC_HOST")
	if grpcHost == "" {
		grpcHost = "localhost"
	}

	gRPCPort := os.Getenv("GRPC_PORT")
	if gRPCPort == "" {
		gRPCPort = "50051"
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8081"
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
	mux := runtime.NewServeMux()
	// Middleware CORS
	corsHandler := corsMiddleware(mux)
	grpcServerEndpoint := fmt.Sprintf("%s:%s", grpcHost, gRPCPort)
	RunHTTPGateway(ctx, mux, grpcServerEndpoint, httpPort) // HTTP gateway di port 8080
	// HTTP Listener
	httpListener, err := net.Listen("tcp", ":"+httpPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Sync WaitGroup
	var wg sync.WaitGroup
	// wg.Add(2)
	wg.Add(2)

	// Menjalankan gRPC server dalam Goroutine
	go func() {
		defer wg.Done()
		log.Printf("gRPC server berjalan di :%s", gRPCPort)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Menjalankan HTTP Gateway dalam Goroutine
	go func() {
		defer wg.Done()
		log.Printf("HTTP gateway berjalan di :%s", httpPort)
		if err := http.Serve(httpListener, corsHandler); err != nil {
			log.Fatalf("Failed to serve HTTP Gateway: %v", err)
		}
	}()

	// Menunggu sinyal shutdown
	wg.Add(1) // Tambahkan WaitGroup untuk shutdown goroutine
	go func() {
		defer wg.Done() // Pastikan WaitGroup diberi tahu setelah selesai
		<-signalChan
		fmt.Println("\nShutting down servers...")

		// Timeout shutdown dalam 5 detik
		_, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel()

		// Matikan server gRPC
		grpcServer.GracefulStop()

		// Matikan HTTP Gateway
		if err := httpListener.Close(); err != nil {
			log.Printf("Error while closing HTTP listener: %v", err)
		}

		// Batalkan context utama agar semua operasi berhenti
		cancel()
	}()

	// Menunggu semua Goroutine selesai
	wg.Wait()
	fmt.Println("Server shutdown complete")
}

func corsMiddleware(h http.Handler) http.Handler {
	frontend := os.Getenv("FRONTEND")
	if frontend == "" {
		frontend = "http://localhost:5173"
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", frontend)
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
