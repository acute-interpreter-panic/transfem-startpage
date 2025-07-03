package rendering

import (
	"html/template"
	"log"
	"os"
)

func getFileContent() string {
	content, err := os.ReadFile("frontend/index.html")

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

var IndexTemplate = template.Must(template.New("index").Parse(getFileContent()))
