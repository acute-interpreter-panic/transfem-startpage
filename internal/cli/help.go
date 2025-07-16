package cli

import (
	"fmt"
	"log"
)

func generalHelp() error {
	fmt.Println()
	fmt.Println(HelpHeader)
	for _, p := range Programs {
		fmt.Println(" - " + p.Name + ":\t" + p.ShortDescription)
	}
	return nil
}

func Help() error {
	log.Println("running help")

	return generalHelp()
}
