package rendering

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"gitea.elara.ws/Hazel/transfem-startpage/internal/diyhrt"
	"github.com/pelletier/go-toml"
)

type ActiveCard string

const (
	DiyHrtStores   ActiveCard = "stores"
	DiyHrtListings ActiveCard = "listings"
	Websites       ActiveCard = "websites"
)

type AnchorTarget string

const (
	OpenInNewTab     AnchorTarget = "_blank"
	OpenInCurrentTab AnchorTarget = "_self"
	OpenInParent     AnchorTarget = "_parent"
	OpenInTopWindow  AnchorTarget = "_top"
)

type ServerConfig struct {
	Port int
}

type TemplateConfig struct {
	HeaderPhrases     []string
	BackgroundScrollX string
	BackgroundScrollY string
	PageTitle         string
	SearchPlaceholder string
	SearchFormAction  string
	SearchInputName   string

	Listings     []diyhrt.Listing
	Stores       []diyhrt.Store
	DiyHrtTarget AnchorTarget

	ActiveCard ActiveCard

	Websites      []Website
	WebsiteTarget AnchorTarget
}

type Config struct {
	Server   ServerConfig
	Template TemplateConfig
	DiyHrt   diyhrt.DiyHrtConfig
}

func NewConfig() Config {
	return Config{
		Server: ServerConfig{
			Port: 5500,
		},
		DiyHrt: diyhrt.DiyHrtConfig{
			ApiKey: os.Getenv("API_KEY"),
			StoreFilter: diyhrt.StoreFilter{
				Limit:      0,
				IncludeIds: []int{7},
			},
			ListingFilter: diyhrt.ListingFilter{
				FromStores: []int{7},
			},
		},
		Template: TemplateConfig{
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

			ActiveCard: DiyHrtListings,

			DiyHrtTarget: OpenInCurrentTab,

			Websites: []Website{
				{Url: "https://gitea.elara.ws/Hazel/transfem-startpage", Name: "Transfem Startpage", ImageUrl: "https://gitea.elara.ws/assets/img/logo.svg"},
			},
			WebsiteTarget: OpenInCurrentTab,
		},
	}
}

func (rc *Config) ScanForConfigFile(profile string) error {
	profileFile := profile + ".toml"

	baseDir, cacheDirErr := os.UserConfigDir()
	if cacheDirErr == nil {
		configFile := filepath.Join(baseDir, "startpage", profileFile)

		if err := rc.LoadConfigFile(configFile); !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}

	if err := rc.LoadConfigFile(profileFile); !errors.Is(err, os.ErrNotExist) {
		return err
	}

	if err := rc.LoadConfigFile("." + profileFile); !errors.Is(err, os.ErrNotExist) {
		return err
	}

	return errors.New("no config file found")
}

func (rc *Config) LoadConfigFile(file string) error {
	if _, err := os.Stat(file); err != nil {
		return err
	}

	log.Println("loading config file", file)

	content, err := os.ReadFile(file)

	if err != nil {
		return err
	}

	return toml.Unmarshal(content, rc)
}

func (c *Config) Init() error {
	log.Println("downloading website icons...")
	for i := range c.Template.Websites {
		c.Template.Websites[i].Cache()
	}

	return nil
}
