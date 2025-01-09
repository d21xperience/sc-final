package main

// import (
// 	"context"
// 	"log"

// 	generated "gateway_service/generated"

// 	"github.com/gin-gonic/gin"
// 	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
// 	"google.golang.org/grpc"
// )

// func main() {
// 	// Koneksi ke backend gRPC server
// 	conn, err := grpc.NewClient("backend/auth_service:50051", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("Failed to connect to gRPC server: %v", err)
// 	}
// 	defer conn.Close()

// 	// Membuat gRPC-Gateway mux
// 	mux := runtime.NewServeMux()
// 	err = generated.RegisterAuthServiceHandler(context.Background(), mux, conn)
// 	if err != nil {
// 		log.Fatalf("Failed to register gRPC Gateway: %v", err)
// 	}

// 	// Inisialisasi Gin
// 	router := gin.Default()

// 	// Menggunakan gRPC-Gateway sebagai handler utama
// 	router.Any("/*any", gin.WrapH(mux))

	// Route tambahan untuk komunikasi langsung dengan backend
	// router.POST("/login", func(c *gin.Context) {
	// 	var req generated.LoginRequest
	// 	if err := c.ShouldBindJSON(&req); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	// Membuat client gRPC untuk AuthService
	// 	client := generated.NewAuthServiceClient(conn)
	// 	resp, err := client.Login(context.Background(), &req)
	// 	if err != nil {
	// 		st, _ := status.FromError(err)
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"token": resp.Token,
	// 		"user":  resp.User,
	// 	})
	// })

	// Menjalankan server pada port 8080
// 	log.Println("Gateway running on port 8080")
// 	router.Run(":8080")
// }
