package utils

import (
	crand "crypto/rand" // Alias untuk crypto/rand
	"fmt"
	"math/big"
	mrand "math/rand" // Alias untuk math/rand
	"strings"
	"time"
)

// GenerateUsername generates a unique username based on name and ID.
// func GenerateUsername(name string, id int, isUsernameTaken func(string) bool) string {
// 	// Normalize the name: lowercase, remove special characters, and trim
// 	name = strings.ToLower(name)
// 	name = regexp.MustCompile(`[^a-z0-9]+`).ReplaceAllString(name, "")
// 	name = strings.TrimSpace(name)

// 	// Fallback if the name becomes empty after normalization
// 	if len(name) == 0 {
// 		name = "user"
// 	}

// 	// Base username
// 	username := fmt.Sprintf("%s%d", name, id)

// 	// Ensure uniqueness
// 	counter := 0
// 	for isUsernameTaken(username) {
// 		counter++
// 		username = fmt.Sprintf("%s%d", name, id+counter)
// 	}

// 	return username
// }

// GenerateUsername membuat username dengan format {inisial_sekolah}_{nama_depan}{4_digit_acak}
func GenerateUsername(namaSekolah, namaSiswa string) string {
	// Ambil inisial sekolah
	inisialSekolah := strings.ToLower(strings.ReplaceAll(namaSekolah, " ", ""))
	// Ambil nama depan siswa
	namaDepan := strings.Split(strings.ToLower(namaSiswa), " ")[0]

	// Menggunakan local random generator
	source := mrand.NewSource(time.Now().UnixNano())
	randomGen := mrand.New(source)
	randomNumber := randomGen.Intn(9000) + 1000 // Angka antara 1000-9999

	return fmt.Sprintf("%s_%s%d", inisialSekolah, namaDepan, randomNumber)
}

// GeneratePassword membuat password acak dengan panjang tertentu
func GeneratePassword(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+"
	password := make([]byte, length)
	for i := range password {
		randomChar, err := crand.Int(crand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", fmt.Errorf("failed to generate random character: %w", err)
		}
		password[i] = charset[randomChar.Int64()]
	}
	return string(password), nil
}
