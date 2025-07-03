package backend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const endpoint = "https://diyhrt.market/api/listings"

func GetListings() ([]Listing, error) {
	apiKey := os.Getenv("API_KEY")

	// why put api key in url parameter
	resp, err := http.NewRequest("GET", endpoint+"?api_token="+apiKey, nil)

	if err != nil {
		fmt.Print(err.Error())
		return []Listing{}, err
	}

	var listings []Listing
	if err := json.NewDecoder(resp.Body).Decode(&listings); err != nil {
		return []Listing{}, err
	}

	return listings, nil
}
