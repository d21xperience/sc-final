package utils

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
