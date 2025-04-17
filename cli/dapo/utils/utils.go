package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// =======================
// KONVERSI PROTOBUF
// =======================
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

// konversi pb ke model
// schemaName := req.GetSchemaName()
// pbs := req.pb  //repeated proto
// models := ConvertPBToModels(pbs, func(items *pb.pbs) *models.models {}
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

// =======================
// KONVERSI INT
// =======================
// Fungsi helper untuk mengubah string ke int
func StringToInt(value string) int {
	i, _ := strconv.Atoi(value)
	return i
}
func StringToInt32(s string) (int32, error) {
	val, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(val), nil
}

// =======================
// KONVERSI UUID
// =======================
// Fungsi helper untuk mengubah pointer string ke UUID
func ParseUuidFromPointerString(value *string) uuid.UUID {
	i, _ := uuid.Parse(*value)
	return i
}

// Fungsi helper untuk mengubah  string ke UUID
func StringToUUID(value string) uuid.UUID {
	i, _ := uuid.Parse(value)
	return i
}
func PointerToUUID(value string) *uuid.UUID {
	i, err := uuid.Parse(value)
	if err != nil {
		return nil
	}
	return &i
}
func UUIDToPointer(u uuid.UUID) *uuid.UUID {
	return &u
}

// Mengembalikan uuid.UUID dari pointer string. Jika nil atau invalid, kembalikan uuid.Nil
func PointerStringToUUID(s *string) uuid.UUID {
	if s == nil || *s == "" {
		return uuid.Nil
	}
	id, err := uuid.Parse(*s)
	if err != nil {
		return uuid.Nil
	}
	return id
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

// =======================
// KONVERSI LAINNYA
// =======================

// FUNGSI untuk Menangani ERROR pada POINTER

// fungsi helper untuk menangani pointer nil pada string
func SafeString(s *string) string {
	if s == nil {
		return "" // atau return nilai default lain yang sesuai
	}
	return *s
}

// =======================
// KONVERSI UINT
// =======================
// ToUint16Pointer mengubah nilai uint16 menjadi pointer
func Uint16ToPointer(value uint16) *uint16 {
	return &value
}

// ToUint32Pointer mengubah nilai uint32 menjadi pointer
func Uint32ToPointer(value uint32) *uint32 {
	return &value
}

// ToUint64Pointer mengubah nilai uint64 menjadi pointer
func Uint64ToPointer(value uint64) *uint64 {
	return &value
}

// Uint16Value mengembalikan nilai uint16 dari pointer, jika nil akan default ke 0
func PointerToUint16(ptr *uint16) uint16 {
	if ptr == nil {
		return 0
	}
	return *ptr
}

// Uint32Value mengembalikan nilai uint32 dari pointer, jika nil akan default ke 0
func PointerToUint32(ptr *uint32) uint32 {
	if ptr == nil {
		return 0
	}
	return *ptr
}
func Uint16ToUint32Pointer(ptr *uint16) *uint32 {
	if ptr == nil {
		return nil
	}
	val := uint32(*ptr)
	return &val
}

// Uint64Value mengembalikan nilai uint64 dari pointer, jika nil akan default ke 0
func PointerToUint64(ptr *uint64) uint64 {
	if ptr == nil {
		return 0
	}
	return *ptr
}

// ConvertStringToUint mengonversi string ke tipe bilangan unsigned (generic)
func StringToUint[T uint8 | uint16 | uint32 | uint64](str string) (T, error) {
	num, err := strconv.ParseUint(str, 10, strconv.IntSize)
	if err != nil {
		return 0, fmt.Errorf("konversi gagal: %w", err)
	}
	return T(num), nil
}

// ConvertUintToString mengonversi tipe uint (uint8, uint16, uint32, uint64) ke string
func UintToString[T uint8 | uint16 | uint32 | uint64](num T) string {
	return strconv.FormatUint(uint64(num), 10)
}

// =======================
// KONVERSI WAKTU
// =======================
func TimeToPointer(dateStr string) *time.Time {
	if dateStr == "" {
		return nil
	}
	parsedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil // Bisa juga return error jika ingin menangani kesalahan parsing
	}
	return &parsedDate
}

// Function to convert time.Time to string (Tidak perlu generic)
func TimeToString(t time.Time, layout string) string {
	return t.Format(layout)
}

// Generic function to convert string to time.Time
func StringToTime[T ~string](str T, layout string) (time.Time, error) {
	parsedTime, err := time.Parse(layout, string(str))
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

// =======================
// KONVERSI
// =======================
// Fungsi generic untuk memastikan semua pointer string dalam struct tidak nil
func HandleNilPointers[T any](data *T) {
	val := reflect.ValueOf(data).Elem()

	for i := range val.NumField() {
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

// =======================
// KONVERSI SLICE
// =======================
// DereferenceSlice mengonversi slice pointer []*T menjadi slice biasa []T.
func PointerToSlice[T any](input []*T) []T {
	output := make([]T, len(input))
	for i, v := range input {
		if v != nil {
			output[i] = *v
		}
	}
	return output
}
func SliceToPointer[T any](input []T) []*T {
	output := make([]*T, len(input))
	for i, v := range input {
		val := v // buat salinan lokal
		output[i] = &val
	}
	return output
}
