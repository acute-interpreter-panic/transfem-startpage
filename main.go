package main

import (
	"fmt"

	"gitea.elara.ws/Hazel/transfem-startpage/backend"
)

func main() {
	fmt.Println("running transfem startpage")

	for _, l := range backend.GetListings() {
		fmt.Println(l)
	}
}
