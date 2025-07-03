package backend

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const endpoint = "https://diyhrt.market/api/listings"

func GetListings() ([]Listing, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("API_KEY environment variable not set")
	}

	// Create HTTP client
	client := &http.Client{Timeout: 10 * time.Second}

	// Create request
	req, err := http.NewRequest("GET", endpoint+"?api_token="+apiKey, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request failed: %w", err)
	}

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, string(body))
	}

	// Decode response
	var listings []Listing
	if err := json.NewDecoder(resp.Body).Decode(&listings); err != nil {
		return nil, fmt.Errorf("decoding failed: %w", err)
	}

	return listings, nil
}
