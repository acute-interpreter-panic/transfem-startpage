package cli

import (
	"fmt"
	"log"

	"github.com/TwiN/go-color"
)

func padString(s string, n int) string {
	missing := n - len(s)
	if missing <= 0 {
		return s
	}

	for _ = range missing {
		s = s + " "
	}

	return s
}

func generalHelp() error {
	fmt.Println()
	fmt.Println(HelpHeader)
	fmt.Println()

	for _, p := range Programs {
		fmt.Print(color.Bold + padString(p.Name, 7) + color.Reset)

		argumentString := color.Purple
		for _, a := range p.Arguments {
			requiredString := ""
			if a.Required {
				requiredString = "*"
			}
			argumentString = argumentString + " [" + requiredString + a.Name + ":" + a.Type + "]"
		}
		argumentString = argumentString + color.Reset

		fmt.Print(padString(argumentString, 40) + p.ShortDescription + "\n")
	}
	return nil
}

func Help() error {
	log.Println("running help")

	return generalHelp()
}
