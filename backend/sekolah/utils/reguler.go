package utils

import (
	"fmt"
	"regexp"
)

func SearchString(input string) string {
	// input := "invalid request: missing fields - TahunAjaran"

	// Gunakan regex untuk mencari kata setelah "- "
	re := regexp.MustCompile(`- (\w+)`)
	match := re.FindStringSubmatch(input)

	if len(match) > 1 {
		field := match[1]
		fmt.Println("Field:", field) // Output: TahunAjaran
		return field
	} else {
		fmt.Println("Tidak ada field ditemukan")
		return ""
	}
}

func ConvertModelsToPB[T any, U any](models []*T, converter func(*T) *U) []*U {
	var pbList []*U
	for _, model := range models {
		pbList = append(pbList, converter(model))
	}
	return pbList
}
