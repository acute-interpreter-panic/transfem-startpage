package main

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"

	"gitea.elara.ws/Hazel/transfem-startpage/internal/diyhrt"
	"gitea.elara.ws/Hazel/transfem-startpage/internal/rendering"
	"github.com/labstack/echo/v4"
)

var CurrentConfig = rendering.NewConfig()

func FetchDiyHrt() error {
	fmt.Println("Fetch DiyHrt Marketplaces...")

	l, err := diyhrt.GetListings(CurrentConfig.DiyHrt.ApiKey)
	if err != nil {
		return err
	}
	CurrentConfig.LoadDiyHrt(l)
	return nil
}

//go:embed frontend/*
var frontendFiles embed.FS

func getFileContent() string {
	content, err := frontendFiles.ReadFile("frontend/index.html")

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

var IndexTemplate = template.Must(template.New("index").Parse(getFileContent()))

func getIndex(c echo.Context) error {
	var tpl bytes.Buffer
	IndexTemplate.Execute(&tpl, CurrentConfig.Template)

	return c.HTML(http.StatusOK, tpl.String())
}

func getFileSystem() http.FileSystem {
	fsys, err := fs.Sub(frontendFiles, "frontend")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}

func main() {
	profile := "default"
	if len(os.Args) > 1 {
		profile = os.Args[1]
	}
	fmt.Println("loading profile " + profile)

	err := CurrentConfig.ScanForConfigFile(profile)
	if err != nil {
		fmt.Println(err)
	}

	err = CurrentConfig.Init()
	if err != nil {
		fmt.Println(err)
	}

	err = FetchDiyHrt()
	if err != nil {
		fmt.Println(err)
	}

	e := echo.New()

	// statically serve the file
	cacheDir, err := rendering.GetCacheDir()
	if err == nil {
		e.Static("/cache", cacheDir)
	} else {
		fmt.Println(err)
	}

	// https://echo.labstack.com/docs/cookbook/embed-resources
	staticHandler := http.FileServer(getFileSystem())
	e.GET("/assets/*", echo.WrapHandler(http.StripPrefix("/", staticHandler)))
	e.GET("/scripts/*", echo.WrapHandler(http.StripPrefix("/", staticHandler)))
	e.GET("/", getIndex)

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(CurrentConfig.Server.Port)))
}
