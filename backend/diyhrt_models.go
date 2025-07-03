package backend

type ActiveIngredient struct {
	Name        string `json:"name"`
	Ester       string `json:"ester"`
	DisplayName string `json:"display_name"`
}

type Product struct {
	Id               int              `json:"id"`
	Name             string           `json:"name"`
	Image            string           `json:"image"`
	ActiveIngredient ActiveIngredient `json:"active_ingredient"`
}

type Store struct {
	Id                 int    `json:"id"`
	Name               string `json:"name"`
	Url                string `json:"url"`
	Description        string `json:"description"`
	ShipsFromCountry   string `json:"ships_from_country"`
	ShipsToCountry     string `json:"ships_to_country"`
	ServiceStatus      string `json:"service_status"`
	ServiceStatusNotes string `json:"service_status_notes"`
	PaymentMethods     string `json:"payment_methods"`
	CategoryName       string `json:"category_name"`
}

type Listing struct {
	Id             int     `json:"id"`
	ProductName    string  `json:"product_name"`
	StoreName      string  `json:"store_name"`
	Price          string  `json:"price"`
	PriceCurrency  string  `json:"price_currency"`
	State          string  `json:"state"`
	InStock        bool    `json:"in_stock"`
	Url            string  `json:"url"`
	PricingPerUnit string  `json:"pricing_per_unit"`
	Product        Product `json:"product"`
	Store          Store   `json:"store"`
}
