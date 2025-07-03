package rendering

import "reflect"

type RenderingConfig struct {
	HeaderPhrases     []string
	BackgroundScrollX string
	BackgroundScrollY string
	PageTitle         string
	SearchPlaceholder string
	SearchFormAction  string
	SearchInputName   string
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

func (rc *RenderingConfig) Set(key string, value string) {
	// https://gist.github.com/kilfu0701/77c614386483782f68bc5538b6100730
	r := reflect.ValueOf(rc)
	f := reflect.Indirect(r).FieldByName(key)
	if f.Kind() != reflect.Invalid {
		f.SetString(value)
	}
}
