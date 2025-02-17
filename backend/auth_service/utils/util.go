package utils

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// ConvertUintToString mengonversi tipe uint (uint8, uint16, uint32, uint64) ke string
func ConvertUintToString[T uint8 | uint16 | uint32 | uint64](num T) string {
	return strconv.FormatUint(uint64(num), 10)
}

func ValidateFields(req interface{}, fieldNames []string) error {
	log.Printf("Received data request: %+v\n", req)

	// Cek apakah req kosong
	if req == nil {
		log.Println("Request is nil")
		return errors.New("invalid request: request is nil")
	}

	val := reflect.ValueOf(req)
	if val.Kind() == reflect.Ptr {
		val = val.Elem() // Dereference pointer
	}
	if val.Kind() != reflect.Struct {
		return errors.New("invalid request: expected a struct")
	}

	// Mapping nama field dalam struct ke bentuk case-insensitive
	fieldMap := make(map[string]reflect.Value)
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		fieldMap[strings.ToLower(field.Name)] = val.Field(i)
	}

	// Periksa setiap field
	var missingFields []string
	for _, fieldName := range fieldNames {
		fieldValue, exists := fieldMap[strings.ToLower(fieldName)]
		if !exists || fieldValue.IsZero() {
			log.Printf("%s is missing or nil in request\n", fieldName)
			missingFields = append(missingFields, fieldName)
		}
	}

	// Jika ada field yang kosong, kembalikan error
	if len(missingFields) > 0 {
		return errors.New("invalid request: missing fields - " + strings.Join(missingFields, ", "))
	}

	return nil
}

// Generic function untuk konversi antara string dan UUID
func ConvertUUIDToStringViceVersa[T any](input T) (any, error) {
	switch v := any(input).(type) {
	case string:
		// Konversi string ke UUID
		parsedUUID, err := uuid.Parse(v)
		if err != nil {
			return nil, fmt.Errorf("gagal mengonversi string ke UUID: %w", err)
		}
		return parsedUUID, nil
	case uuid.UUID:
		// Konversi UUID ke string
		return v.String(), nil
	default:
		return nil, fmt.Errorf("tipe tidak didukung: %T", v)
	}
}
