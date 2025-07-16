package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/TwiN/go-color"
)

type ProgramFunction func() error
type Program struct {
	Name             string
	Function         ProgramFunction
	ShortDescription string
	Arguments        []Argument
}
type Argument struct {
	Name        string
	Type        string
	Required    bool
	Description string
}

var HelpHeader = `This is the help page of transfem-startpage.
` + color.Bold + `transfem-startpage {program} {...args}` + color.Reset + `
The following Programs are available:`
var Programs = []Program{
	{
		Name:             "help",
		ShortDescription: "get more information on how the cli in general or a specific program works",
		Arguments: []Argument{
			{
				Name:        "program",
				Type:        "string",
				Required:    false,
				Description: "defines the program you want to know more about",
			},
		},
	},
	{
		Name:             "start",
		Function:         Start,
		ShortDescription: "start the webserver",
		Arguments: []Argument{
			{
				Name:        "profile",
				Type:        "string",
				Required:    false,
				Description: "tells the program which config to load, default is 'default'",
			},
		},
	},
	{
		Name:             "cache",
		Function:         Cache,
		ShortDescription: "do something with the cache",
		Arguments: []Argument{
			{
				Name:        "action",
				Type:        "enum(clear;clean)",
				Required:    true,
				Description: "defines what to do with the cache",
			},
		},
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
