package backend

import (
	"fmt"
	"net/http"
	"os"
)

const endpoint = "https://diyhrt.market/api/listings"

func GetListings() []Listing {
	apiKey := os.Getenv("API_KEY")
	fmt.Println(apiKey)

	// why put api key in url parameter
	req, err := http.NewRequest("GET", endpoint+"?api_token="+apiKey, nil)

	if err != nil {
		fmt.Print(err.Error())
		return []Listing{}
	}

	fmt.Println(req.Body.Read())

	return []Listing{}
}
