package backend

type ActiveIngredient struct {
	Name        string
	Ester       string
	DisplayName string
}

type Product struct {
	Id               int
	Name             string
	Image            string
	ActiveIngredient ActiveIngredient
}

type Store struct {
	Id                 int
	Name               string
	Url                string
	Description        string
	ShipsFromCountry   string
	ShipsToCountry     string
	ServiceStatus      string
	ServiceStatusNotes string
	PaymentMethods     string
	CategoryName       string
}

type Listing struct {
	Id             int
	ProductName    string
	StoreName      string
	Price          string
	PriceCurrency  string
	State          string
	InStock        bool
	Url            string
	PricingPerUnit string
	Product        Product
	Store          Store
}
