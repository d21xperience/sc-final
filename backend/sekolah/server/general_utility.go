package server

// // RequestRequirement digunakan untuk menyimpan dependensi service dan protobuf
// type RequestRequirement[S any, P any] struct {
// 	service  S // Service untuk mengakses data
// 	protoBuf P // Protobuf atau struct untuk representasi data
// }

// func (r *RequestRequirement[S, P]) GetModelByID(ctx context.Context, id string, fetcher func(S, string) (*P, error)) (*P, error) {
// 	// Gunakan fetcher untuk mengambil data berdasarkan ID
// 	result, err := fetcher(r.service, id)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get model by ID: %w", err)
// 	}

// 	return result, nil
// }

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
