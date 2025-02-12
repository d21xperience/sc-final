package utils

import (
	"strconv"

	"github.com/google/uuid"
)

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

// Fungsi helper untuk mengubah string ke int
func ParseInt(value string) int {
	i, _ := strconv.Atoi(value)
	return i
}

// Fungsi helper untuk mengubah string ke int
func ParseUuid(value *string) uuid.UUID {
	i, _ := uuid.Parse(*value)
	return i
}
