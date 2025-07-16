package diyhrt

type DiyHrtConfig struct {
	ApiKey         string
	FetchIntervals int

	StoreFilter   StoreFilter
	ListingFilter ListingFilter
}
