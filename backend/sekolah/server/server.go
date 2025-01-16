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

// func StartGRPCServer(schemaServices services.SchemaService, sekolahService services.SekolahService) {

// 	// gRPC Listener
// 	listener, err := net.Listen("tcp", ":50052")
// 	if err != nil {
// 		log.Fatalf("Failed to listen: %v", err)
// 	}

// 	// gRPC Server
// 	grpcServer := grpc.NewServer()
// 	pb.RegisterSchoolServiceServer(grpcServer, &SekolahServiceServer{
// 		schemaService:  schemaServices,
// 		sekolahService: sekolahService,
// 		// userProfile:    userProfileService,
// 	})

// 	// HTTP Gateway
// 	mux := runtime.NewServeMux()
// 	ctx := context.Background()
// 	// Buat context dengan timeout 3 detik
// 	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	// defer cancel()
// 	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

// 	err = pb.RegisterSchoolServiceHandlerFromEndpoint(ctx, mux, "localhost:50052", opts)
// 	if err != nil {
// 		log.Fatalf("Failed to register gRPC Gateway: %v", err)
// 	}

// 	// HTTP Listener
// 	httpListener, err := net.Listen("tcp", ":8081")
// 	if err != nil {
// 		log.Fatalf("Failed to listen on HTTP Gateway: %v", err)
// 	}

// 	// Middleware CORS
// 	corsHandler := corsMiddleware(mux)

// 	// Sync WaitGroup
// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	go func() {
// 		defer wg.Done()
// 		fmt.Println("Auth Service running on port 50052 (gRPC)")
// 		if err := grpcServer.Serve(listener); err != nil {
// 			log.Fatalf("Failed to serve gRPC: %v", err)
// 		}
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		fmt.Println("HTTP Gateway running on port 8081")
// 		if err := http.Serve(httpListener, corsHandler); err != nil {
// 			log.Fatalf("Failed to serve HTTP Gateway: %v", err)
// 		}
// 	}()

// 	wg.Wait()
// }

func StartGRPCServer(schemaServices services.SchemaService, sekolahService services.SekolahService, pesertaDidikService services.PesertaDidikService) {
	// Buat context utama dengan cancel untuk shutdown yang aman
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Pastikan resource dibersihkan saat fungsi keluar

	// Menangani signal dari OS (Ctrl+C, SIGTERM)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// gRPC Listener
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// gRPC Server
	grpcServer := grpc.NewServer()
	pb.RegisterSchoolServiceServer(grpcServer, &SekolahServiceServer{
		schemaService:       schemaServices,
		sekolahService:      sekolahService,
		pesertaDidikService: pesertaDidikService,
	})

	// HTTP Gateway
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err = pb.RegisterSchoolServiceHandlerFromEndpoint(ctx, mux, "localhost:50052", opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Gateway: %v", err)
	}

	// HTTP Listener
	httpListener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Failed to listen on HTTP Gateway: %v", err)
	}

	// Middleware CORS
	corsHandler := corsMiddleware(mux)

	// Sync WaitGroup
	var wg sync.WaitGroup
	wg.Add(2)

	// Menjalankan gRPC server dalam Goroutine
	go func() {
		defer wg.Done()
		fmt.Println("gRPC Service running on port 50052")
		if err := grpcServer.Serve(listener); err != nil {
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
		grpcServer.GracefulStop()

		// Matikan HTTP Gateway
		httpListener.Close()

		// Batalkan context utama agar semua operasi berhenti
		cancel()
	}()

	// Menunggu semua Goroutine selesai
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
