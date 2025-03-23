package scraper

import "fmt"

// ScrapeSemesters mengambil data semester dari URL yang diberikan
func ScrapeSemesters(url string) ([]string, error) {
	doc, err := FetchHTML(url)
	if err != nil {
		return nil, err
	}

	semesters := ExtractSemesterOptions(doc)
	if len(semesters) == 0 {
		return nil, fmt.Errorf("tidak ada data semester yang ditemukan")
	}

	return semesters, nil
}
