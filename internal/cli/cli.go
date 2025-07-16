package cli

import (
	"fmt"
	"log"
	"os"
)

type ProgramFunction func() error
type Program struct {
	Name        string
	Function    ProgramFunction
	Description string
}

var HelpHeader = `Meow
Ze
Dong`
var Programs = []Program{
	{
		Name:        "help",
		Function:    Help,
		Description: "get more information on how the cli or a program works",
	},
	{
		Name:        "start",
		Function:    Start,
		Description: "start the webserver",
	},
	{
		Name:        "cache",
		Function:    Cache,
		Description: "do something with the cache",
	},
}

func Cli() {
	fmt.Println("running transfem startpage")

	programName := "help"
	if len(os.Args) > 1 {
		programName = os.Args[1]
	}

	log.Println("running program", programName)

	var selectedProgram *Program = nil
	for i, p := range Programs {
		if p.Name == programName {
			selectedProgram = &Programs[i]
			break
		}
	}
	if selectedProgram == nil {
		log.Panicln("couldn't find program", programName, ". EXITING")
	}

	selectedProgram.Function()
}
