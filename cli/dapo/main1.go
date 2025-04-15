/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
// package main

// func main() {
// 	// cmd.Execute()

// }
package main

import (
	"dapo/cmd"
	"log"

	"github.com/joho/godotenv"
)

// Fungsi utama untuk mengirim request dengan Bearer token
func main1() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default env.")
	}
	cmd.Execute()
}
