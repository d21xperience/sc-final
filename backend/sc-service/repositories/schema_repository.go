package repositories

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

type SchemaRepository interface {
	InitializeDatabase(ctx context.Context, schemaFile, schemaName string) error
	SetSchema(schemaName string) error
}

type schemaRepositoryImpl struct {
	db *gorm.DB
}

func NewSchemaRepository(db *gorm.DB) SchemaRepository {
	return &schemaRepositoryImpl{db: db}
}

// ExecuteSQL menjalankan perintah SQL dari string
func (r *schemaRepositoryImpl) ExecuteSQL(query string) error {
	if err := r.db.Exec(query).Error; err != nil {
		return err
	}
	return nil
}

// LoadSQLFile membaca isi file SQL lalu mengganti placeholder {{schema_name}}
// File dibaca baris demi baris untuk mengurangi penggunaan memori.
func (r *schemaRepositoryImpl) LoadSQLFile(filePath, schemaName string) (string, error) {
	// Validasi parameter
	if filePath == "" {
		return "", fmt.Errorf("file path is empty")
	}
	if schemaName == "" {
		return "", fmt.Errorf("schema name is empty")
	}

	// Validasi schemaName untuk karakter yang aman
	isValidSchemaName := regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString
	if !isValidSchemaName(schemaName) {
		return "", fmt.Errorf("invalid schema name: %s", schemaName)
	}

	// Buka file untuk dibaca
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Buffer untuk menampung konten SQL
	var builder strings.Builder
	scanner := bufio.NewScanner(file)

	// Iterasi setiap baris dalam file
	for scanner.Scan() {
		line := scanner.Text()

		// Ganti placeholder {{schema_name}} dengan schemaName
		line = strings.ReplaceAll(line, "{{schema_name}}", schemaName)

		// Tambahkan baris ke buffer
		builder.WriteString(line)
		builder.WriteString("\n") // Tambahkan newline untuk setiap baris
	}

	// Periksa error pada scanner
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	return builder.String(), nil
}

func (r *schemaRepositoryImpl) InitializeDatabase(ctx context.Context, schemaFile, schemaName string) error {
	// Mulai transaction
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Muat isi file SQL
		sqlContent, err := r.LoadSQLFile(schemaFile, schemaName)
		if err != nil {
			log.Printf("Error loading SQL file: %v", err)
			return fmt.Errorf("failed to load SQL file: %w", err)
		}

		// Eksekusi SQL
		if err := tx.Exec(sqlContent).Error; err != nil {
			log.Printf("Error executing SQL: %v", err)
			return fmt.Errorf("failed to execute SQL: %w", err)
		}

		// (Opsional) Simpan informasi schema ke database
		if err := tx.Exec("INSERT INTO schema_logs (schema_name, created_at) VALUES (?, NOW())", schemaName).Error; err != nil {
			log.Printf("Error saving schema log: %v", err)
			return fmt.Errorf("failed to save schema log: %w", err)
		}

		// Commit transaction
		log.Printf("Schema %s successfully initialized", schemaName)
		return nil
	})
}

func (r *schemaRepositoryImpl) SetSchema(schemaName string) error {
	return r.db.Exec(fmt.Sprintf("SET search_path TO %s", schemaName)).Error
}
