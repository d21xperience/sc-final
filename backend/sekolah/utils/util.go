package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// ConvertModelsToPB mengonversi slice model ke slice protobuf
func ConvertModelsToPB[T any, U any](models []T, convert func(T) U) []U {
	var pbModels []U
	for _, model := range models {
		pbModels = append(pbModels, convert(model))
	}
	return pbModels
}

//	func ConvertModelsToPB[T any, U any](models []*T, converter func(*T) *U) []*U {
//		var pbList []*U
//		for _, model := range models {
//			pbList = append(pbList, converter(model))
//		}
//		return pbList
//	}
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

// Fungsi helper untuk mengubah string ke int
func ParseInt(value string) int {
	i, _ := strconv.Atoi(value)
	return i
}

// Fungsi helper untuk mengubah string ke UUID
func ParseUuid(value *string) uuid.UUID {
	i, _ := uuid.Parse(*value)
	return i
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

// FUNGSI untuk Menangani ERROR pada POINTER

// fungsi helper untuk menangani pointer nil pada string
func SafeString(s *string) string {
	if s == nil {
		return "" // atau return nilai default lain yang sesuai
	}
	return *s
}

// Generic function to convert string to time.Time
func StringToTime[T ~string](str T, layout string) (time.Time, error) {
	parsedTime, err := time.Parse(layout, string(str))
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

// Function to convert time.Time to string (Tidak perlu generic)
func TimeToString(t time.Time, layout string) string {
	return t.Format(layout)
}

// Fungsi generic untuk memastikan semua pointer string dalam struct tidak nil
func HandleNilPointers[T any](data *T) {
	val := reflect.ValueOf(data).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		// Cek apakah field adalah pointer ke string (*string)
		if field.Kind() == reflect.Ptr && field.Type().Elem().Kind() == reflect.String {
			if field.IsNil() {
				emptyStr := ""
				field.Set(reflect.ValueOf(&emptyStr)) // Set pointer ke string kosong
			}
		}
	}
}
