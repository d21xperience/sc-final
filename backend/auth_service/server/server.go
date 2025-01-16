package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	pb "auth_service/generated"
	"auth_service/services"

	"github.com/go-redis/redis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// Inisialisasi Redis client global
func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "redis_container:6379",
	})
}

func StartGRPCServer(authService services.AuthService, sekolahService services.SekolahService, userProfileService services.UserProfileService) {
	// gRPC Listener
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// gRPC Server
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, &AuthServiceServer{
		authService:    authService,
		sekolahService: sekolahService,
		userProfile:    userProfileService,
	})

	// HTTP Gateway
	mux := runtime.NewServeMux()
	ctx := context.Background()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err = pb.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Gateway: %v", err)
	}

	// HTTP Listener
	httpListener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen on HTTP Gateway: %v", err)
	}

	// Middleware CORS
	corsHandler := corsMiddleware(mux)

	// Sync WaitGroup
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Auth Service running on port 50051 (gRPC)")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	go func() {
		defer wg.Done()
		fmt.Println("HTTP Gateway running on port 8080")
		if err := http.Serve(httpListener, corsHandler); err != nil {
			log.Fatalf("Failed to serve HTTP Gateway: %v", err)
		}
	}()

	wg.Wait()
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
