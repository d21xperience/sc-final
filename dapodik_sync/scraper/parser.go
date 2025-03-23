package scraper

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

// ExtractSemesterOptions mengekstrak daftar semester dari dokumen HTML
func ExtractSemesterOptions(doc *goquery.Document) []string {
	var semesters []string

	// Cari elemen <select name="semester_id"> lalu ambil semua <option>
	doc.Find("select[name='semester_id'] option").Each(func(i int, s *goquery.Selection) {
		value, exists := s.Attr("value")
		if exists {
			text := s.Text()
			semesters = append(semesters, fmt.Sprintf("%s - %s", value, text))
		}
	})

	return semesters
}
