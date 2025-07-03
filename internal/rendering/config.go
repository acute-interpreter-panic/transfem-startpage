package rendering

import (
	"fmt"

	"gitea.elara.ws/Hazel/transfem-startpage/internal/diyhrt"
)

type RenderingConfig struct {
	HeaderPhrases     []string
	BackgroundScrollX string
	BackgroundScrollY string
	PageTitle         string
	SearchPlaceholder string
	SearchFormAction  string
	SearchInputName   string

	Listings []diyhrt.Listing
	Stores   []diyhrt.Store
}

func DefaultRenderingConfig() RenderingConfig {
	return RenderingConfig{
		HeaderPhrases: []string{
			"GirlJuice.Inject()",
			"Child.CrowdKill()",
			"CopCar.Burn()",
			"You.Cute = true",
			"You.Gay = true",
			"Nazi.Punch()",
			"Dolls.GiveGuns()",
		},
		BackgroundScrollX: "1",
		BackgroundScrollY: "0",
		PageTitle:         "TransRights",
		SearchPlaceholder: "Search on DuckDuckGo",
		SearchFormAction:  "https://duckduckgo.com/",
		SearchInputName:   "q",
	}
}

func (rc *RenderingConfig) LoadDiyHrt(listings []diyhrt.Listing) {
	existingStores := make(map[int]struct{})
	stores := make([]diyhrt.Store, 0)

	for _, listing := range listings {
		fmt.Println(listing)
		if _, ok := existingStores[listing.Store.Id]; ok {
			continue
		}

		stores = append(stores, listing.Store)
	}

	rc.Listings = listings
	rc.Stores = stores
}
