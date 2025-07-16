package server

import (
	"bytes"
	"embed"
	"io/fs"
	"log"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

var FrontendFiles embed.FS

func getFileContent() string {
	content, err := FrontendFiles.ReadFile("frontend/index.html")

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func getIndex(c echo.Context) error {
	IndexTemplate := template.Must(template.New("index").Parse(getFileContent()))

	var tpl bytes.Buffer
	IndexTemplate.Execute(&tpl, Config.Template)

	return c.HTML(http.StatusOK, tpl.String())
}

func getFileSystem() http.FileSystem {
	fsys, err := fs.Sub(FrontendFiles, "frontend")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
