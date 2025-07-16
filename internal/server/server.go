package server

import (
	"log"
	"net/http"
	"strconv"

	"gitea.elara.ws/Hazel/transfem-startpage/internal/rendering"
	"github.com/labstack/echo/v4"
)

var Config = rendering.NewConfig()

func Start(profile string) error {
	err := Config.ScanForConfigFile(profile)
	if err != nil {
		return err
	}

	err = Config.Init()
	if err != nil {
		return err
	}

	err = Config.FetchDiyHrt()
	if err != nil {
		log.Println(err)
	}

	e := echo.New()

	// statically serve the file
	cacheDir, err := rendering.GetCacheDir()
	if err == nil {
		e.Static("/cache", cacheDir)
	} else {
		log.Println("didn't enable cache dir", err)
	}

	// https://echo.labstack.com/docs/cookbook/embed-resources
	staticHandler := http.FileServer(getFileSystem())
	e.GET("/assets/*", echo.WrapHandler(http.StripPrefix("/", staticHandler)))
	e.GET("/scripts/*", echo.WrapHandler(http.StripPrefix("/", staticHandler)))
	e.GET("/", getIndex)

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(Config.Server.Port)))
	return nil
}
