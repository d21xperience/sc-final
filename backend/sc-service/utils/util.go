package utils

import (
	"errors"
	"log"
	"reflect"
	"strings"
)

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
