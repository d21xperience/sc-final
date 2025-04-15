package services

import (
	"dapo/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Fungsi untuk mengirim permintaan GET dan mengembalikan respons JSON
func FetchJSON(url string) ([]byte, error) {
	// Kirim permintaan GET
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error sending GET request: %w", err)
	}
	defer resp.Body.Close()

	// Pastikan status response 200 OK
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 response code: %d", resp.StatusCode)
	}

	// Memeriksa apakah respons berupa JSON
	if resp.Header.Get("Content-Type") != "application/json" {
		return nil, fmt.Errorf("expected JSON response, got: %s", resp.Header.Get("Content-Type"))
	}

	// Membaca response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return body, nil
}

// Fungsi untuk memparsing respons JSON ke dalam struct PesertaDidik
// func parsePesertaDidik(body []byte) ([]PesertaDidik, error) {
// 	var pesertaDidik []PesertaDidik

// 	// Parsing JSON ke struct PesertaDidik
// 	err := json.Unmarshal(body, &pesertaDidik)
// 	if err != nil {
// 		return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
// 	}

// 	return pesertaDidik, nil
// }

// FileExists checks if a file exists
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

// SaveToFile saves data to a JSON file
func SaveToFile(data []byte, filePath string) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	// Write data to file
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}

// ReadStudentData reads student data from JSON file
func ReadStudentData(filePath string) (*models.PesertaDidik, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var studentData models.PesertaDidik
	if err := json.Unmarshal(data, &studentData); err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return &studentData, nil
}

// ReadJSONFile reads JSON file into the given target struct
func ReadJSONFile(filePath string) (interface{}, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	defer file.Close()

	// Baca isi file
	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Gagal membaca file: %v", err)
	}

	// Unmarshal JSON
	var response models.RegistrasiResponse
	err = json.Unmarshal(byteValue, &response)
	if err != nil {
		log.Fatalf("Gagal parse JSON: %v", err)
	}
	// for _, registrasi := range response.Rows {
	// 	fmt.Printf("Nama: %s, NISN: %s, Asal Sekolah: %s\n", registrasi.Nama, registrasi.Nisn, registrasi.SekolahAsal)
	// }
	return &response.Rows, nil
}
func ReadJSONToStruct[T any](filePath string, out *T) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("gagal membuka file: %w", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("gagal membaca file: %w", err)
	}

	err = json.Unmarshal(byteValue, out)
	if err != nil {
		return fmt.Errorf("gagal parsing JSON: %w", err)
	}

	return nil
}
