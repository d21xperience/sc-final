package utils

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

// func ValidateCfg(req interface{}) error {
// 	log.Printf("Received data request: %+v\n", req)
// 	var fieldNames = []string{
// 		"NetworkId",
// 		"BlockchainType", // "ethereum", "quorum", atau "hyperledger"
// 		"RPCURL",         // URL RPC untuk Ethereum/Quorum
// 		// Untuk Hyperledger Fabric
// 		"FabricConfigPath",
// 		"FabricWallet",
// 		"FabricIdentity",
// 	}
// 	// Cek apakah req kosong
// 	if req == nil {
// 		log.Println("Request is nil")
// 		return errors.New("invalid request: request is nil")
// 	}

// 	val := reflect.ValueOf(req)
// 	if val.Kind() == reflect.Ptr {
// 		val = val.Elem() // Dereference pointer
// 	}
// 	if val.Kind() != reflect.Struct {
// 		return errors.New("invalid request: expected a struct")
// 	}

// 	// Mapping nama field dalam struct ke bentuk case-insensitive
// 	fieldMap := make(map[string]reflect.Value)
// 	for i := 0; i < val.NumField(); i++ {
// 		field := val.Type().Field(i)
// 		fieldMap[strings.ToLower(field.Name)] = val.Field(i)
// 	}

// 	var fieldValue reflect.Value
// 	var exists bool
// 	// Periksa setiap field
// 	var missingFields []string
// 	for _, fieldName := range fieldNames {
// 		fieldValue, exists = fieldMap[strings.ToLower(fieldName)]
// 		if !exists || fieldValue.IsZero() {
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

func ConvertModelsToPB[T any, U any](models []*T, converter func(*T) *U) []*U {
	var pbList []*U
	for _, model := range models {
		pbList = append(pbList, converter(model))
	}
	return pbList
}
func ConvertPBToModels[T any, U any](pbs []*T, converter func(*T) *U) []*U {
	var modelList []*U
	for _, model := range pbs {
		modelList = append(modelList, converter(model))
	}
	return modelList
}

func ConvertModelToPB[T any, U any](model *T, converter func(*T) *U) *U {
	if model == nil {
		return nil
	}
	return converter(model)
}

// ConvertStringToUint mengonversi string ke tipe bilangan unsigned (generic)
func ConvertStringToUint[T uint8 | uint16 | uint32 | uint64](str string) (T, error) {
	num, err := strconv.ParseUint(str, 10, strconv.IntSize)
	if err != nil {
		return 0, fmt.Errorf("konversi gagal: %w", err)
	}
	return T(num), nil
}

// ConvertUintToString mengonversi tipe uint (uint8, uint16, uint32, uint64) ke string
func ConvertUintToString[T uint8 | uint16 | uint32 | uint64](num T) string {
	return strconv.FormatUint(uint64(num), 10)
}
