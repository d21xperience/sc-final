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
	// return redis.NewClient(&redis.Options{
	// 	Addr: "redis_container:6378",
	// })
	redisAddr := "redis_container:6378"
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr, //redisAddr,
		// Password: redisPassword,
		// DB:       0, // Menggunakan database default (DB 0)
	})

	// Cek apakah Redis bisa dihubungi
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis at", redisAddr)
	return rdb

}

func StartGRPCServer(authService services.AuthService, sekolahService services.SekolahService, userProfileService services.UserProfileService) {

	// // Inisialisasi Redis
	// rdb := InitRedis()
	// defer rdb.Close()
	// // Subscribe ke channel Redis
	// sub := subscribeToChannel(rdb, "school_registration")
	// defer sub.Close()

	// fmt.Println("Listening for messages on channel 'school_registration'...")

	// // Menerima pesan dari channel
	// ch := sub.Channel()
	// for msg := range ch {
	// 	handleMessage(msg)
	// }

	//---------------------- gRCP server
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
		RedisClient:    rdb,
	})

	fmt.Println("Auth Service running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
