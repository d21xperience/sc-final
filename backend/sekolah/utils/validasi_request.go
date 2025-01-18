package utils

import (
	"errors"
	"log"
	"reflect"
	"strings"
)

// // ValidateFields memvalidasi beberapa field dari sebuah request
// func ValidateFields(req interface{}, fieldNames []string) error {
// 	// Debugging: Cek nilai request yang diterima
// 	log.Printf("Received data request: %+v\n", req)

// 	// Cek apakah req kosong
// 	if req == nil {
// 		log.Println("Request is nil")
// 		return errors.New("invalid request: request is nil")
// 	}

// 	// Gunakan refleksi untuk memvalidasi field tertentu
// 	val := reflect.ValueOf(req)
// 	if val.Kind() == reflect.Ptr {
// 		val = val.Elem() // Dereference pointer
// 	}
// 	if val.Kind() != reflect.Struct {
// 		return errors.New("invalid request: expected a struct")
// 	}

// 	// Periksa setiap field
// 	var missingFields []string
// 	for _, fieldName := range fieldNames {
// 		field := val.FieldByName(fieldName)
// 		if !field.IsValid() || field.IsZero() {
// 			log.Printf("%s is missing or nil in request\n", fieldName)
// 			missingFields = append(missingFields, fieldName)
// 		}
// 	}

// 	// Jika ada field yang kosong, kembalikan error
// 	if len(missingFields) > 0 {
// 		return errors.New("invalid request: missing fields - " + strings.Join(missingFields, ", "))
// 	}

// 	return nil
// }

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
