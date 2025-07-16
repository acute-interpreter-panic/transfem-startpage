package cli

import (
	"fmt"
	"log"
	"os"
)

type ProgramFunction func() error
type Program struct {
	Name             string
	Function         ProgramFunction
	ShortDescription string
}

var HelpHeader = `This is the help page of transfem-startpage.
transfem-startpage {program} {...args}
The following Programs are available:`
var Programs = []Program{
	{
		Name:             "help",
		ShortDescription: "get more information on how the cli in general or a specific program works",
	},
	{
		Name:             "start",
		Function:         Start,
		ShortDescription: "start the webserver",
	},
	{
		Name:             "cache",
		Function:         Cache,
		ShortDescription: "do something with the cache",
	},
}

func Cli() {
	fmt.Println("running transfem startpage")

	// getting around initialization cycle
	Programs[0].Function = Help

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
