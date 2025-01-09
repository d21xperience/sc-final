package server

import (
	"fmt"
	"log"
	"net"

	pb "auth_service/generated"
	"auth_service/services"

	"github.com/go-redis/redis"
	"google.golang.org/grpc"
)

// Inisialisasi Redis client global
func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "redis_container:6379",
	})
}
func StartGRPCServer(authService services.AuthService, sekolahService services.SekolahService, userProfileService services.UserProfileService) {

	// Inisialisasi Redis sekali saja
	redisClient := InitRedis()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	// Inisialisasi gRPC server
	grpcServer := grpc.NewServer()
	// Register service
	pb.RegisterAuthServiceServer(grpcServer, &AuthServiceServer{
		authService:    authService,
		sekolahService: sekolahService,
		userProfile:    userProfileService,
		RedisClient:    redisClient,
	})

	fmt.Println("Auth Service running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
