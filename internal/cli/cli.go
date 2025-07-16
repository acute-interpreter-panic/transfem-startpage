package cli

import (
	"log"
	"os"

	"gitea.elara.ws/Hazel/transfem-startpage/internal/utils"
	"github.com/TwiN/go-color"
)

type ProgramFunction func() error
type Program struct {
	Name             string
	Function         ProgramFunction
	ShortDescription string
	LongDescription  string
	Arguments        []Argument
}
type Argument struct {
	Name        string
	Type        string
	Required    bool
	Description string
}

var HelpHeader = `This is the help page of ` + utils.Name + `.
` + color.Purple + utils.BinaryName + ` {program} {...args}` + color.Reset + `
The following Programs are available:`
var Programs = []Program{
	{
		Name:             "help",
		ShortDescription: "get more information on how the cli in general or a specific program works",
		LongDescription:  "What did you expect to find here?",
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
		LongDescription: `The start program starts the webserver.
It loads the config file of the according profile.
It uses the default values if no config file was found.`,
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
		LongDescription: `Does something with the cache.
- clear: delete the whole cache
- clean: delete all files that aren't used by any program.`,
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

func GetProgram(programName string) Program {
	for i, p := range Programs {
		if p.Name == programName {
			return Programs[i]
		}
	}

	log.Panicln("couldn't find program", programName, ". EXITING")
	return Program{}
}

func Cli() {
	// getting around initialization cycle
	Programs[0].Function = Help

	programName := "help"
	if len(os.Args) > 1 {
		programName = os.Args[1]
	}

	var selectedProgram Program = GetProgram(programName)
	err := selectedProgram.Function()
	if err != nil {
		log.Panicln(err)
	}
}
