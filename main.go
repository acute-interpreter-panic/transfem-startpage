package main

import (
	"fmt"

	"gitea.elara.ws/Hazel/transfem-startpage/backend"
)

func main() {
	fmt.Println("running transfem startpage")

	listings, _ := backend.GetListings()
	for _, l := range listings {
		fmt.Println(l)
	}
}
