package scraper

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// FetchHTML mengambil HTML dari URL yang diberikan dan mengembalikan objek *goquery.Document
func FetchHTML(url string) (*goquery.Document, error) {
	// Kirim request HTTP GET
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil halaman: %w", err)
	}
	defer resp.Body.Close()

	// Parsing HTML dengan goquery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca HTML: %w", err)
	}

	return doc, nil
}

// fetchData mengambil data dari URL yang diberikan dan mengembalikan hasilnya dalam bentuk map
func FetchData(url, bearerToken string) (map[string]interface{}, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+bearerToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Membaca body respons
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Mencetak body respons mentah untuk debugging
	fmt.Printf("\n[DEBUG] Data mentah dari %s:\n%s\n", url, string(body))

	// Parsing JSON
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}
