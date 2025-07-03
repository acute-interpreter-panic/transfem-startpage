package backend

type ActiveIngredient struct {
	Name        string `json:"name"`
	Ester       string `json:"ester,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
}

type Product struct {
	Id               int              `json:"id,omitempty"`
	Name             string           `json:"name"`
	Image            string           `json:"image,omitempty"`
	ActiveIngredient ActiveIngredient `json:"active_ingredient"`
}

type Store struct {
	Id                 int    `json:"id,omitempty"`
	Name               string `json:"name"`
	Url                string `json:"url"`
	Description        string `json:"description"`
	ShipsFromCountry   string `json:"ships_from_country,omitempty"`
	ShipsToCountry     string `json:"ships_to_country,omitempty"`
	ServiceStatus      string `json:"service_status"`
	ServiceStatusNotes string `json:"service_status_notes,omitempty"`
	PaymentMethods     string `json:"payment_methods,omitempty"`
	CategoryName       string `json:"category_name,omitempty"`
}

type Listing struct {
	Id             int     `json:"id,omitempty"`
	ProductName    string  `json:"product_name,omitempty"`
	StoreName      string  `json:"store_name"`
	Price          string  `json:"price,omitempty"`
	PriceCurrency  string  `json:"price_currency,omitempty"`
	State          string  `json:"state,omitempty"`
	InStock        bool    `json:"in_stock"`
	Url            string  `json:"url,omitempty"`
	PricingPerUnit string  `json:"pricing_per_unit,omitempty"`
	Product        Product `json:"product"`
	Store          Store   `json:"store"`
}
