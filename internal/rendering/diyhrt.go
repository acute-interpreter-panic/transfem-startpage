package rendering

import (
	"maps"
	"slices"

	"gitea.elara.ws/Hazel/transfem-startpage/internal/diyhrt"
)

func (c *Config) LoadDiyHrt(listings []diyhrt.Listing) {
	existingStores := make(map[int]diyhrt.Store)

	for _, listing := range listings {
		existingStores[listing.Store.Id] = listing.Store
	}

	c.Template.Listings = c.DiyHrt.ListingFilter.Filter(listings)
	c.Template.Stores = c.DiyHrt.StoreFilter.Filter(slices.Collect(maps.Values(existingStores)))
}

func (c *Config) FetchDiyHrt() error {
	l, err := diyhrt.GetListings(c.DiyHrt.ApiKey)
	if err != nil {
		return err
	}
	c.LoadDiyHrt(l)
	return nil
}
