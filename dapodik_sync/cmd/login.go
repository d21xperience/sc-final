package cmd

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	authpb "dapodik_sync/generated"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login ke sistem",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Username: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)

		fmt.Print("Password: ")
		passwordBytes, _ := term.ReadPassword(int(os.Stdin.Fd()))
		fmt.Println()
		password := strings.TrimSpace(string(passwordBytes))

		conn, err := grpc.NewClient(viper.GetString("AUTH_SERVICE_ADDRESS"), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("gagal terhubung ke auth service: %v", err)
		}
		defer conn.Close()

		client := authpb.NewAuthServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		res, err := client.Login(ctx, &authpb.LoginRequest{
			Username: username,
			Password: password,
		})

		if err != nil {
			log.Fatalf("Login gagal: %v", err)
		}

		fmt.Println("Login berhasil. Token:", res.Token)
	},
}
