package diyhrt

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"
)

const endpoint = "https://diyhrt.market/api/listings"

func GetListings() ([]Listing, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return nil, errors.New("API_KEY environment variable not set")
	}

	// Create HTTP client
	client := &http.Client{Timeout: 10 * time.Second}

	// Create request
	req, err := http.NewRequest("GET", endpoint+"?api_token="+apiKey, nil)
	if err != nil {
		return nil, err
	}

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected status code")
	}

	// Decode response
	var listings []Listing
	if err := json.NewDecoder(resp.Body).Decode(&listings); err != nil {
		return nil, err
	}

	return listings, nil
}
