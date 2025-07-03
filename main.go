package main

import (
	"bytes"
	"fmt"
	"net/http"

	"gitea.elara.ws/Hazel/transfem-startpage/internal/diyhrt"
	"gitea.elara.ws/Hazel/transfem-startpage/internal/rendering"
	"github.com/labstack/echo/v4"
)

var CurrentRenderingConfig = rendering.DefaultRenderingConfig()

func FetchDiyHrt() error {
	fmt.Println("Fetch DiyHrt Marketplaces...")

	l, err := diyhrt.GetListings()
	if err != nil {
		return err
	}
	CurrentRenderingConfig.LoadDiyHrt(l)
	return nil
}

func setConfig(c echo.Context) error {
	err := c.Bind(&CurrentRenderingConfig)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "OK")
}

func getIndex(c echo.Context) error {
	var tpl bytes.Buffer
	rendering.IndexTemplate.Execute(&tpl, CurrentRenderingConfig)

	return c.HTML(http.StatusOK, tpl.String())
}

func main() {
	fmt.Println("running transfem startpage")
	err := FetchDiyHrt()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(CurrentRenderingConfig.Stores)

	e := echo.New()
	e.Static("/assets", "frontend/assets")
	e.GET("/", getIndex)

	// this is for me to later setup the ctl such that I can config the running program on the command line
	e.POST("/api/config", setConfig)

	e.Logger.Fatal(e.Start(":5500"))
}
