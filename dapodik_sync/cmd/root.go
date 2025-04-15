package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd adalah command utama untuk aplikasi CLI
var rootCmd = &cobra.Command{
	Use:   "dapodik_sync_cli",
	Short: "CLI untuk melakukan sinkronisasi data dapodik",
	Run: func(cmd *cobra.Command, args []string) {
		// Menampilkan help jika tidak ada sub-command yang dipanggil
		fmt.Println("Pilih perintah yang ingin dijalankan")
	},
}

// Execute menjalankan root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
