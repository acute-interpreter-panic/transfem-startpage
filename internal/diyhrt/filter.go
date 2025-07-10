package diyhrt

import (
	"slices"
)

type StoreFilter struct{
    Limit int

    IncludeIds []int
    ExcludeIds []int

    ShipsTo []string
}


func (f StoreFilter) Filter (stores []Store) []Store {
    result := make([]Store, 0)

    if len(f.IncludeIds) > 0 {
        for _, s := range stores {
            if f.Limit > 0 && len(result) >= f.Limit {
                break
            }

            if slices.Contains(f.IncludeIds, s.Id) {
                result = append(result, s)
            }
        }
    }


    for _, s := range stores {
        if f.Limit > 0 && len(result) >= f.Limit {
            break
        }


        if slices.Contains(f.ExcludeIds, s.Id) || slices.Contains(f.IncludeIds, s.Id) {
            continue
        }

        result = append(result, s)
    }

    return  result
}

type ListingFilter struct{
    Limit int

    IncludeIds []int
    ExcludeIds []int

    FromStores []int
}


func (f ListingFilter) Filter (listings []Listing) []Listing {
    result := make([]Listing, 0)

    if len(f.IncludeIds) > 0 {
        for _, l := range listings {
            if f.Limit > 0 && len(result) >= f.Limit {
                break
            }

            if slices.Contains(f.IncludeIds, l.Id) {
                result = append(result, l)
            }
        }
    }


    for _, l := range listings {
        if f.Limit > 0 && len(result) >= f.Limit {
            break
        }


        if slices.Contains(f.ExcludeIds, l.Id) || slices.Contains(f.IncludeIds, l.Id) {
            continue
        }

        if len(f.FromStores) > 0 && !slices.Contains(f.FromStores, l.Store.Id) {
            continue
        }

        result = append(result, l)
    }

    return  result
}
