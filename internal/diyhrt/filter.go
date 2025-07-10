package diyhrt

import (
	"fmt"
	"slices"
)

type StoreFilter struct{
    Limit int

    IncludeIds []int
    ExcludeIds []int
}


func (f StoreFilter) Filter (stores []Store) []Store {
    result := make([]Store, 0)

    for _, s := range stores {
        if f.Limit > 0 && len(result) >= f.Limit {
            break
        }

        if len(f.IncludeIds) > 0 {
            if slices.Contains(f.IncludeIds, s.Id) {
                result = append(result, s)
            }
            continue
        }


        if slices.Contains(f.ExcludeIds, s.Id) {
            continue
        }


        result = append(result, s)
    }

    fmt.Println(len(result))
    return  result
}
