package rendering

import (
	"errors"
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"slices"

	"gitea.elara.ws/Hazel/transfem-startpage/internal/diyhrt"
	"github.com/kirsle/configdir"
	"github.com/pelletier/go-toml"
)

type RenderingConfig struct {
	HeaderPhrases     []string
	BackgroundScrollX string
	BackgroundScrollY string
	PageTitle         string
	SearchPlaceholder string
	SearchFormAction  string
	SearchInputName   string

	StoreFilter diyhrt.StoreFilter
	ListingFilter diyhrt.ListingFilter

	Listings []diyhrt.Listing
	Stores   []diyhrt.Store
}

func DefaultRenderingConfig() RenderingConfig {
	return RenderingConfig{
		HeaderPhrases: []string{
			"GirlJuice.Inject();",
			"Child.CrowdKill();",
			"CopCar.Burn();",
			"You.Cute = true;",
			"You.Gay = true;",
			"Nazi.Punch();",
			"Dolls.GiveGuns();",
		},
		BackgroundScrollX: "1",
		BackgroundScrollY: "0",
		PageTitle:         "TransRights",
		SearchPlaceholder: "Search on DuckDuckGo",
		SearchFormAction:  "https://duckduckgo.com/",
		SearchInputName:   "q",

		StoreFilter: diyhrt.StoreFilter{
			Limit: 0,
			IncludeIds: []int{7},
		},

		ListingFilter: diyhrt.ListingFilter{
			FromStores: []int{7},
		},
	}
}

func (rc *RenderingConfig) LoadDiyHrt(listings []diyhrt.Listing) {
	existingStores := make(map[int]diyhrt.Store)

	for _, listing := range listings {
		existingStores[listing.Store.Id] = listing.Store
	}

	rc.Listings = rc.ListingFilter.Filter(listings)
	rc.Stores = rc.StoreFilter.Filter(slices.Collect(maps.Values(existingStores)))
}


func (rc *RenderingConfig) ScanForConfigFile(profile string) error {
	profileFile := profile + ".toml"

	configPath := configdir.LocalConfig("startpage")
	configFile := filepath.Join(configPath, profileFile)

	if err := rc.LoadConfigFile(configFile); !errors.Is(err, os.ErrNotExist) {
		return err
	}

	if err := rc.LoadConfigFile(profileFile); !errors.Is(err, os.ErrNotExist) {
		return err
	}

	if err := rc.LoadConfigFile("." + profileFile); !errors.Is(err, os.ErrNotExist) {
		return err
	}

	return errors.New("No config file found")
}

func (rc *RenderingConfig) LoadConfigFile(file string) error {
	if _, err := os.Stat(file); err != nil {
		return err
	}

	fmt.Println("loading config file: " + file)


	content, err := os.ReadFile(file)

	if err != nil {
		return err
	}

	err = toml.Unmarshal(content, rc)
	if err != nil {
		return err
    }
	return nil
}
