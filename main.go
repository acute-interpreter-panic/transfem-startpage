package main

import (
	"fmt"

	"gitea.elara.ws/Hazel/transfem-startpage/internal/diyhrt"
)

func main() {
	fmt.Println("running transfem startpage")

	listings, err := diyhrt.GetListings()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, l := range listings {
		fmt.Println(l)
	}
}
