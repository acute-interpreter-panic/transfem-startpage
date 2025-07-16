package cli

import (
	"fmt"
	"os"
	"strings"

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

func getSingleArgumentString(a Argument) string {
	requiredString := ""
	if a.Required {
		requiredString = "*"
	}
	return requiredString + a.Name + ":" + a.Type
}

func getArgumentString(arguments []Argument) string {
	argumentString := color.Blue
	for _, a := range arguments {
		argumentString = argumentString + " [" + getSingleArgumentString(a) + "]"
	}
	return argumentString + color.Reset
}

func generalHelp() error {
	fmt.Println()
	fmt.Println(HelpHeader)
	fmt.Println()

	for _, p := range Programs {
		fmt.Print(color.Bold + padString(p.Name, 7) + color.Reset)
		fmt.Print(padString(getArgumentString(p.Arguments), 40) + p.ShortDescription + "\n")
	}
	return nil
}

func specificHelp(programName string) error {
	program := GetProgram(programName)

	fmt.Println(color.Bold + "MAN PAGE FOR " + strings.ToUpper(programName) + color.Reset)
	fmt.Println()

	fmt.Println(color.Purple + "transfem-startpage " + programName + color.Reset + getArgumentString(program.Arguments))
	fmt.Println()

	fmt.Println(color.Bold + "arguments" + color.Reset)

	argumentStrings := make([]string, len(program.Arguments))
	maxArgumentString := 0
	for i, a := range program.Arguments {
		s := getSingleArgumentString(a)
		argumentStrings[i] = s
		if len(s) > maxArgumentString {
			maxArgumentString = len(s)
		}
	}

	for i, a := range program.Arguments {
		fmt.Println(padString(argumentStrings[i], maxArgumentString+4) + a.Description)
	}

	fmt.Println()
	fmt.Println(program.LongDescription)

	return nil
}

func Help() error {
	if len(os.Args) > 2 {
		return specificHelp(os.Args[2])
	}

	return generalHelp()
}
