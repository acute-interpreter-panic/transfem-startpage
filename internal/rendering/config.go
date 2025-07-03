package rendering

type RenderingConfig struct {
	HeaderPhrases     []string
	BackgroundScrollX float32
	BackgroundScrollY float32
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
		BackgroundScrollX: 1,
		BackgroundScrollY: 0,
		PageTitle:         "TransRights",
		SearchPlaceholder: "Search on DuckDuckGo",
		SearchFormAction:  "https://duckduckgo.com/",
		SearchInputName:   "q",
	}
}
