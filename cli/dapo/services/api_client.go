package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"time"
)

// APIClient represents a client for making API requests
type APIClient struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

// NewAPIClient creates a new APIClient instance
func NewAPIClient(baseURL, token string) *APIClient {
	return &APIClient{
		baseURL: baseURL,
		token:   token,
		httpClient: &http.Client{
			Timeout: 10 * time.Second, // Always set a timeout
		},
	}
}

// RequestOptions contains options for making HTTP requests
type RequestOptions struct {
	Method      string
	Path        string
	QueryParams map[string]string
	Headers     map[string]string
	Body        interface{}
}

// Response contains the HTTP response details
type Response struct {
	StatusCode int
	Body       []byte
	Headers    http.Header
}

const outputDir = "./data/"

// DoRequest makes an HTTP request with the given options
func (c *APIClient) DoRequest(opts RequestOptions) (*Response, error) {
	// Build the full URL
	url := c.baseURL + opts.Path
	if len(opts.QueryParams) > 0 {
		url += "?"
		params := ""
		for key, value := range opts.QueryParams {
			if params != "" {
				params += "&"
			}
			params += key + "=" + value
		}
		url += params
	}

	// Prepare the request body if any
	var bodyReader io.Reader
	if opts.Body != nil {
		bodyBytes, err := json.Marshal(opts.Body)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request body: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	// Create the request
	req, err := http.NewRequest(opts.Method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Set default headers
	req.Header.Add("Authorization", "Bearer "+c.token)
	req.Header.Add("Content-Type", "application/json")

	// Add custom headers if any
	for key, value := range opts.Headers {
		req.Header.Add(key, value)
	}

	// Execute the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Return the response
	return &Response{
		StatusCode: resp.StatusCode,
		Body:       body,
		Headers:    resp.Header,
	}, nil
}

// versi 1.1
func GetOrFetch[T any](client *APIClient, npsn, semesterID, outputDir, path string, query map[string]string) (*T, error) {
	// Generate filename dari parameter
	endpoint := filepath.Base(path)
	filename := fmt.Sprintf("%s_%s_%s.json", endpoint, npsn, semesterID)
	filePath := filepath.Join(outputDir, filename)

	if FileExists(filePath) {
		fmt.Printf("Reading existing file: %s\n", filePath)

		var data T
		if err := ReadJSONToStruct(filePath, &data); err != nil {
			return nil, fmt.Errorf("error reading JSON file: %w", err)
		}
		return &data, nil
	}

	fmt.Printf("Fetching data from API: %s?%v\n", path, query)
	opts := RequestOptions{
		Method:      "GET",
		Path:        path,
		QueryParams: query,
	}

	resp, err := client.DoRequest(opts)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if err := SaveToFile(resp.Body, filePath); err != nil {
		return nil, fmt.Errorf("error saving to file: %w", err)
	}

	fmt.Printf("Successfully saved response to: %s\n", filePath)

	var data T
	if err := json.Unmarshal(resp.Body, &data); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return &data, nil
}
